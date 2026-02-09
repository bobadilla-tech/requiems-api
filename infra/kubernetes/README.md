# Kubernetes Deployment Guide

This directory contains Kubernetes manifests for deploying Requiem API on a Kubernetes cluster.

## Architecture

All services run in the `requiem` namespace on the same cluster:

- **PostgreSQL** (db) - Database with persistent storage (ClusterIP service)
- **Redis** (redis) - Cache and job queue (ClusterIP service)
- **Go API** (api) - Internal backend API (ClusterIP service)
- **Rails Dashboard** (dashboard) - Web interface (ClusterIP service with Ingress)
- **Sidekiq** (sidekiq) - Background job processor (Deployment only, no service)

### Service Types

All services use `ClusterIP` for internal communication within the cluster:
- `db:5432` - PostgreSQL (internal only)
- `redis:6379` - Redis (internal only)
- `api:8080` - Go API (internal only, exposed via Ingress to internal.requiems.xyz)
- `dashboard:80` - Rails dashboard (exposed via Ingress to requiems.xyz)

External access is managed through the Ingress controller.

## Prerequisites

1. **Kubernetes cluster** (v1.24+)
   - Can be a single-node cluster, multi-node cluster, or managed Kubernetes service
   - For single server deployment, use tools like:
     - [k3s](https://k3s.io/) - Lightweight Kubernetes
     - [MicroK8s](https://microk8s.io/) - Minimal Kubernetes
     - [kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/) - Standard Kubernetes

2. **kubectl** configured to access your cluster

3. **Ingress Controller** (for external access)
   ```bash
   # Install nginx ingress controller
   kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.1/deploy/static/provider/cloud/deploy.yaml
   ```

4. **cert-manager** (optional, for automatic SSL certificates)
   ```bash
   kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.0/cert-manager.yaml
   ```

5. **Docker images** built and available:
   ```bash
   # Build the Go API image
   cd apps/api
   docker build -t requiem-api:latest -f ../../infra/docker/api.Dockerfile .
   
   # Build the Rails dashboard image
   cd apps/dashboard
   docker build -t requiem-dashboard:latest -f ../../infra/docker/dashboard.Dockerfile .
   
   # If using a registry, tag and push:
   docker tag requiem-api:latest your-registry/requiem-api:latest
   docker push your-registry/requiem-api:latest
   docker tag requiem-dashboard:latest your-registry/requiem-dashboard:latest
   docker push your-registry/requiem-dashboard:latest
   ```

## Quick Start

### 1. Update Configuration

Edit `secrets.yaml` to set your production values:
```bash
# Generate a secure secret key base
SECRET_KEY_BASE=$(openssl rand -hex 64)

# Edit secrets.yaml and update:
# - SECRET_KEY_BASE
# - POSTGRES_PASSWORD (if changing from default)
# - CLOUDFLARE_* variables for production
# - BACKEND_SECRET for internal API authentication
```

### 2. Deploy to Kubernetes

Apply all manifests in order:

```bash
cd infra/kubernetes

# Create namespace
kubectl apply -f namespace.yaml

# Create configuration
kubectl apply -f configmap.yaml
kubectl apply -f secrets.yaml

# Deploy database and Redis
kubectl apply -f database.yaml
kubectl apply -f redis.yaml

# Wait for database and Redis to be ready
kubectl wait --for=condition=ready pod -l app=postgres -n requiem --timeout=300s
kubectl wait --for=condition=ready pod -l app=redis -n requiem --timeout=300s

# Deploy applications
kubectl apply -f api.yaml
kubectl apply -f dashboard.yaml

# Deploy ingress (external access)
kubectl apply -f ingress.yaml
```

### 3. Verify Deployment

```bash
# Check all pods are running
kubectl get pods -n requiem

# Check services
kubectl get svc -n requiem

# Check ingress
kubectl get ingress -n requiem

# View logs
kubectl logs -l app=api -n requiem
kubectl logs -l app=dashboard -n requiem
```

### 4. Access the Application

Once the ingress is configured and DNS is set up:
- **Dashboard**: https://requiems.xyz
- **Internal API**: https://internal.requiems.xyz (protected by X-Backend-Secret header)

## Scaling

Scale deployments as needed:

```bash
# Scale API instances
kubectl scale deployment api -n requiem --replicas=3

# Scale dashboard instances
kubectl scale deployment dashboard -n requiem --replicas=3

# Scale sidekiq workers
kubectl scale deployment sidekiq -n requiem --replicas=2
```

## Storage

The manifests use PersistentVolumeClaims for data persistence:
- `postgres-pvc`: 10Gi for PostgreSQL data
- `redis-pvc`: 1Gi for Redis data

These will automatically provision volumes from your cluster's default StorageClass.

## Updates and Rollouts

Update deployments with new images:

```bash
# Update API
kubectl set image deployment/api api=requiem-api:v1.2.0 -n requiem

# Update dashboard
kubectl set image deployment/dashboard dashboard=requiem-dashboard:v1.2.0 -n requiem
kubectl set image deployment/sidekiq sidekiq=requiem-dashboard:v1.2.0 -n requiem

# Check rollout status
kubectl rollout status deployment/api -n requiem
kubectl rollout status deployment/dashboard -n requiem
```

## Troubleshooting

### Check pod status
```bash
kubectl describe pod <pod-name> -n requiem
```

### View logs
```bash
kubectl logs <pod-name> -n requiem
kubectl logs -l app=api -n requiem --tail=100
```

### Execute commands in pods
```bash
# Access Rails console
kubectl exec -it deployment/dashboard -n requiem -- bundle exec rails console

# Access database
kubectl exec -it deployment/postgres -n requiem -- psql -U requiem
```

### Common Issues

1. **Pods stuck in Pending**: Check PVC status and StorageClass availability
2. **ImagePullBackOff**: Ensure Docker images are built and accessible
3. **CrashLoopBackOff**: Check logs for application errors
4. **Database connection issues**: Verify DATABASE_URL and db service is running

## Production Considerations

1. **Secrets Management**: Use Kubernetes Secrets or external secret management (Vault, AWS Secrets Manager, etc.)
2. **Resource Limits**: Adjust resource requests/limits based on your workload
3. **High Availability**: Run multiple replicas and use PodDisruptionBudgets
4. **Monitoring**: Set up Prometheus/Grafana for metrics and alerts
5. **Backups**: Regular PostgreSQL backups using pg_dump or volume snapshots
6. **SSL/TLS**: Use cert-manager for automatic certificate management
7. **Network Policies**: Restrict traffic between pods for security

## Single-Server Deployment with k3s

For a single-server deployment, k3s is recommended:

```bash
# Install k3s on your server
curl -sfL https://get.k3s.io | sh -

# Copy kubeconfig for kubectl
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown $USER ~/.kube/config

# k3s comes with Traefik ingress by default
# You can use it instead of nginx ingress

# Deploy all manifests as described above
cd infra/kubernetes
kubectl apply -f namespace.yaml
kubectl apply -f configmap.yaml
kubectl apply -f secrets.yaml
kubectl apply -f database.yaml
kubectl apply -f redis.yaml
kubectl apply -f api.yaml
kubectl apply -f dashboard.yaml
kubectl apply -f ingress.yaml
```

## Comparison with Docker Compose

### Docker Compose (Current Setup)
- Simpler for single-server deployments
- Good for development and small production setups
- Less overhead
- Manual scaling and updates

### Kubernetes (This Setup)
- Better for production at scale
- Automatic healing and restarts
- Easy horizontal scaling
- Rolling updates with zero downtime
- Better resource management
- Production-grade orchestration

Choose Docker Compose for simple deployments, Kubernetes for production scale.
