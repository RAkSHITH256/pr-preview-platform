# PR Preview Platform — Validation Evidence

**Validation date:** 16 November 2026
**Test PR:** #1 — `test: validate first PR preview build`
**Test commit:** `b6640a3b01bb16a39fac8ad07f3b73d3c8e97f1e`

## 1. Preview Environment Creation

Verified that GitHub Actions built and published the PR image to GHCR.

Argo CD ApplicationSet detected PR #1 and created:

- Application: `preview-1`
- Namespace: `preview-1`
- Deployment: `preview-app`
- Service: `preview-app`

Application status: Synced / Healthy.

## 2. Dynamic PR Metadata

Verified that ApplicationSet and Kustomize injected:

```text
APP_ENV=preview
PR_NUMBER=1
GIT_SHA=b6640a3b01bb16a39fac8ad07f3b73d3c8e97f1e
