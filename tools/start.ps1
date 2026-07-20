[CmdletBinding(DefaultParameterSetName = 'Both')]
param(
  [Parameter(ParameterSetName = 'FrontendOnly')]
  [switch]$FrontendOnly,

  [Parameter(ParameterSetName = 'BackendOnly')]
  [switch]$BackendOnly
)

$ErrorActionPreference = 'Stop'

$repoRoot = (Resolve-Path (Join-Path $PSScriptRoot '..')).Path
$stateRoot = Join-Path $PSScriptRoot '.state'
$logRoot = Join-Path $stateRoot 'logs'
$frontendPort = 5173
$backendPort = 8000
$stateFile = Join-Path $stateRoot 'launch-state.json'
$goPath = $null

function Initialize-ProcessEnvironment {
  $machinePath = [System.Environment]::GetEnvironmentVariable('Path', 'Machine')
  $userPath = [System.Environment]::GetEnvironmentVariable('Path', 'User')
  $pathParts = @()

  if (-not [string]::IsNullOrWhiteSpace($machinePath)) {
    $pathParts += $machinePath
  }

  if (-not [string]::IsNullOrWhiteSpace($userPath)) {
    $pathParts += $userPath
  }

  Remove-Item Env:PATH -ErrorAction SilentlyContinue
  $env:Path = ($pathParts -join ';')
}

function Invoke-CheckedWebRequest {
  param([string]$Url)

  if ($PSVersionTable.PSVersion.Major -lt 6) {
    return Invoke-WebRequest -Uri $Url -TimeoutSec 5 -UseBasicParsing
  }

  return Invoke-WebRequest -Uri $Url -TimeoutSec 5
}

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
      try {
        $listener.Stop()
      } catch {
      }
    }
  }
}

function Get-LogTail {
  param(
    [string]$Path,
    [int]$Tail = 40
  )

  if (-not (Test-Path $Path)) {
    return ''
  }

  $lines = Get-Content -Path $Path -Tail $Tail -ErrorAction SilentlyContinue
  return ($lines -join [Environment]::NewLine)
}

function Write-LaunchState {
  param([array]$Services)

  New-Item -ItemType Directory -Force -Path $stateRoot | Out-Null
  $json = @($Services) | ConvertTo-Json -Depth 6
  Set-Content -Path $stateFile -Value $json -Encoding utf8
}

function Remove-LaunchState {
  Remove-Item -Path $stateFile -Force -ErrorAction SilentlyContinue
}

function Get-TrackedProcessFingerprint {
  param([System.Diagnostics.Process]$Process)

  $Process.Refresh()
  return [int64]$Process.StartTime.ToUniversalTime().ToFileTimeUtc()
}

function Test-TrackedProcess {
  param([psobject]$Service)

  if ($null -eq $Service.Id) {
    return $false
  }

  try {
    $process = Get-Process -Id $Service.Id -ErrorAction Stop
  } catch {
    return $false
  }

  if ($null -ne $Service.StartTimeFileTimeUtc) {
    try {
      $startTimeFileTimeUtc = [int64]$process.StartTime.ToUniversalTime().ToFileTimeUtc()
      if ($startTimeFileTimeUtc -ne [int64]$Service.StartTimeFileTimeUtc) {
        return $false
      }
    } catch {
      return $false
    }
  }

  return $true
}

function Stop-ProcessOnPort {
  param([int]$Port)

  try {
    $connections = Get-NetTCPConnection -LocalPort $Port -State Listen -ErrorAction Stop
    foreach ($conn in $connections) {
      if ($null -ne $conn.OwningProcess -and $conn.OwningProcess -gt 0) {
        try {
          & taskkill.exe /PID $conn.OwningProcess /T /F 2>&1 | Out-Null
        } catch {
        }
      }
    }
  } catch {
    try {
      $procs = Get-CimInstance Win32_Process -Filter "CommandLine LIKE '%--port $Port%' OR CommandLine LIKE '%:$Port%'" -ErrorAction Stop
      foreach ($proc in $procs) {
        try {
          & taskkill.exe /PID $proc.ProcessId /T /F 2>&1 | Out-Null
        } catch {
        }
      }
    } catch {
    }
  }
}

function Get-PortFromUrl {
  param([string]$Url)

  if ([string]::IsNullOrWhiteSpace($Url)) {
    return $null
  }

  try {
    $uri = [System.Uri]::new($Url)
    return $uri.Port
  } catch {
    return $null
  }
}

function Stop-TrackedProcess {
  param([psobject]$Service)

  if (-not (Test-TrackedProcess -Service $Service)) {
    return
  }

  try {
    Stop-Process -Id $Service.Id -Force -ErrorAction SilentlyContinue
  } catch {
  }

  for ($attempt = 0; $attempt -lt 10; $attempt++) {
    if (-not (Test-TrackedProcess -Service $Service)) {
      break
    }

    Start-Sleep -Milliseconds 300
  }

  if (Test-TrackedProcess -Service $Service) {
    try {
      & taskkill.exe /PID $Service.Id /T /F | Out-Null
    } catch {
    }

    for ($attempt = 0; $attempt -lt 10; $attempt++) {
      if (-not (Test-TrackedProcess -Service $Service)) {
        break
      }

      Start-Sleep -Milliseconds 300
    }
  }

  $port = Get-PortFromUrl -Url $Service.Url
  if ($null -ne $port) {
    Start-Sleep -Milliseconds 500
    Stop-ProcessOnPort -Port $port
  }
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

  $stdoutPath = Join-Path $logRoot "$Name.out.log"
  $stderrPath = Join-Path $logRoot "$Name.err.log"
  Remove-Item -Path $stdoutPath, $stderrPath -Force -ErrorAction SilentlyContinue

  $process = Start-Process `
    -FilePath $FilePath `
    -ArgumentList $ArgumentList `
    -WorkingDirectory $WorkingDirectory `
    -PassThru `
    -WindowStyle Hidden `
    -RedirectStandardOutput $stdoutPath `
    -RedirectStandardError $stderrPath

  $startTimeFileTimeUtc = $null
  for ($attempt = 0; $attempt -lt 10; $attempt++) {
    try {
      $startTimeFileTimeUtc = Get-TrackedProcessFingerprint -Process $process
      break
    } catch {
      Start-Sleep -Milliseconds 100
    }
  }

  if ($null -eq $startTimeFileTimeUtc) {
    throw "$Name started, but its process fingerprint could not be captured."
  }

  $trackedProcess = [pscustomobject]@{
    Id = $process.Id
    StartTimeFileTimeUtc = $startTimeFileTimeUtc
  }

  $deadline = (Get-Date).AddSeconds(30)
  $ready = $false
  while ((Get-Date) -lt $deadline) {
    if (-not (Test-TrackedProcess -Service $trackedProcess)) {
      throw "$Name exited immediately after launch."
    }

    try {
      Invoke-CheckedWebRequest -Url $ReadyUrl | Out-Null
      $ready = $true
      break
    } catch {
      Start-Sleep -Seconds 1
    }
  }

  if (-not $ready) {
    Stop-TrackedProcess -Service $trackedProcess

    $logTail = @(
      "=== $Name stdout ==="
      (Get-LogTail -Path $stdoutPath)
      "=== $Name stderr ==="
      (Get-LogTail -Path $stderrPath)
    ) -join [Environment]::NewLine

    throw "$Name did not become ready at $ReadyUrl within 30 seconds.`n$logTail"
  }

  [pscustomobject]@{
    Name = $Name
    Id = $process.Id
    Url = $ReadyUrl
    Stdout = $stdoutPath
    Stderr = $stderrPath
    StartTimeFileTimeUtc = $startTimeFileTimeUtc
  }
}

$nodePath = $null
if (-not $BackendOnly) {
  foreach ($candidate in @('node.exe', 'node')) {
    $command = Get-Command $candidate -ErrorAction SilentlyContinue | Select-Object -First 1
    if ($null -ne $command) {
      if ($null -ne $command.Path -and $command.Path.Length -gt 0) {
        $nodePath = $command.Path
        break
      }

      if ($null -ne $command.Source -and $command.Source.Length -gt 0) {
        $nodePath = $command.Source
        break
      }
    }
  }

  if ($null -eq $nodePath) {
    throw 'Unable to locate node on PATH.'
  }
}

if (-not $FrontendOnly) {
  if (-not (Test-PortAvailable -Port $backendPort)) {
    throw "Port $backendPort is already in use."
  }
}

if (-not $BackendOnly) {
  if (-not (Test-PortAvailable -Port $frontendPort)) {
    throw "Port $frontendPort is already in use."
  }
}

foreach ($candidate in @('go.exe', 'go')) {
  $command = Get-Command $candidate -ErrorAction SilentlyContinue | Select-Object -First 1
  if ($null -ne $command) {
    if ($null -ne $command.Path -and $command.Path.Length -gt 0) {
      $goPath = $command.Path
      break
    }

    if ($null -ne $command.Source -and $command.Source.Length -gt 0) {
      $goPath = $command.Source
      break
    }
  }
}

$services = @()

try {
  Initialize-ProcessEnvironment

  if (-not $FrontendOnly) {
    if ($null -eq $goPath) {
      throw 'Unable to locate go on PATH.'
    }

    $services += Start-LoggedProcess `
      -Name 'backend' `
      -FilePath $goPath `
      -ArgumentList @('run', './cmd/server', '--addr', '127.0.0.1:8000') `
      -WorkingDirectory (Join-Path $repoRoot 'backend') `
      -ReadyUrl 'http://127.0.0.1:8000/ready'
    Write-LaunchState -Services $services
  }

  if (-not $BackendOnly) {
    $devScript = Join-Path $repoRoot 'tools\dev.mjs'
    if (-not (Test-Path $devScript)) {
      throw "Unable to locate dev launcher at $devScript."
    }

    $services += Start-LoggedProcess `
      -Name 'frontend' `
      -FilePath $nodePath `
      -ArgumentList @($devScript) `
      -WorkingDirectory $repoRoot `
      -ReadyUrl 'http://127.0.0.1:5173/'
    Write-LaunchState -Services $services
  }
} catch {
  foreach ($service in $services) {
    try {
      Stop-TrackedProcess -Service $service
    } catch {
    }
  }

  Remove-LaunchState
  throw
}

Write-LaunchState -Services $services
$services

if ($services.Count -gt 0) {
  Wait-Process -Id @($services | ForEach-Object { $_.Id })
  Remove-LaunchState
}
