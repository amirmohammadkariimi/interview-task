# Interview Task

This guide explains how to build, run, lint, and deploy the application, as well as configure environment variables for development and deployment.

## Building the Application

### On Linux
To build the application on a Linux system, use the following command:
- `go build ./cmd`

---

### Using Docker
To build the application using Docker, run:
- `docker build -t interview_task`

---

## Running the Linter

To run the linter and ensure code quality, use `golangci-lint` with the following command:
- `golangci-lint run ./...`

---

## Kubernetes Deployment

### Manual Deployment

To manually deploy the application to a Kubernetes cluster, follow these steps:

1. Apply the secret configuration:
`kubectl apply -f k8s/secret.yaml -n <namespace>`
2. Deploy the application using Helm:
`helm upgrade --install interview-task -n <namespace> -f interview_task/values.yaml ./interview_task`

---

## Environment Variables

Configure the following environment variables for both local development and Kubernetes deployments:

| Variable           | Description                                         | Example                     |
|--------------------|-----------------------------------------------------|-----------------------------|
| `PORT`             | Port for the webserver (as a string)                | `":3000"`                   |
| `DATABASE_NAME`     | Name of the database                                | `"my_database"`             |
| `DATABASE_USER`     | Username for the database                           | `"db_user"`                 |
| `DATABASE_PASSWORD` | Password for the database user                      | `"password123"`             |
| `DATABASE_ADDRESS`  | Address for the database connection                 | `"tcp(127.0.0.1:3306)"`     |

---

### Notes:
- Ensure that the `secret.yaml` file contains the appropriate database credentials and other sensitive information before deploying to Kubernetes.
- Modify the `<namespace>` placeholder with the actual namespace where you want to deploy the application.
