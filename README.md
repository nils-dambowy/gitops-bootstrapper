# GitOps Bootstrapper

A Zero-Touch CLI tool to automate the application lifecycle and provide "Golden Paths" for developers.

## The Problem

Manually creating repositories, writing Kubernetes manifests, configuring ArgoCD applications, and setting up HashiCorp Vault secrets for every new microservice is error-prone and doesn't scale.

## The Solution

`gitops-bootstrapper` automates the boilerplate. With a single command, it provisions the complete infrastructure-as-code repository for a new application.

## Planned Features

- [ ] **GitLab Integration:** Automatically create a new repository with a base skeleton (Containerfile, CI/CD pipeline).
- [ ] **GitOps / ArgoCD:** Generate K8s base manifests (Deployment, Service, NetworkPolicies - default deny!) and push them to the central GitOps repository.
- [ ] **Secret Management:** Interface with HashiCorp Vault to create AppRoles and inject secret references.
- [ ] **Observability:** Scaffold default Prometheus `ServiceMonitor` and basic alerts.

## Getting Started

### Prerequisites

- Go 1.21+

### Running locally

For now, the project is in early development. To run the CLI:

```bash
go run .
```
