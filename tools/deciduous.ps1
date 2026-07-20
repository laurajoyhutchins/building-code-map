[CmdletBinding()]
param(
  [Parameter(ValueFromRemainingArguments = $true)]
  [string[]]$Arguments
)

$ErrorActionPreference = 'Stop'

function Get-DeciduousCommandPath {
  $candidates = @(
    (Get-Command deciduous -ErrorAction SilentlyContinue | Select-Object -First 1),
    (Join-Path $PSScriptRoot '..\cache\cargo\bin\deciduous.exe'),
    (Join-Path $env:USERPROFILE '.cargo\bin\deciduous.exe'),
    (Join-Path $env:USERPROFILE '.cargo\bin\deciduous.cmd'),
    (Join-Path $env:LOCALAPPDATA 'Programs\Deciduous\deciduous.exe'),
    'C:\Program Files\Deciduous\deciduous.exe',
    'C:\Program Files (x86)\Deciduous\deciduous.exe'
  )

  foreach ($candidate in $candidates) {
    if ($null -eq $candidate) {
      continue
    }

    if ($candidate -is [System.Management.Automation.CommandInfo]) {
      if ($null -ne $candidate.Path -and (Test-Path $candidate.Path)) {
        return $candidate.Path
      }

      if ($null -ne $candidate.Source -and (Test-Path $candidate.Source)) {
        return $candidate.Source
      }

      continue
    }

    if (Test-Path $candidate) {
      return $candidate
    }
  }

  return $null
}

$deciduousPath = Get-DeciduousCommandPath

if ($null -eq $deciduousPath) {
  $repoRoot = (Resolve-Path (Join-Path $PSScriptRoot '..')).Path
  throw @"
Unable to locate a deciduous executable.

Checked PATH and these common locations:
- $(Join-Path $repoRoot 'cache\cargo\bin\deciduous.exe')
- $(Join-Path $env:USERPROFILE '.cargo\bin\deciduous.exe')
- $(Join-Path $env:USERPROFILE '.cargo\bin\deciduous.cmd')
- $(Join-Path $env:LOCALAPPDATA 'Programs\Deciduous\deciduous.exe')
- C:\Program Files\Deciduous\deciduous.exe
- C:\Program Files (x86)\Deciduous\deciduous.exe
"@
}

& $deciduousPath @Arguments
