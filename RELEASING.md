# Releasing

Releases are cut by [release-please](https://github.com/googleapis/release-please) in the
private `privy-io/go-sdk-private` repo, then promoted to the public
[`privy-io/go-sdk`](https://github.com/privy-io/go-sdk) repo, where the Go module proxy picks up the tag.
Tags are `vX.Y.Z`.

Public `main` only moves when a release is promoted, so it always matches the latest stable
release.

## Stable releases (from `main`)

1. 👤 Merge changes into `main` in `go-sdk-private`.
2. 🤖 On every push to `main`, release-please opens or updates the release PR (`release: X.Y.Z`)
   in `go-sdk-private`.
3. 👤 Approve and merge the release PR.
4. 🤖 CI runs the [release workflow](./.github/workflows/release-please.yml), which tags
   `vX.Y.Z` and creates the GitHub Release in `go-sdk-private`.
5. 🤖 The release triggers the [promote workflow](./.github/workflows/stlc-promote.yml), which:
   - fast-forwards `main` in `go-sdk` to the tag commit;
   - pushes the tag to `go-sdk`;
   - creates the matching GitHub Release in `go-sdk`.
6. 👤 Update the changelog in [our public docs](https://docs.privy.io).

The public release notes link to `go-sdk` commits and leave out PR numbers, since those
refer to private PRs.
