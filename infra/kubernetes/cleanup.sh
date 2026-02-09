#!/bin/bash
# Delete all Requiem API resources from Kubernetes
# Usage: ./cleanup.sh [namespace]

set -e

NAMESPACE=${1:-requiem}
KUBECTL="kubectl"

echo "🗑️  Cleaning up Requiem API from Kubernetes namespace: $NAMESPACE"
echo ""

read -p "Are you sure you want to delete all resources in namespace $NAMESPACE? (yes/no): " CONFIRM
if [ "$CONFIRM" != "yes" ]; then
    echo "Aborted."
    exit 0
fi

echo "Deleting ingress..."
$KUBECTL delete -f ingress.yaml --ignore-not-found=true

echo "Deleting applications..."
$KUBECTL delete -f dashboard.yaml --ignore-not-found=true
$KUBECTL delete -f api.yaml --ignore-not-found=true

echo "Deleting Redis..."
$KUBECTL delete -f redis.yaml --ignore-not-found=true

echo "Deleting database..."
$KUBECTL delete -f database.yaml --ignore-not-found=true

echo "Deleting configuration..."
$KUBECTL delete -f secrets.yaml --ignore-not-found=true
$KUBECTL delete -f configmap.yaml --ignore-not-found=true

echo "Deleting namespace (this will remove PVCs and all data)..."
$KUBECTL delete -f namespace.yaml --ignore-not-found=true

echo ""
echo "✅ Cleanup complete!"
echo ""
echo "⚠️  Note: PersistentVolumes may still exist. Check with:"
echo "  kubectl get pv"
echo ""
