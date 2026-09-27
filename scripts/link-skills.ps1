# link-skills.ps1 — mount this repository once under the shared Agent Skills root.
# Usage:
#   powershell -ExecutionPolicy Bypass -File scripts\link-skills.ps1
#   powershell -ExecutionPolicy Bypass -File scripts\link-skills.ps1 -Remove
#
# The single collection junction makes edits, additions, and removals visible without relinking.
# Existing third-party skills in ~/.agents/skills are left untouched.

param(
    [switch]$Remove
)

$ErrorActionPreference = "Stop"
$RepoRoot = [System.IO.Path]::GetFullPath((Split-Path -Parent $PSScriptRoot))
$Base = Join-Path $env:USERPROFILE ".agents\skills"
$Link = Join-Path $Base "myskills"

function Get-NormalizedTarget([System.IO.FileSystemInfo]$Item) {
    if (-not $Item.Target) { return $null }
    return [System.IO.Path]::GetFullPath([string]($Item.Target | Select-Object -First 1)).TrimEnd('\')
}

function Test-OwnedJunction([string]$Path, [string]$Target) {
    if (-not (Test-Path -LiteralPath $Path)) { return $false }
    $item = Get-Item -LiteralPath $Path
    return $item.LinkType -eq "Junction" -and
        (Get-NormalizedTarget $item) -eq ([System.IO.Path]::GetFullPath($Target).TrimEnd('\'))
}

function Remove-LegacyLinks {
    Get-ChildItem -Directory $RepoRoot | Where-Object {
        Test-Path -LiteralPath (Join-Path $_.FullName "SKILL.md")
    } | ForEach-Object {
        $legacy = Join-Path $Base $_.Name
        if (Test-OwnedJunction $legacy $_.FullName) {
            [System.IO.Directory]::Delete($legacy, $false)
            Write-Host "removed legacy $legacy"
        }
    }
}

if (-not (Test-Path -LiteralPath $Base)) {
    New-Item -ItemType Directory -Path $Base | Out-Null
}

if ($Remove) {
    if (Test-OwnedJunction $Link $RepoRoot) {
        [System.IO.Directory]::Delete($Link, $false)
        Write-Host "removed  $Link"
    } elseif (Test-Path -LiteralPath $Link) {
        Write-Warning "skip     $Link is not a junction owned by this repository"
    }
    Remove-LegacyLinks
    exit 0
}

if (Test-Path -LiteralPath $Link) {
    if (-not (Test-OwnedJunction $Link $RepoRoot)) {
        Write-Error "$Link already exists and is not a junction owned by this repository"
    }
    Write-Host "ok       $Link"
} else {
    New-Item -ItemType Junction -Path $Link -Target $RepoRoot | Out-Null
    Write-Host "mounted  $Link -> $RepoRoot"
}

# Migrate links created by older releases only after the collection mount is valid.
Remove-LegacyLinks
