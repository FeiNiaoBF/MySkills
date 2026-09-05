# link-skills.ps1 — 把本仓库的每个 skill 目录以 Junction 链接到 agent skills 目录
# 用法:
#   powershell -ExecutionPolicy Bypass -File scripts\link-skills.ps1 -Only frontend-guide          # 链接指定 skill（推荐：做好一个链一个）
#   powershell -ExecutionPolicy Bypass -File scripts\link-skills.ps1 -Remove -Only frontend-guide   # 移除指定 skill 的链接
#   powershell -ExecutionPolicy Bypass -File scripts\link-skills.ps1            # 链接全部 skill
#
# 换机器 clone 后按需跑；Junction 是本机对象，不入 git。
# 已存在但不是指向本仓库 Junction 的同名目录会被跳过（不覆盖手动部署的版本）。

param(
    [switch]$Remove,
    [string[]]$Only
)

$ErrorActionPreference = "Stop"
$RepoRoot = Split-Path -Parent $PSScriptRoot
$Links = @(
    (Join-Path $env:USERPROFILE ".agents\skills"),
    (Join-Path $env:USERPROFILE ".pi\agent\skills")
)

# 发现 skill：仓库根下含 SKILL.md 的一级目录
$skills = Get-ChildItem -Directory $RepoRoot | Where-Object {
    (Test-Path (Join-Path $_.FullName "SKILL.md")) -and
    (-not $Only -or $Only -contains $_.Name)
}

if (-not $skills) {
    Write-Warning "未在 $RepoRoot 下找到任何含 SKILL.md 的目录"
    exit 1
}

foreach ($base in $Links) {
    if (-not (Test-Path $base)) { New-Item -ItemType Directory -Path $base | Out-Null }

    foreach ($skill in $skills) {
        $link = Join-Path $base $skill.Name

        if ($Remove) {
            if ((Test-Path $link) -and (Get-Item $link).LinkType -eq "Junction") {
                # Junction 用 Remove-Item 只拆链接、不递归删目标内容
                [System.IO.Directory]::Delete($link, $false)
                Write-Host "removed  $link"
            }
            continue
        }

        if (Test-Path $link) {
            $item = Get-Item $link
            if ($item.LinkType -eq "Junction" -and $item.Target -eq $skill.FullName) {
                Write-Host "ok       $link"
            } else {
                Write-Warning "skip     $link 已存在且不是指向本仓库的 Junction，请手动处理"
            }
        } else {
            New-Item -ItemType Junction -Path $link -Target $skill.FullName | Out-Null
            Write-Host "linked   $link -> $($skill.FullName)"
        }
    }
}
