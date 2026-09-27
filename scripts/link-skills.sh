#!/usr/bin/env bash
# link-skills.sh — mount this repository once under the shared Agent Skills root.
# Usage:
#   bash scripts/link-skills.sh
#   bash scripts/link-skills.sh --remove
#
# The single collection symlink makes edits, additions, and removals visible without relinking.
# Existing third-party skills in ~/.agents/skills are left untouched.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd -P)"
BASE="$HOME/.agents/skills"
LINK="$BASE/myskills"
REMOVE=0

case "${1:-}" in
    --remove|-r) REMOVE=1 ;;
    "") ;;
    -h|--help) sed -n '2,7p' "${BASH_SOURCE[0]}"; exit 0 ;;
    *) echo "ERROR: unknown argument: $1" >&2; exit 2 ;;
esac

# MSYS/Git Bash may emulate `ln -s` by copying a directory. Delegate to the
# Windows implementation so the live mount remains a real Junction.
case "$(uname -s)" in
    MINGW*|MSYS*|CYGWIN*)
        args=()
        [ "$REMOVE" = 1 ] && args+=("-Remove")
        powershell.exe -NoProfile -ExecutionPolicy Bypass -File \
            "$(cygpath -w "$SCRIPT_DIR/link-skills.ps1")" "${args[@]}"
        exit $?
        ;;
esac

owned_link() {
    [ -L "$1" ] && [ "$(readlink -f "$1")" = "$(readlink -f "$2")" ]
}

remove_legacy_links() {
    local dir name legacy
    for dir in "$REPO_ROOT"/*/; do
        [ -f "$dir/SKILL.md" ] || continue
        name="$(basename "$dir")"
        legacy="$BASE/$name"
        if owned_link "$legacy" "$dir"; then
            rm "$legacy"
            echo "removed legacy $legacy"
        fi
    done
}

mkdir -p "$BASE"

if [ "$REMOVE" = 1 ]; then
    if owned_link "$LINK" "$REPO_ROOT"; then
        rm "$LINK"
        echo "removed  $LINK"
    elif [ -e "$LINK" ] || [ -L "$LINK" ]; then
        echo "WARN: skip $LINK; it is not a symlink owned by this repository" >&2
    fi
    remove_legacy_links
    exit 0
fi

if [ -e "$LINK" ] || [ -L "$LINK" ]; then
    if ! owned_link "$LINK" "$REPO_ROOT"; then
        echo "ERROR: $LINK already exists and is not a symlink owned by this repository" >&2
        exit 1
    fi
    echo "ok       $LINK"
else
    ln -s "$REPO_ROOT" "$LINK"
    echo "mounted  $LINK -> $REPO_ROOT"
fi

# Migrate links created by older releases only after the collection mount is valid.
remove_legacy_links
