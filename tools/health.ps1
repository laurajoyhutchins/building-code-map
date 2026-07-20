[CmdletBinding(DefaultParameterSetName = 'Both')]
param(
  [Parameter(ParameterSetName = 'FrontendOnly')]
  [switch]$FrontendOnly,

  [Parameter(ParameterSetName = 'BackendOnly')]
  [switch]$BackendOnly
)

$ErrorActionPreference = 'Stop'

function Invoke-CheckedWebRequest {
  param([string]$Url)

  if ($PSVersionTable.PSVersion.Major -lt 6) {
    return Invoke-WebRequest -Uri $Url -TimeoutSec 5 -UseBasicParsing
  }

  return Invoke-WebRequest -Uri $Url -TimeoutSec 5
}

function Test-Endpoint {
  param(
    [string]$Name,
    [string]$Url
  )

  try {
    $response = Invoke-CheckedWebRequest -Url $Url
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
  $results += Test-Endpoint -Name 'backend' -Url 'http://127.0.0.1:8000/ready'
}

$failures = @($results | Where-Object { -not $_.Healthy })
if ($failures.Count -gt 0) {
  $failures
  throw 'One or more health checks failed.'
}

$results
