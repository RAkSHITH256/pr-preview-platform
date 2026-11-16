# First PR Preview Test

This pull request validates the PR Preview Platform CI pipeline.

Expected workflow:
1. GitHub Actions checks out the PR source.
2. Docker builds the preview application.
3. The image is published to GitHub Container Registry.
4. The image receives a PR-specific tag and a commit SHA tag.
