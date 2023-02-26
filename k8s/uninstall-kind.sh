#!/usr/bin/env bash

set -euo pipefail

kind delete cluster --name=kind

echo "kind cluster deleted successfully"

rm /usr/local/bin/kind

echo "kind uninstalled successfully"