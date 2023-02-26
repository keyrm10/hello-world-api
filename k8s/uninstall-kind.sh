#!/usr/bin/env bash

set -euo pipefail

kind delete cluster --name=kind

rm /usr/local/bin/kind

echo "kind uninstalled successfully"