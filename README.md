# User Profile Service

The **User Profile Service** is a RESTful API built in Go that allows users to manage their profiles. It supports CRUD operations and is containerized using Docker and deployed on Kubernetes. This service is designed to be scalable and easy to integrate with other microservices.

### Features:
- Create, retrieve, update, and delete user profiles.
- Kubernetes deployment with scalability.
- Integration with SQLite or PostgreSQL.
- Extensible architecture for adding new features.

## Technologies Used

- **Programming Language**: Go (Golang)
- **Web Framework**: net/http
- **Database**: SQLite (local) / PostgreSQL (production)
- **Containerization**: Docker
- **Orchestration**: Kubernetes (Minikube)


## Setup Instructions

### Prerequisites
- Install [Go](https://golang.org/doc/install) (1.20+)
- Install [Docker](https://www.docker.com/)
- Install [Minikube](https://minikube.sigs.k8s.io/docs/start/)
- Install `kubectl`

### Clone the Repository

git clone https://github.com/anishhgoel/user_profile_service.git
cd user-profile-service

### Build and Run Locally
go build -o main .
./main


### Run with Docker

1. Build the Docker image:  docker build -t user-profile-service:latest .
2. Run the container: docker run -p 8080:8080 user-profile-service:latest

### Deploy on Kubernetes

1. Start Minikube: minikube start
2. Enable Docker for Minikube: eval $(minikube docker-env)
3. Apply Kubernetes manifests: kubectl apply -f deployment.yaml   
kubectl apply -f service.yaml
4. Access the service: minikube service user-profile-service




