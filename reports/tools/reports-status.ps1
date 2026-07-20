[CmdletBinding()]
param(
  [Parameter(ValueFromRemainingArguments = $true)]
  [string[]]$Arguments
)

$ErrorActionPreference = 'Stop'

$scriptPath = Join-Path $PSScriptRoot 'generate_report_tracker.py'
if (-not (Test-Path $scriptPath)) {
  throw "Unable to locate tracker script at $scriptPath."
}

$pythonPath = $null
foreach ($candidate in @('python.exe', 'python', 'py.exe')) {
  $command = Get-Command $candidate -ErrorAction SilentlyContinue | Select-Object -First 1
  if ($null -ne $command) {
    if ($null -ne $command.Path -and $command.Path.Length -gt 0) {
      $pythonPath = $command.Path
      break
    }

    if ($null -ne $command.Source -and $command.Source.Length -gt 0) {
      $pythonPath = $command.Source
      break
    }
  }
}

if ($null -eq $pythonPath) {
  throw 'Unable to locate Python on PATH.'
}

& $pythonPath $scriptPath @Arguments
exit $LASTEXITCODE
