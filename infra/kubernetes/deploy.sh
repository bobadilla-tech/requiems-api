#!/bin/bash
# Deploy Requiem API to Kubernetes
# Usage: ./deploy.sh [namespace]

set -e

NAMESPACE=${1:-requiem}
KUBECTL="kubectl"

echo "🚀 Deploying Requiem API to Kubernetes namespace: $NAMESPACE"
echo ""

# Check if kubectl is installed
if ! command -v $KUBECTL &> /dev/null; then
    echo "❌ kubectl is not installed. Please install kubectl first."
    exit 1
fi

# Check if cluster is accessible
if ! $KUBECTL cluster-info &> /dev/null; then
    echo "❌ Cannot connect to Kubernetes cluster. Please configure kubectl."
    exit 1
fi

echo "✅ Connected to Kubernetes cluster"
echo ""

# Create namespace
echo "📦 Creating namespace..."
$KUBECTL apply -f namespace.yaml

# Create configuration
echo "⚙️  Creating configuration..."
$KUBECTL apply -f configmap.yaml
$KUBECTL apply -f secrets.yaml

# Deploy database
echo "🗄️  Deploying PostgreSQL database..."
$KUBECTL apply -f database.yaml

# Deploy Redis
echo "📮 Deploying Redis..."
$KUBECTL apply -f redis.yaml

# Wait for database to be ready
echo "⏳ Waiting for database to be ready..."
$KUBECTL wait --for=condition=ready pod -l app=postgres -n $NAMESPACE --timeout=300s || {
    echo "❌ Database failed to start. Check logs with: kubectl logs -l app=postgres -n $NAMESPACE"
    exit 1
}

# Wait for Redis to be ready
echo "⏳ Waiting for Redis to be ready..."
$KUBECTL wait --for=condition=ready pod -l app=redis -n $NAMESPACE --timeout=300s || {
    echo "❌ Redis failed to start. Check logs with: kubectl logs -l app=redis -n $NAMESPACE"
    exit 1
}

# Deploy API
echo "🔌 Deploying Go API..."
$KUBECTL apply -f api.yaml

# Deploy Dashboard and Sidekiq
echo "🎨 Deploying Rails Dashboard and Sidekiq..."
$KUBECTL apply -f dashboard.yaml

# Deploy Ingress
echo "🌐 Deploying Ingress..."
$KUBECTL apply -f ingress.yaml

echo ""
echo "✅ Deployment complete!"
echo ""
echo "📊 Check status with:"
echo "  kubectl get pods -n $NAMESPACE"
echo "  kubectl get svc -n $NAMESPACE"
echo "  kubectl get ingress -n $NAMESPACE"
echo ""
echo "📝 View logs with:"
echo "  kubectl logs -l app=api -n $NAMESPACE"
echo "  kubectl logs -l app=dashboard -n $NAMESPACE"
echo ""
echo "🔍 Describe resources with:"
echo "  kubectl describe pod <pod-name> -n $NAMESPACE"
echo ""
