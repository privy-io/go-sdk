#!/usr/bin/env bash
# Prints the markdown body of the sticky `release-proposal` comment for a release PR: the
# source SHA release-please ran on, the last published tag, and a compare link between them.
#
# Read-only: everything comes from the GitHub API, so no full clone is needed. The caller
# posts the output (e.g. marocchino/sticky-pull-request-comment with header
# `release-proposal`).
#
# Runs in the release-please workflow: the source SHA is $GITHUB_SHA, and GH_TOKEN needs
# read access to contents.
#
# Usage: proposal-comment.sh > body.md
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=SCRIPTDIR/lib.sh
source "$script_dir/lib.sh"

repo_url="$GITHUB_SERVER_URL/$GITHUB_REPOSITORY"
short_sha="${GITHUB_SHA:0:7}"

if last_version="$(manifest_version "$GITHUB_SHA")"; then
  last_tag="v$last_version"
  last_tag_cell="[\`$last_tag\`]($repo_url/releases/tag/$last_tag)"
  changes_cell="[\`$last_tag...$short_sha\`]($repo_url/compare/$last_tag...$GITHUB_SHA)"
else
  last_tag_cell="_none found in ${MANIFEST_PATH}_"
  changes_cell="_n/a_"
fi

cat <<MD
### Release proposal

| | |
|---|---|
| Source SHA | [\`$short_sha\`]($repo_url/commit/$GITHUB_SHA) |
| Last published tag | $last_tag_cell |
| Changes | $changes_cell |
MD

if [[ -n "${GITHUB_RUN_ID:-}" ]]; then
  echo
  echo "<sub>Updated by [release-please run $GITHUB_RUN_ID]($repo_url/actions/runs/$GITHUB_RUN_ID).</sub>"
fi
