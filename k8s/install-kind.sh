#!/usr/bin/env bash

set -euo pipefail

# Check if kind is already installed
if command -v kind &> /dev/null; then
  echo "kind is already installed"
else
  # Check OS type (Linux or macOS), and set environment variable for kind binary
  if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    KIND_BINARY_URL="https://kind.sigs.k8s.io/dl/v0.17.0/kind-linux-amd64"
  elif [[ "$OSTYPE" == "darwin"* ]]; then
    MACHINE_TYPE=$(uname -m)
    # For Intel Macs (amd64)
    if [[ "$MACHINE_TYPE" == "x86_64" ]]; then
      KIND_BINARY_URL="https://kind.sigs.k8s.io/dl/v0.17.0/kind-darwin-amd64"
    # For M1/M2 Macs (arm64)
    elif [[ "$MACHINE_TYPE" == "arm64" ]]; then
      KIND_BINARY_URL="https://kind.sigs.k8s.io/dl/v0.17.0/kind-darwin-arm64"
    else
      echo "Unsupported machine type: $MACHINE_TYPE"
      exit 1
    fi
  else
    echo "Unsupported OS: $OSTYPE"
    exit 1
  fi
  # Download kind binary, and place it into a $PATH directory
  curl -Lo ./kind $KIND_BINARY_URL
  chmod +x ./kind
  sudo mv ./kind /usr/local/bin/kind
  echo "kind installed successfully"
fi

# Create a local Kubernetes cluster
kind create cluster --config kind-config.yml

# Get kubeconfig
kubectl cluster-info --context kind-kind
