# shellcheck shell=bash
# Shared helpers for the release scripts in this directory. Source it; don't run it.

MANIFEST_PATH=.release-please-manifest.json

# manifest_version <ref>
# Prints the version recorded in the release-please manifest at <ref> in this repo, read
# through the contents API (no checkout needed).
manifest_version() {
  gh api -H "Accept: application/vnd.github.raw+json" \
    "repos/$GITHUB_REPOSITORY/contents/$MANIFEST_PATH?ref=$1" |
    jq -er '.["."] // empty'
}
