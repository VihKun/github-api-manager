# GitHub API Manager

A Go-based API to manage GitHub repositories and pull requests using the GitHub REST API, with configurations for deployment to Kubernetes.

## Features

- List GitHub repositories.
- Create and delete repositories.
- List pull requests for a given repository.
- Health check endpoint.
- Deploy the application to a Kubernetes cluster (Minikube).

## Prerequisites

Before running this project locally or deploying it, ensure you have the following:

- **Go (1.23.4)**: This project is written in Go and requires Go 1.23.4 or later.
- **Docker**: Required for building and pushing the Docker image to a container registry.
- **Minikube**: For local Kubernetes clusters (used for deploying the application).
- **kubectl**: The Kubernetes command-line tool to interact with your cluster.
- **GitHub Access Token (GITHUB_TOKEN)**: Required to authenticate with the GitHub API.

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/VihKun/github-api-manager.git
cd github-api-manager
```

### 2. Set up Go environment

Make sure that Go 1.23.4 is installed. You can download it from [Go's official website.](https://golang.org/dl/)

### 3. Install dependencies

Run the following commands to install the dependencies:

```bash
go mod tidy
go mod vendor
```

### 4. Set up environment variables

You need to set the GIT_TOKEN environment variable for authentication with GitHub. You can obtain this token from GitHub by following these steps:

- Go to GitHub tokens page.
- Generate a new personal access token with the appropriate permissions for repository management and pull requests.

Set the GIT_TOKEN in your environment:

```bash
export GIT_TOKEN=<your-github-token>
```

### 5. Build the Docker image

Build the Docker image for the application:

```bash
docker build -t github-api-manager .
```

### 6. Start Minikube

Ensure that you have Minikube installed and running. Start Minikube:

```bash
minikube start
```

### 7. Apply Kubernetes configurations

Once Minikube is running, apply the Kubernetes configurations for deployment:

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

### 8. Expose the API

You can expose the API service using the following command:

```bash
kubectl port-forward deployment/github-api-manager 8080:8080
```

In this example, it will forward the service to localhost:8080.

## API Endpoints

### 1. /api/v1/health

- **Method**: GET
- **Description**: Returns the health status of the API.

Example:

```bash
curl -X GET localhost:8080/api/v1/health
```

### 2. /api/v1/repos

- **Method**: GET
- **Description**: Lists all repositories for the authenticated user (associated to the token used).

Example:

```bash
curl -X GET localhost:8080/api/v1/repos
```

### 3. /api/v1/repos

- **Method**: POST
- **Description**: Creates a new repository.

Example:

```bash
curl -X POST http://localhost:8080/api/v1/repos \
    -H "Content-Type: application/json" \
    -d '{"name": "name_of_repo", "description": "description_of_repo", "private": true}'
```

### 4. /api/v1/repos/{name}/pulls

- **Method**: GET
- **Description**: Lists pull requests for a specific repository.

Example:

```bash
curl -X GET localhost:8080/api/v1/repos/github-api-manager/pulls
```

### 5. /api/v1/repos/{name}

- **Method**: DELETE
- **Description**: Deletes a specified repository.

Example:

```bash
curl -X DELETE localhost:8080/api/v1/repos/github-api-manager
```

## Testing

To run tests locally:

```bash
go test ./...
```

## Linting

To run linters:

```bash
golint ./...
```

(**golint needs to be installed**)

## CI/CD Pipeline

The project includes a GitHub Actions CI/CD pipeline, which performs the following tasks:

- Runs tests.
- Lints the code.
- Runs security checks.
- Builds a Docker image.
- Deploys to Minikube.

The CI/CD pipeline is triggered on each push to the main branch or when a pull request is created.
Secrets Required for GitHub Actions:

- **DOCKER_USERNAME**: Docker Hub username.
- **DOCKER_PASSWORD**: Docker Hub password.
- **KUBECONFIG**: The kubeconfig file for accessing the Kubernetes cluster.
- **GIT_TOKEN**: GitHub personal access token.

Everything related to the pipeline can be modified using the `.github/workflows/ci-cd.yml` file.

For testing purposes, a local runner can be setup. Refer to the official GitHub documentation for setting up and configuring a self-hosted runner for GitHub Actions.

## Troubleshooting

- If you encounter issues related to kubectl or Minikube during deployment, ensure your kubectl is properly configured to point to the correct Minikube cluster.
- Ensure that the necessary environment variables (Ex.: **GIT_TOKEN**) are set in your environment or GitHub Secrets.