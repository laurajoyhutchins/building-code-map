# Launch Scripts Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a single Windows PowerShell launcher for the frontend and backend, plus a reusable health-check script.

**Architecture:** Put the runnable entrypoints in `tools/` and keep them self-contained. The launcher should resolve the repo root, start each service in the right working directory, and fail fast when a requested port is already occupied. The health script should probe the backend and frontend over HTTP so launch success is visible without opening the browser first.

**Tech Stack:** PowerShell 7+, local Go backend via `go run`, npm/Vite, HTTP probes via `Invoke-WebRequest`.

---

### Task 1: Add the launcher script

**Files:**
- Create: `tools/start.ps1`

- [ ] **Step 1: Write the launcher**

```powershell
[CmdletBinding(DefaultParameterSetName = 'Both')]
param(
  [Parameter(ParameterSetName = 'FrontendOnly')]
  [switch]$FrontendOnly,

  [Parameter(ParameterSetName = 'BackendOnly')]
  [switch]$BackendOnly
)

$ErrorActionPreference = 'Stop'

$repoRoot = (Resolve-Path (Join-Path $PSScriptRoot '..')).Path
$toolsRoot = $PSScriptRoot
$logRoot = Join-Path $toolsRoot '.logs'
$frontendPort = 5173
$backendPort = 8000

function Test-PortAvailable {
  param([int]$Port)

  $listener = $null
  try {
    $listener = [System.Net.Sockets.TcpListener]::new([System.Net.IPAddress]::Loopback, $Port)
    $listener.Start()
    return $true
  } catch {
    return $false
  } finally {
    if ($null -ne $listener) {
      try { $listener.Stop() } catch { }
    }
  }
}

function Get-LogTail {
  param([string]$Path)

  if (-not (Test-Path $Path)) {
    return ''
  }

  $lines = Get-Content -Path $Path -Tail 40 -ErrorAction SilentlyContinue
  return ($lines -join [Environment]::NewLine)
}

function Wait-ForEndpoint {
  param(
    [string]$Name,
    [string]$Url,
    [int]$TimeoutSeconds = 45
  )

  $deadline = (Get-Date).AddSeconds($TimeoutSeconds)
  while ((Get-Date) -lt $deadline) {
    try {
      $response = Invoke-WebRequest -Uri $Url -TimeoutSec 5
      if ($response.StatusCode -ge 200 -and $response.StatusCode -lt 300) {
        return
      }
    } catch {
      Start-Sleep -Milliseconds 500
      continue
    }

    Start-Sleep -Milliseconds 500
  }

  throw "$Name did not become ready at $Url within $TimeoutSeconds seconds."
}

function Start-LoggedProcess {
  param(
    [string]$Name,
    [string]$FilePath,
    [string[]]$ArgumentList,
    [string]$WorkingDirectory,
    [string]$ReadyUrl
  )

  New-Item -ItemType Directory -Force -Path $logRoot | Out-Null

  $stdout = Join-Path $logRoot "$Name.out.log"
  $stderr = Join-Path $logRoot "$Name.err.log"
  Remove-Item -Path $stdout, $stderr -Force -ErrorAction SilentlyContinue

  $process = Start-Process `
    -FilePath $FilePath `
    -ArgumentList $ArgumentList `
    -WorkingDirectory $WorkingDirectory `
    -PassThru `
    -WindowStyle Hidden `
    -RedirectStandardOutput $stdout `
    -RedirectStandardError $stderr

  try {
    Wait-ForEndpoint -Name $Name -Url $ReadyUrl
  } catch {
    try { Stop-Process -Id $process.Id -Force } catch { }

    $logTail = @(
      "=== $Name stdout ==="
      (Get-LogTail -Path $stdout)
      "=== $Name stderr ==="
      (Get-LogTail -Path $stderr)
    ) -join [Environment]::NewLine

    throw "$Name failed to start.`n$logTail"
  }

  [pscustomobject]@{
    Name = $Name
    Id = $process.Id
    Url = $ReadyUrl
    Stdout = $stdout
    Stderr = $stderr
  }
}

$services = @()

if (-not $BackendOnly) {
  if (-not (Test-PortAvailable -Port $frontendPort)) {
    throw "Port $frontendPort is already in use."
  }
}

if (-not $FrontendOnly) {
  if (-not (Test-PortAvailable -Port $backendPort)) {
    throw "Port $backendPort is already in use."
  }
}

try {
  if (-not $FrontendOnly) {
    $goPath = (Get-Command go.exe -ErrorAction Stop).Path

    $services += Start-LoggedProcess `
      -Name 'backend' `
      -FilePath $goPath `
      -ArgumentList @('run', './cmd/server', '--addr', '127.0.0.1:8000') `
      -WorkingDirectory (Join-Path $repoRoot 'backend') `
      -ReadyUrl 'http://127.0.0.1:8000/ready'
  }

  if (-not $BackendOnly) {
    $npmPath = (Get-Command npm.cmd -ErrorAction Stop).Path

    $services += Start-LoggedProcess `
      -Name 'frontend' `
      -FilePath $npmPath `
      -ArgumentList @('run', 'dev', '--', '--host', '127.0.0.1', '--port', '5173', '--strictPort') `
      -WorkingDirectory $repoRoot `
      -ReadyUrl 'http://127.0.0.1:5173/'
  }
} catch {
  foreach ($service in $services) {
    try { Stop-Process -Id $service.Id -Force } catch { }
  }

  throw
}

$services
```

- [ ] **Step 2: Verify the script starts the requested service set**

Run:
```powershell
.\tools\start.ps1 -BackendOnly
```
Expected: backend starts, the script returns an object with `backend`, and `http://127.0.0.1:8000/health` responds.

- [ ] **Step 3: Verify the two-service path**

Run:
```powershell
.\tools\start.ps1
```
Expected: both frontend and backend start, and the script returns two service objects.

### Task 2: Add the health-check script

**Files:**
- Create: `tools/health.ps1`

- [ ] **Step 1: Write the health probe**

```powershell
[CmdletBinding(DefaultParameterSetName = 'Both')]
param(
  [Parameter(ParameterSetName = 'FrontendOnly')]
  [switch]$FrontendOnly,

  [Parameter(ParameterSetName = 'BackendOnly')]
  [switch]$BackendOnly
)

$ErrorActionPreference = 'Stop'

function Test-Endpoint {
  param(
    [string]$Name,
    [string]$Url
  )

  try {
    $response = Invoke-WebRequest -Uri $Url -TimeoutSec 5
    [pscustomobject]@{
      Name = $Name
      Url = $Url
      StatusCode = $response.StatusCode
      Healthy = $true
    }
  } catch {
    [pscustomobject]@{
      Name = $Name
      Url = $Url
      StatusCode = $null
      Healthy = $false
      Error = $_.Exception.Message
    }
  }
}

$results = @()

if (-not $BackendOnly) {
  $results += Test-Endpoint -Name 'frontend' -Url 'http://127.0.0.1:5173/'
}

if (-not $FrontendOnly) {
  $results += Test-Endpoint -Name 'backend' -Url 'http://127.0.0.1:8000/health'
}

$failures = @($results | Where-Object { -not $_.Healthy })
if ($failures.Count -gt 0) {
  $failures
  throw 'One or more health checks failed.'
}

$results
```

- [ ] **Step 2: Verify the backend-only probe**

Run:
```powershell
.\tools\health.ps1 -BackendOnly
```
Expected: object output with `Healthy = True` for the backend.

- [ ] **Step 3: Verify the combined probe**

Run:
```powershell
.\tools\health.ps1
```
Expected: object output with both frontend and backend healthy when the launcher is running.

### Task 3: Update repo docs and ignore rules

**Files:**
- Modify: `.gitignore`
- Modify: `README.md`
- Modify: `docs/reference/configuration.md`

- [ ] **Step 1: Ignore the script log directory**

```gitignore
tools/.logs/
```

- [ ] **Step 2: Document the new launch flow**

```markdown
## Local Launch Scripts

- `.\tools\start.ps1`: start the frontend and backend together.
- `.\tools\start.ps1 -FrontendOnly`: start only the frontend.
- `.\tools\start.ps1 -BackendOnly`: start only the backend.
- `.\tools\health.ps1`: probe both services.
- `.\tools\health.ps1 -FrontendOnly`: probe only the frontend.
- `.\tools\health.ps1 -BackendOnly`: probe only the backend.
```

- [ ] **Step 3: Update the configuration reference**

```markdown
## Operational Scripts

- [tools/start.ps1](/tools/start.ps1): Windows launcher for local frontend/backend development.
- [tools/health.ps1](/tools/health.ps1): HTTP health checks for the launched services.
```

- [ ] **Step 4: Verify docs format stays intact**

Run:
```powershell
npm run format:check
```
Expected: pass with the new script and docs entries formatted consistently.
