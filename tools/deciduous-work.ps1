[CmdletBinding()]
param(
  [Parameter(Mandatory = $true, Position = 0)]
  [string]$Title,

  [Parameter()]
  [int]$Confidence = 90,

  [switch]$DryRun
)

$ErrorActionPreference = 'Stop'

$deciduousCmd = Join-Path $PSScriptRoot 'deciduous.cmd'
$recoverCmd = Join-Path $PSScriptRoot 'deciduous-recover.cmd'

$promptText = [Console]::In.ReadToEnd()
if ([string]::IsNullOrWhiteSpace($promptText)) {
  throw 'Provide the verbatim user request on stdin when starting a Deciduous work transaction.'
}

Write-Host '=== Create Goal ==='
if ($DryRun) {
  Write-Host "Dry run only. Would create goal:`n- Title: $Title`n- Confidence: $Confidence"
  Write-Host '- Prompt:'
  Write-Host $promptText
  exit 0
}

$promptText | & $deciduousCmd add goal $Title -c $Confidence --prompt-stdin
if ($LASTEXITCODE -ne 0) {
  throw 'Failed to create Deciduous goal.'
}

Write-Host ''
Write-Host '=== Recovery Summary ==='
& $recoverCmd
if ($LASTEXITCODE -ne 0) {
  throw 'Failed to run Deciduous recovery summary.'
}
