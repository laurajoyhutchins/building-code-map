[CmdletBinding()]
param(
    [string] $Repository = "laurajoyhutchins/building-code-map",
    [string] $RulesetName = "Exact-head review clearance"
)

$ErrorActionPreference = "Stop"
$apiVersion = "2026-03-10"

function Invoke-GitHubApi {
    param(
        [Parameter(Mandatory)]
        [ValidateSet("GET", "POST", "PUT", "PATCH", "DELETE")]
        [string] $Method,

        [Parameter(Mandatory)]
        [string] $Endpoint,

        [object] $Body
    )

    $arguments = @(
        "api",
        "--method", $Method,
        "--header", "Accept: application/vnd.github+json",
        "--header", "X-GitHub-Api-Version: $apiVersion",
        $Endpoint
    )

    if ($null -eq $Body) {
        & gh @arguments
    }
    else {
        $json = $Body | ConvertTo-Json -Depth 20 -Compress
        $json | & gh @arguments --input -
    }

    if ($LASTEXITCODE -ne 0) {
        throw "GitHub API request failed: $Method $Endpoint"
    }
}

function Invoke-OptionalGitHubApi {
    param(
        [Parameter(Mandatory)]
        [ValidateSet("PUT", "PATCH")]
        [string] $Method,

        [Parameter(Mandatory)]
        [string] $Endpoint,

        [object] $Body,

        [Parameter(Mandatory)]
        [string] $FeatureName
    )

    try {
        Invoke-GitHubApi -Method $Method -Endpoint $Endpoint -Body $Body | Out-Null
        Write-Host "Enabled $FeatureName."
    }
    catch {
        Write-Warning "Could not enable $FeatureName automatically. Review the repository's Code security and analysis settings. $($_.Exception.Message)"
    }
}

if (-not (Get-Command gh -ErrorAction SilentlyContinue)) {
    throw "GitHub CLI (gh) is required. Install it and authenticate with an account that has repository administration permission."
}

& gh auth status
if ($LASTEXITCODE -ne 0) {
    throw "GitHub CLI is not authenticated. Run 'gh auth login' first."
}

Write-Host "Configuring repository settings for $Repository..."

$repositorySettings = @{
    allow_squash_merge          = $true
    allow_merge_commit          = $false
    allow_rebase_merge          = $false
    allow_auto_merge            = $false
    delete_branch_on_merge      = $true
    allow_update_branch         = $true
    squash_merge_commit_title   = "PR_TITLE"
    squash_merge_commit_message = "PR_BODY"
}

Invoke-GitHubApi -Method PATCH -Endpoint "repos/$Repository" -Body $repositorySettings | Out-Null

$ruleset = @{
    name        = $RulesetName
    target      = "branch"
    enforcement = "active"
    conditions  = @{
        ref_name = @{
            include = @("~DEFAULT_BRANCH")
            exclude = @()
        }
    }
    rules       = @(
        @{
            type = "deletion"
        },
        @{
            type = "non_fast_forward"
        },
        @{
            type = "required_linear_history"
        },
        @{
            type       = "pull_request"
            parameters = @{
                allowed_merge_methods              = @("squash")
                dismiss_stale_reviews_on_push      = $true
                require_code_owner_review          = $false
                require_last_push_approval         = $false
                required_approving_review_count    = 0
                required_review_thread_resolution = $true
            }
        },
        @{
            type       = "required_status_checks"
            parameters = @{
                do_not_enforce_on_create              = $true
                required_status_checks                = @(
                    @{ context = "Frontend" },
                    @{ context = "Backend" }
                )
                strict_required_status_checks_policy = $true
            }
        }
    )
}

$existingRulesets = Invoke-GitHubApi -Method GET -Endpoint "repos/$Repository/rulesets?per_page=100"
$matchingRuleset = @($existingRulesets | ConvertFrom-Json) | Where-Object { $_.name -eq $RulesetName } | Select-Object -First 1

if ($null -eq $matchingRuleset) {
    Write-Host "Creating ruleset '$RulesetName'..."
    Invoke-GitHubApi -Method POST -Endpoint "repos/$Repository/rulesets" -Body $ruleset | Out-Null
}
else {
    Write-Host "Updating ruleset '$RulesetName' (ID $($matchingRuleset.id))..."
    Invoke-GitHubApi -Method PUT -Endpoint "repos/$Repository/rulesets/$($matchingRuleset.id)" -Body $ruleset | Out-Null
}

Invoke-OptionalGitHubApi -Method PUT -Endpoint "repos/$Repository/private-vulnerability-reporting" -FeatureName "private vulnerability reporting"
Invoke-OptionalGitHubApi -Method PUT -Endpoint "repos/$Repository/vulnerability-alerts" -FeatureName "Dependabot alerts"
Invoke-OptionalGitHubApi -Method PUT -Endpoint "repos/$Repository/automated-security-fixes" -FeatureName "Dependabot security updates"

$secretScanningSettings = @{
    security_and_analysis = @{
        secret_scanning = @{
            status = "enabled"
        }
        secret_scanning_push_protection = @{
            status = "enabled"
        }
    }
}

Invoke-OptionalGitHubApi -Method PATCH -Endpoint "repos/$Repository" -Body $secretScanningSettings -FeatureName "secret scanning and push protection"

Write-Host "Verifying repository settings..."
$repository = Invoke-GitHubApi -Method GET -Endpoint "repos/$Repository" | ConvertFrom-Json
$verifiedRulesets = Invoke-GitHubApi -Method GET -Endpoint "repos/$Repository/rulesets?per_page=100" | ConvertFrom-Json
$verifiedRuleset = @($verifiedRulesets) | Where-Object { $_.name -eq $RulesetName } | Select-Object -First 1

[pscustomobject]@{
    Repository             = $repository.full_name
    Visibility             = $repository.visibility
    DefaultBranch          = $repository.default_branch
    SquashMerge            = $repository.allow_squash_merge
    MergeCommits           = $repository.allow_merge_commit
    RebaseMerge            = $repository.allow_rebase_merge
    AutoMerge              = $repository.allow_auto_merge
    DeleteBranchOnMerge    = $repository.delete_branch_on_merge
    AllowUpdateBranch      = $repository.allow_update_branch
    Ruleset                = $verifiedRuleset.name
    RulesetEnforcement     = $verifiedRuleset.enforcement
    SecretScanning         = $repository.security_and_analysis.secret_scanning.status
    SecretPushProtection   = $repository.security_and_analysis.secret_scanning_push_protection.status
} | Format-List

Write-Host "Repository hardening configuration complete."
