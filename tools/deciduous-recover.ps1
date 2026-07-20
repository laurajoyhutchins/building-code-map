[CmdletBinding()]
param()

$ErrorActionPreference = 'Stop'

$deciduousCmd = Join-Path $PSScriptRoot 'deciduous.cmd'
$auditCmd = Join-Path $PSScriptRoot 'deciduous-audit.cmd'

function Invoke-RepoCommand {
  param(
    [Parameter(Mandatory = $true)]
    [string]$Label,

    [Parameter(Mandatory = $true)]
    [string]$FilePath,

    [Parameter()]
    [string[]]$ArgumentList = @()
  )

  Write-Host "=== $Label ==="
  & $FilePath @ArgumentList
  if ($LASTEXITCODE -ne 0) {
    throw "$Label failed."
  }

  Write-Host ''
}

Write-Host 'Deciduous session recovery for this repo'
Write-Host ''

Invoke-RepoCommand -Label 'Pulse Summary' -FilePath $deciduousCmd -ArgumentList @('pulse', '--summary')
Invoke-RepoCommand -Label 'Graph Audit' -FilePath $auditCmd

Write-Host 'Next checks:'
Write-Host '- If a goal is awaiting decision, decide whether the repo has already chosen an approach or the question is still open.'
Write-Host '- If you start a new repo-visible initiative, create a goal with the verbatim prompt before implementation.'
Write-Host '- Use docs/how-to/use-decision-graph-for-repo-work.md for the repo-specific logging rules.'
