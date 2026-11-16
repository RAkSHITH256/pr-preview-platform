# Second PR Preview Test

Purpose: Validate a second independent pull request preview.

Checks:
- GitHub Actions builds and publishes the PR image.
- Argo CD ApplicationSet detects the open PR.
- A separate preview namespace is created.
- PR number and commit SHA are injected.
- Application endpoints return HTTP 200.
- Closing the PR removes its preview resources.
