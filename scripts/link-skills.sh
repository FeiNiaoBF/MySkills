#!/usr/bin/env bash
# link-skills.sh — 把本仓库的每个 skill 目录以符号链接挂到 agent skills 目录（Linux / macOS / WSL）
# 用法:
#   ./scripts/link-skills.sh frontend-guide              # 链接指定 skill（推荐：做好一个链一个）
#   ./scripts/link-skills.sh --remove frontend-guide     # 移除指定 skill 的链接
#   ./scripts/link-skills.sh                             # 链接全部 skill
#
# 换机器 clone 后按需跑；符号链接不进 git。
# 已存在但不是指向本仓库的同名目录会被跳过（不覆盖手动部署的版本）。
# 链接使用绝对路径，与 Windows 侧的 Junction 行为一致：仓库搬家后重跑本脚本即可修复。

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

REMOVE=0
declare -a ONLY=()
for arg in "$@"; do
    case "$arg" in
        --remove|-r) REMOVE=1 ;;
        -h|--help) sed -n '2,7p' "${BASH_SOURCE[0]}"; exit 0 ;;
        *) ONLY+=("$arg") ;;
    esac
done

LINK_BASES=(
    "$HOME/.agents/skills"
    "$HOME/.pi/agent/skills"
)

# 发现 skill：仓库根下含 SKILL.md 的一级目录（给了位置参数则只取同名目录）
skills=()
for dir in "$REPO_ROOT"/*/; do
    name="$(basename "$dir")"
    [ -f "$dir/SKILL.md" ] || continue
    if [ ${#ONLY[@]} -gt 0 ]; then
        keep=0
        for want in "${ONLY[@]}"; do
            [ "$want" = "$name" ] && keep=1
        done
        [ "$keep" = 1 ] || continue
    fi
    skills+=("$(basename "$dir")")
done

if [ ${#skills[@]} -eq 0 ]; then
    echo "WARN: 未在 $REPO_ROOT 下找到匹配的 skill 目录（需含 SKILL.md）" >&2
    exit 1
fi

for base in "${LINK_BASES[@]}"; do
    mkdir -p "$base"

    for name in "${skills[@]}"; do
        target="$REPO_ROOT/$name"
        link="$base/$name"

        if [ "$REMOVE" = 1 ]; then
            if [ -L "$link" ] && [ "$(readlink "$link")" = "$target" ]; then
                rm "$link"
                echo "removed  $link"
            fi
            continue
        fi

        if [ -e "$link" ] || [ -L "$link" ]; then
            if [ -L "$link" ] && [ "$(readlink "$link")" = "$target" ]; then
                echo "ok       $link"
            else
                echo "SKIP     $link 已存在且不是指向本仓库的符号链接，请手动处理" >&2
            fi
        else
            ln -s "$target" "$link"
            echo "linked   $link -> $target"
        fi
    done
done
