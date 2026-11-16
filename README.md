# PR Preview Platform

**Automated ephemeral Kubernetes environments for every GitHub Pull Request.**

A GitOps-driven platform that automatically builds, deploys, and manages isolated preview environments for pull requests. Each PR gets its own Kubernetes namespace and application deployment. When the PR closes, its preview environment is automatically removed.

## Architecture

```mermaid
flowchart TD
    A[Developer opens Pull Request] --> B[GitHub Actions]
    B --> C[Build Docker Image]
    C --> D[Push Image to GHCR]
    D --> E[Argo CD ApplicationSet]
    E --> F[GitHub PR Generator]
    F --> G[Generate preview-N Application]
    G --> H[Kustomize]
    H --> I[Create preview-N Namespace]
    I --> J[Deploy Application and Service]
    J --> K[Preview Environment Running]

    L[Pull Request Closed] --> M[ApplicationSet Detects Closure]
    M --> N[Delete Generated Application]
    N --> O[Argo CD Cleanup]
    O --> P[Delete Preview Namespace]
```

## Screenshots

### GitHub Actions – Image Build
![GitHub Actions Build](docs/screenshots/github-actions-build.png)

### Argo CD – Preview Application
![Argo CD Preview](docs/screenshots/argocd-preview.png)

### Running Preview Environment
![Preview Application](docs/screenshots/preview-application.png)

### Automatic Cleanup
![Automatic Cleanup](docs/screenshots/automatic-cleanup.png)

## Technology Stack

- **Containerization:** Docker
- **Orchestration:** Kubernetes, kind
- **GitOps:** Argo CD, ApplicationSet
- **CI/CD:** GitHub Actions
- **Container Registry:** GitHub Container Registry (GHCR)
- **Configuration Management:** Kustomize
- **Application:** Go

## Key Features

- Automated Docker image builds for pull requests.
- Isolated Kubernetes namespace for every open PR.
- Dynamic preview application generation using Argo CD ApplicationSet.
- Automated synchronization and deployment.
- Automatic environment cleanup when a PR closes.
- Health and version endpoints for deployment verification.

## Validation

Successfully tested two complete PR preview lifecycles.

| Validation | Result |
|---|---|
| GitHub Actions image build | Passed |
| Preview deployment for PR #1 | Passed |
| Preview deployment for PR #2 | Passed |
| Application health and version endpoints | Passed |
| Automatic namespace cleanup | Passed |
| Existing GitOps application preserved | Passed |

Detailed evidence: [`docs/validation-evidence.md`](docs/validation-evidence.md).

## Project Structure

```text
pr-preview-platform/
├── .github/workflows/
├── app/
├── argocd/
├── k8s/
│   ├── base/
│   └── preview/
├── docs/
│   ├── screenshots/
│   └── validation-evidence.md
├── Dockerfile
└── README.md
```

## Project Outcome

Built and validated an automated PR preview platform using Kubernetes and GitOps, demonstrating dynamic environment provisioning, continuous deployment, and automatic resource cleanup.
