# Local Build and Test Instructions

## Build Local Image

```bash
# Build the litmus-go image locally with 'ci' tag
cd /home/bkylberg/git/litmus/litmus-go
make build-amd64

# This creates: litmuschaos/go-runner:ci
```

## Load Image to Kubernetes

If using kind/minikube:

```bash
# For kind
kind load docker-image litmuschaos/go-runner:ci

# For minikube
minikube image load litmuschaos/go-runner:ci
```

## Deploy and Test

```bash
# 1. Deploy sample Java app
kubectl apply -f experiments/generic/jvm-cpu-hog/sample-app.yaml

# 2. Wait for app to be ready
kubectl wait --for=condition=ready pod -l app=java-app --timeout=60s

# 3. Apply RBAC
kubectl apply -f experiments/generic/jvm-cpu-hog/rbac.yaml

# 4. Install experiment (uses local image)
kubectl apply -f experiments/generic/jvm-cpu-hog/experiment.yaml

# 5. Run chaos
kubectl apply -f experiments/generic/jvm-cpu-hog/engine.yaml

# 6. Monitor
kubectl get chaosengine -w
kubectl get chaosresult
kubectl logs -l name=jvm-cpu-hog -f
```

## Image Configuration

The experiment.yaml now uses:
- **Image**: `litmuschaos/go-runner:ci` (local build tag)
- **ImagePullPolicy**: `IfNotPresent` (uses local image)

This prevents pulling from Docker Hub and uses your local build.
