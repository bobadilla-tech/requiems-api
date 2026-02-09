# Service Networking Design Decision

## Problem Statement
> "Decide what service of kubernetes to use for this, the ruby on rails and the go app and the db will be in the same server"

## Solution: ClusterIP Services for Internal Communication

All components (Rails dashboard, Go API, PostgreSQL, Redis) run in the same Kubernetes cluster and use **ClusterIP** service type for internal communication.

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                     Kubernetes Cluster                       │
│  ┌────────────────────────────────────────────────────────┐ │
│  │                  Namespace: requiem                     │ │
│  │                                                         │ │
│  │  ┌──────────────────────────────────────────────────┐ │ │
│  │  │          ClusterIP Services (Internal)            │ │ │
│  │  │                                                    │ │ │
│  │  │  ┌─────────────┐    ┌──────────────┐            │ │ │
│  │  │  │ db:5432     │    │ redis:6379   │            │ │ │
│  │  │  │ PostgreSQL  │    │ Redis        │            │ │ │
│  │  │  └──────┬──────┘    └──────┬───────┘            │ │ │
│  │  │         │                   │                     │ │ │
│  │  │         └───────┬───────────┴──────┬────────────┐ │ │
│  │  │                 │                  │            │ │ │
│  │  │         ┌───────▼──────┐   ┌──────▼─────┐     │ │ │
│  │  │         │ api:8080     │   │ dashboard:80│     │ │ │
│  │  │         │ Go Backend   │   │ Rails App    │     │ │ │
│  │  │         └───────┬──────┘   └──────┬───────┘     │ │ │
│  │  └─────────────────┼─────────────────┼─────────────┘ │ │
│  │                    │                 │                │ │
│  └────────────────────┼─────────────────┼────────────────┘ │
│                       │                 │                   │
│  ┌────────────────────▼─────────────────▼────────────────┐ │
│  │              Ingress Controller                        │ │
│  │  • requiems.xyz → dashboard:80                        │ │
│  │  • internal.requiems.xyz → api:8080                   │ │
│  └────────────────────────────────────────────────────────┘ │
└───────────────────────────────┬──────────────────────────────┘
                                │
                        Internet Access
```

## Service Type Decision: ClusterIP

### What is ClusterIP?

**ClusterIP** is the default Kubernetes service type that:
- Creates an internal IP address accessible only within the cluster
- Provides stable DNS names for service discovery (e.g., `db`, `redis`, `api`, `dashboard`)
- Enables pod-to-pod communication without exposing services externally
- Is the most secure option for internal services

### Why ClusterIP for All Services?

Since all components run on the **same Kubernetes cluster** (same server or multi-node cluster), they can communicate efficiently using internal cluster networking:

1. **PostgreSQL (db:5432)** - ClusterIP
   - Only needs to be accessed by Go API and Rails dashboard
   - No external access required
   - Security: Not exposed to internet

2. **Redis (redis:6379)** - ClusterIP
   - Only needs to be accessed by Rails dashboard and Sidekiq
   - No external access required
   - Security: Not exposed to internet

3. **Go API (api:8080)** - ClusterIP
   - Accessed by Rails dashboard for internal operations
   - Accessed via Ingress controller for external API calls to internal.requiems.xyz
   - Security: Only exposed through Ingress with proper authentication

4. **Rails Dashboard (dashboard:80)** - ClusterIP
   - Accessed via Ingress controller for web interface at requiems.xyz
   - Security: Only exposed through Ingress

5. **Sidekiq** - No Service
   - Background job processor
   - Doesn't need to receive requests
   - Connects to Redis and database directly

## External Access: Ingress

For external access to the dashboard and API, we use a **Kubernetes Ingress** resource:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
spec:
  rules:
  - host: requiems.xyz
    http:
      paths:
      - backend:
          service:
            name: dashboard  # ClusterIP service
            port: 80
  - host: internal.requiems.xyz
    http:
      paths:
      - backend:
          service:
            name: api  # ClusterIP service
            port: 8080
```

The Ingress controller:
- Acts as a reverse proxy (like Caddy in Docker Compose setup)
- Handles SSL/TLS termination
- Routes external traffic to internal ClusterIP services
- Provides a single point of entry to the cluster

## Alternative Service Types (Not Used)

### NodePort
- Exposes service on each node's IP at a static port (30000-32767)
- **Not used**: Would expose services on non-standard ports
- **Not used**: Less secure, exposes services on all nodes

### LoadBalancer
- Provisions an external load balancer (cloud provider)
- **Not used**: More expensive, creates external IPs for each service
- **Not used**: Overkill when Ingress provides all external routing needs

### ExternalName
- Maps service to a DNS name
- **Not used**: All our services are in the same cluster

## Benefits of This Design

1. **Security**
   - Database and Redis are not exposed to the internet
   - Only Ingress provides external access
   - Internal services communicate on private cluster network

2. **Simplicity**
   - All services use the same type (ClusterIP)
   - Consistent service discovery (DNS names)
   - Easy to understand and maintain

3. **Performance**
   - Internal communication uses cluster network
   - No external routing for inter-service calls
   - Low latency between services

4. **Flexibility**
   - Can scale any service independently
   - Can add more services easily
   - Works on single-node or multi-node clusters

5. **Cost-Effective**
   - No external load balancers per service
   - Single Ingress controller handles all external traffic
   - Efficient resource usage

## Service Discovery

Services can reach each other using DNS names:

```bash
# From Rails dashboard container
$ curl http://api:8080/healthz
$ psql postgres://requiem:requiem@db:5432/requiem

# From Go API container  
$ psql postgres://requiem:requiem@db:5432/requiem
$ redis-cli -h redis -p 6379 PING

# From Sidekiq container
$ redis-cli -h redis -p 6379 PING
$ psql postgres://requiem:requiem@db:5432/requiem
```

Kubernetes automatically resolves:
- `db` → PostgreSQL ClusterIP
- `redis` → Redis ClusterIP
- `api` → Go API ClusterIP
- `dashboard` → Rails Dashboard ClusterIP

## Comparison with Docker Compose

In Docker Compose, we also use internal networking:

```yaml
services:
  api:
    ports:
      - "6969:8080"  # External:Internal
    depends_on:
      - db
      - redis
```

In Kubernetes with ClusterIP:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api
spec:
  type: ClusterIP  # Internal only
  ports:
    - port: 8080
```

Both achieve the same goal: internal communication with selective external access.

## Conclusion

**ClusterIP services for all components** is the correct choice because:
- All services run in the same Kubernetes cluster
- Internal communication is secure and efficient
- External access is controlled through a single Ingress point
- This is the Kubernetes best practice for microservices architecture

This design provides the same networking behavior as the Docker Compose setup but with the added benefits of Kubernetes orchestration, scaling, and high availability.
