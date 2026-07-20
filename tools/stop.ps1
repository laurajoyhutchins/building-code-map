[CmdletBinding(DefaultParameterSetName = 'Both')]
param(
  [Parameter(ParameterSetName = 'FrontendOnly')]
  [switch]$FrontendOnly,

  [Parameter(ParameterSetName = 'BackendOnly')]
  [switch]$BackendOnly
)

$ErrorActionPreference = 'Stop'

$stateFile = Join-Path $PSScriptRoot '.state\launch-state.json'

if (-not (Test-Path $stateFile)) {
  Write-Output 'No launch state file found.'
  exit 0
}

if ([string]::IsNullOrWhiteSpace((Get-Content -Path $stateFile -Raw))) {
  Remove-Item -Path $stateFile -Force -ErrorAction SilentlyContinue
  Write-Output 'Launch state file was empty.'
  exit 0
}

function Read-LaunchState {
  $services = (Get-Content -Path $stateFile -Raw) | ConvertFrom-Json
  if ($services -isnot [System.Collections.IEnumerable]) {
    return @($services)
  }

  return @($services)
}

function Write-LaunchState {
  param([array]$Services)

  if ($Services.Count -gt 0) {
    $Services | ConvertTo-Json -Depth 6 | Set-Content -Path $stateFile -Encoding utf8
  } else {
    Remove-Item -Path $stateFile -Force -ErrorAction SilentlyContinue
  }
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
    return $true
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

  return -not (Test-TrackedProcess -Service $Service)
}

$services = Read-LaunchState
$targets = $services

if ($FrontendOnly) {
  $targets = @($services | Where-Object { $_.Name -eq 'frontend' })
} elseif ($BackendOnly) {
  $targets = @($services | Where-Object { $_.Name -eq 'backend' })
}

$failedStops = @()

foreach ($service in $targets) {
  if (-not (Stop-TrackedProcess -Service $service)) {
    $failedStops += $service.Name
  }
}

if ($FrontendOnly) {
  $remaining = @($services | Where-Object { $_.Name -ne 'frontend' -and (Test-TrackedProcess -Service $_) })
} elseif ($BackendOnly) {
  $remaining = @($services | Where-Object { $_.Name -ne 'backend' -and (Test-TrackedProcess -Service $_) })
} else {
  $remaining = @($services | Where-Object { Test-TrackedProcess -Service $_ })
}

Write-LaunchState -Services $remaining

$targets | Select-Object Name, Id, Url

if ($failedStops.Count -gt 0) {
  throw "One or more services could not be stopped: $($failedStops -join ', ')."
}
