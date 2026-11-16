# PR Preview Platform

**Automated ephemeral Kubernetes environments for every GitHub Pull Request.**

A GitOps-driven platform that automatically builds, deploys, and manages isolated preview environments for pull requests. Each PR gets its own Kubernetes namespace and application deployment. When the PR closes, its preview environment is automatically removed.

## Architecture

```mermaid
flowchart TD
    A[Developer opens Pull Request] --> B[GitHub Actions]
    B --> C[Build Docker Image]
    C --> D[Push PR Image to GHCR]
    D --> E[Argo CD ApplicationSet]
    E --> F[GitHub Pull Request Generator]
    F --> G[Generate preview-N Application]
    G --> H[Kustomize]
    H --> I[Create preview-N Namespace]
    I --> J[Deploy Application and Service]
    J --> K[Preview Environment Running]

    L[Pull Request Closed] --> M[ApplicationSet Detects Closure]
    M --> N[Delete Generated Application]
    N --> O[Argo CD Cascading Cleanup]
    O --> P[Delete Preview Namespace]
