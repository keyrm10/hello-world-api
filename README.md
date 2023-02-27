# Hello World API

This repository contains code for a simple "Hello World" REST API that returns a greeting based on a customer ID. It includes a Dockerfile for building the API into a Docker image and Kubernetes manifest files for running the API in a local Kubernetes cluster using Kind.

It also includes a GitHub Actions workflow that builds the Docker image and pushes it to Docker Hub when changes are pushed to the main branch.

## Overview

The hello-world-api is a simple HTTP server written in Go that returns a greeting in JSON format based on a provided customer ID. The available greetings are stored in a map with the keys "a", "b", and "c". If the provided customer ID does not match any of these keys, the greeting "Hello" is returned.

To use the API, you can send an HTTP GET request to the endpoint `/api/v1/hello/<id>`, where `<id>` is the desired customer ID. For example, to get a greeting for customer "a", you would send a request to `/api/v1/hello/a`. The `<id>` parameter is case-insensitive, so you can use upper or lower case strings.

The API will return a JSON object with the keys "id" and "salutation", where "id" is the provided customer ID and "salutation" is the greeting associated with that ID.

##### Request Parameters

-   `id` (string, required) - The `customerID` parameter.

##### Response Format

-   `id` (string) - The `customerID` parameter.
-   `salutation` (string) - The greeting message based on the `customerID` parameter.

If the customer ID was not provided, the API will return a HTTP 400 Bad Request error with a message "Missing id parameter".

## GitHub Actions workflow

The GitHub Actions workflow defined in `.github/workflows/main.yml` builds the Docker image and pushes it to Docker Hub when changes are pushed to the main branch. The workflow is triggered by the push event on the main branch.

The workflow includes the following steps:

1. Checkout: Checks out the repository code.
2. Login to Docker Hub: Logs in to Docker Hub using the `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN` secrets.
3. Set up Docker Buildx: Sets up Docker Buildx for multi-architecture builds.
4. Build and push: Builds the Docker image using the Dockerfile in the repository root, pushes the image to Docker Hub, and tags it with `DOCKERHUB_USERNAME` and latest.
5. Inspect: Inspects the Docker image to verify that it was built correctly.

## Usage

### Requirements

- [cURL](https://everything.curl.dev/get)
- [Docker](https://docs.docker.com/get-docker/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)

### Building and running the API locally

To build and run the API locally, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/keyrm10/hello-world-api.git
    ```

2. Navigate to the cloned directory:

    ```bash
    cd hello-world-api/
    ```
3. Build the Docker image:

    ```bash
    docker build -t hello-world-api .
    ```

4. Run the Docker image:
    ```bash
    docker run -dp 80:8080 hello-world-api
    ```

### Running the API in a local Kubernetes cluster

To run the API in a local Kubernetes cluster using Kind, follow these steps:

1. Install Kind and create a local cluster:

    ```bash
    ./k8s/install-kind.sh
    ```

2. Deploy the API to the Kubernetes cluster

    ```bash
    kubectl apply -f ./k8s/manifests/
    ```

> After applying the manifests, it will take some time for the Ingress to fully deploy and be functional. You can check the status of the Ingress by running the command `kubectl get ingress hello`.

### Testing the API

You can access the API at `http://localhost/api/v1/hello/<id>`, where `<id>` is a string representing the customer ID. Here is an example of how to use the API with the command-line tool cURL:

```bash
curl http://localhost/hello/A
```

This will return a JSON object like this:

```json
{"id":"A","salutation":"Hi"}
```

## Clean up

If you have built the Docker image locally and run a container, you can stop and delete the container using the following command:

```bash
docker rm -f CONTAINER
```

You can also delete the Docker image by running:

```bash
docker image rm IMAGE
```

Finally, if you installed Kind and created a local Kubernetes cluster, you can delete them by running this script:

```bash
./k8s/uninstall-kind.sh
```
