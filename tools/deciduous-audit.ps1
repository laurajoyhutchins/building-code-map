[CmdletBinding()]
param(
  [switch]$Strict
)

$ErrorActionPreference = 'Stop'

$repoRoot = (Resolve-Path (Join-Path $PSScriptRoot '..')).Path
$deciduousCmd = Join-Path $PSScriptRoot 'deciduous.cmd'

function Invoke-DeciduousText {
  param(
    [Parameter(Mandatory = $true)]
    [string[]]$Arguments
  )

  $output = & $deciduousCmd @Arguments 2>&1
  if ($LASTEXITCODE -ne 0) {
    throw "Deciduous command failed: $($Arguments -join ' ')`n$output"
  }

  return @($output)
}

function Get-DeciduousNodes {
  $lines = Invoke-DeciduousText -Arguments @('nodes')
  $nodes = @()

  foreach ($line in $lines) {
    if ($line -match '^\s*(\d+)\s+(\w+)\s+(\w+)\s+(.+?)\s*$') {
      $nodes += [pscustomobject]@{
        Id = [int]$matches[1]
        Type = $matches[2]
        Status = $matches[3]
        Title = $matches[4]
      }
    }
  }

  return $nodes
}

function Get-DeciduousEdges {
  $lines = Invoke-DeciduousText -Arguments @('edges')
  $edges = @()

  foreach ($line in $lines) {
    if ($line -match '^\s*(\d+)\s+(\d+)\s+(\d+)\s+(\w+)\s+(.+?)\s*$') {
      $edges += [pscustomobject]@{
        Id = [int]$matches[1]
        From = [int]$matches[2]
        To = [int]$matches[3]
        Type = $matches[4]
        Rationale = $matches[5]
      }
    }
  }

  return $edges
}

function Group-EdgeIdsByNode {
  param(
    [Parameter(Mandatory = $true)]
    [array]$Edges,

    [Parameter(Mandatory = $true)]
    [ValidateSet('From', 'To')]
    [string]$Direction
  )

  $map = @{}
  foreach ($edge in $Edges) {
    $key = [int]$edge.$Direction
    if (-not $map.ContainsKey($key)) {
      $map[$key] = @()
    }

    $map[$key] += $edge
  }

  return $map
}

function Get-EdgesForNode {
  param(
    [Parameter(Mandatory = $true)]
    [hashtable]$EdgeMap,

    [Parameter(Mandatory = $true)]
    [int]$NodeId
  )

  if ($EdgeMap.ContainsKey($NodeId)) {
    return ,@($EdgeMap[$NodeId])
  }

  return ,@()
}

function Get-ChildNodesOfType {
  param(
    [Parameter(Mandatory = $true)]
    [pscustomobject]$Node,

    [Parameter(Mandatory = $true)]
    [string]$Type,

    [Parameter(Mandatory = $true)]
    [hashtable]$OutgoingEdges,

    [Parameter(Mandatory = $true)]
    [hashtable]$NodeById
  )

  $children = @()
  foreach ($edge in (Get-EdgesForNode -EdgeMap $OutgoingEdges -NodeId $Node.Id)) {
    $child = $NodeById[$edge.To]
    if ($null -ne $child -and $child.Type -eq $Type) {
      $children += $child
    }
  }

  return $children
}

$nodes = Get-DeciduousNodes
$edges = Get-DeciduousEdges
$nodeById = @{}
foreach ($node in $nodes) {
  $nodeById[$node.Id] = $node
}

$outgoing = Group-EdgeIdsByNode -Edges $edges -Direction From
$incoming = Group-EdgeIdsByNode -Edges $edges -Direction To

$goalCoverageGaps = @()
$openGoalDecisions = @()
$decisionCoverageGaps = @()
$actionCoverageGaps = @()
$orphans = @()

foreach ($node in $nodes) {
  $hasIncoming = (Get-EdgesForNode -EdgeMap $incoming -NodeId $node.Id).Count -gt 0
  if (-not $hasIncoming -and $node.Type -ne 'goal') {
    $orphans += $node
  }

  switch ($node.Type) {
    'goal' {
      $options = Get-ChildNodesOfType -Node $node -Type 'option' -OutgoingEdges $outgoing -NodeById $nodeById
      $decisions = @()

      foreach ($option in $options) {
        $decisions += Get-ChildNodesOfType -Node $option -Type 'decision' -OutgoingEdges $outgoing -NodeById $nodeById
      }

      if ($options.Count -eq 0) {
        $goalCoverageGaps += $node
      } elseif ($decisions.Count -eq 0) {
        $openGoalDecisions += $node
      }
    }
    'decision' {
      $actions = Get-ChildNodesOfType -Node $node -Type 'action' -OutgoingEdges $outgoing -NodeById $nodeById
      if ($actions.Count -eq 0) {
        $decisionCoverageGaps += $node
      }
    }
    'action' {
      $outcomes = Get-ChildNodesOfType -Node $node -Type 'outcome' -OutgoingEdges $outgoing -NodeById $nodeById
      if ($outcomes.Count -eq 0) {
        $actionCoverageGaps += $node
      }
    }
  }
}

$summary = [pscustomobject]@{
  NodeCount = $nodes.Count
  EdgeCount = $edges.Count
  GoalsMissingOptions = $goalCoverageGaps.Count
  GoalsAwaitingDecision = $openGoalDecisions.Count
  DecisionsMissingActions = $decisionCoverageGaps.Count
  ActionsMissingOutcomes = $actionCoverageGaps.Count
  OrphanNonGoals = $orphans.Count
}

Write-Host '=== DECIDUOUS AUDIT ==='
$summary
Write-Host ''

if ($goalCoverageGaps.Count -gt 0) {
  Write-Host 'Goals Missing Options:'
  $goalCoverageGaps | Select-Object Id, Title | Format-Table -AutoSize
  Write-Host ''
}

if ($openGoalDecisions.Count -gt 0) {
  Write-Host 'Goals Awaiting Decision:'
  $openGoalDecisions | Select-Object Id, Title | Format-Table -AutoSize
  Write-Host ''
}

if ($decisionCoverageGaps.Count -gt 0) {
  Write-Host 'Decisions Missing Actions:'
  $decisionCoverageGaps | Select-Object Id, Title | Format-Table -AutoSize
  Write-Host ''
}

if ($actionCoverageGaps.Count -gt 0) {
  Write-Host 'Actions Missing Outcomes:'
  $actionCoverageGaps | Select-Object Id, Title | Format-Table -AutoSize
  Write-Host ''
}

if ($orphans.Count -gt 0) {
  Write-Host 'Orphan Non-Goal Nodes:'
  $orphans | Select-Object Id, Type, Title | Format-Table -AutoSize
  Write-Host ''
}

if ($goalCoverageGaps.Count -eq 0 -and $openGoalDecisions.Count -eq 0 -and $decisionCoverageGaps.Count -eq 0 -and $actionCoverageGaps.Count -eq 0 -and $orphans.Count -eq 0) {
  Write-Host 'No graph hygiene gaps found.'
}

if ($Strict -and ($goalCoverageGaps.Count -gt 0 -or $decisionCoverageGaps.Count -gt 0 -or $actionCoverageGaps.Count -gt 0 -or $orphans.Count -gt 0)) {
  throw 'Deciduous audit found graph hygiene gaps.'
}
