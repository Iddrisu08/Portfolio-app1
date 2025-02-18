[![CI/CD](https://github.com/Iddrisu08/Portfolio-app1/actions/workflows/ci.yaml/badge.svg)](https://github.com/Iddrisu08/Portfolio-app1/actions/workflows/ci.yaml)


# Portfolio Website

A personal portfolio website developed using Go, containerized with Docker, and deployed on AWS EKS with a robust CI/CD pipeline.

### Table of Contents

```Overview
Features
Technologies Used
Installation
Usage
Project Structure
Contributing
License
```

### Overview

This project showcases a personal portfolio website built with Go. The application is containerized using Docker and deployed on an AWS EKS cluster. The deployment process is automated through a CI/CD pipeline implemented with GitHub Actions, ensuring seamless integration of new features and maintaining high availability and fault tolerance through Kubernetes capabilities like autoscaling, load-balancing, and self-healing.

### Features

Personal Portfolio: Displays personal projects, skills, and contact information.
Containerization: Application is containerized using Docker for consistent environments.
CI/CD Pipeline: Automated deployment pipeline using GitHub Actions for continuous integration and deployment.
Kubernetes Deployment: Deployed on AWS EKS with features like autoscaling, load-balancing, and self-healing for high availability.

### Technologies Used

Programming Language: Go
Containerization: Docker
Orchestration: Kubernetes (AWS EKS)
CI/CD: GitHub Actions
Cloud Provider: Amazon Web Services (AWS)

### Installation

To set up the project locally:

Clone the Repository:

`git clone https://github.com/Iddrisu08/Portfolio-app1.git`

`cd Portfolio-app1`

Build the Docker Image:

`docker build -t portfolio-app:latest .`

Run the Docker Container:

`docker run -p 8080:8080 portfolio-app:latest`

The application will be accessible at `http://localhost:8080.`

### Usage

For local development:

Start the Application:

`go run main.go`

Access the Application: Open your browser and navigate to `http://localhost:8080` to view the portfolio website.

### Project Structure

```Portfolio-app1/
├── Dockerfile
├── README.md
├── main.go
├── go.mod
├── go.sum
└── .github/
    └── workflows/
        └── ci-cd.yml
        
main.go: The main application file written in Go.
Dockerfile: Instructions to build the Docker image.
.github/workflows/ci-cd.yml: GitHub Actions workflow for the CI/CD pipeline.
```

### Contributing

Contributions are welcome! To contribute:

Fork the repository.
Create a new branch `(git checkout -b feature-branch).`

Make your changes and commit them `(git commit -m 'Add new feature')`.

Push to the branch `(git push origin feature-branch)`.

Open a Pull Request.


# License

default copyright applies.

