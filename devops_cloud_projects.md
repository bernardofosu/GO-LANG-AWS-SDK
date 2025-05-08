# 🚀 DevOps Projects (Automation, CI/CD, Infra as Code)

## 1. AWS EC2 Provisioner CLI
**What it does**: Launch and manage EC2 instances via CLI.

**Packages**:
- `github.com/aws/aws-sdk-go-v2`
- `cobra` (for CLI)
- `viper` (for config)

## 2. S3 Uploader/Downloader Tool
**What it does**: Upload/download files from S3 buckets.

**Packages**:
- `github.com/aws/aws-sdk-go-v2/service/s3`
- `mime`, `os`, `path/filepath`

## 3. Docker Image Scanner
**What it does**: Scan Docker images for vulnerabilities using an API (e.g., Trivy or Clair).

**Packages**:
- `net/http`
- `encoding/json`
- Trivy CLI or REST API wrapper

## 4. Kubernetes Resource Monitor
**What it does**: Watch CPU/memory usage of pods, auto-scale or alert.

**Packages**:
- `k8s.io/client-go`
- `go-resty/resty/v2` (for APIs)
- `prometheus/client_golang` (if integrating metrics)

## 5. GitOps Bot
**What it does**: Monitors Git repos and triggers deployments (like ArgoCD-lite).

**Packages**:
- `go-git`
- `fsnotify`
- `net/http`

## 6. Terraform Wrapper Tool
**What it does**: Build a CLI that wraps around terraform to enforce policies or naming standards.

**Packages**:
- `os/exec`
- `bufio`
- `regexp`

# 🔐 Cloud Security Projects

## 1. IAM Policy Analyzer
**What it does**: Audits AWS IAM policies for privilege escalation risks.

**Packages**:
- `github.com/aws/aws-sdk-go-v2/service/iam`
- `encoding/json`
- `strings`

## 2. CloudTrail Log Analyzer
**What it does**: Parse CloudTrail logs and detect suspicious actions (e.g., CreateAccessKey in prod).

**Packages**:
- `github.com/aws/aws-sdk-go-v2/service/s3`
- `encoding/json`, `bufio`

## 3. Secrets Scanner
**What it does**: Scan local file systems or Git repos for secrets.

**Packages**:
- `go-git`
- `regexp`
- `filepath`, `os`

## 4. Kubernetes RBAC Visualizer
**What it does**: Parse Roles/ClusterRoles and generate a visual map.

**Packages**:
- `k8s.io/client-go`
- `gonum/graph` or output to Graphviz

## 5. JWT Analyzer CLI
**What it does**: Decode and analyze JWT tokens (e.g., for Auth0, AWS Cognito).

**Packages**:
- `github.com/golang-jwt/jwt/v5`
- `encoding/base64`, `crypto`

# 🛠 Tooling Recommendations

| Purpose         | Package/Tool                  |
|-----------------|-------------------------------|
| CLI Apps        | cobra, urfave/cli             |
| Configs         | viper, envconfig              |
| HTTP Clients    | net/http, resty               |
| AWS SDK         | aws-sdk-go-v2                 |
| GitOps          | go-git, fsnotify              |
| Kubernetes      | k8s.io/client-go, controller-runtime |
| Logging         | logrus, zap                   |
| Testing         | testify, httptest             |
| Security        | golang-jwt, crypto/tls        |
