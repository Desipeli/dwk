# Part 1

## Exercise 1.01: Gettings started

I created a simple program with Go that stores a random uuid to memory and prints it and timestamp every 5 seconds. The app can be found on [https://hub.docker.com/repository/docker/desipeli/dwk-log-output/general](https://hub.docker.com/repository/docker/desipeli/dwk-log-output/general)

Deploy to cluster:

```
kubectl create deployment log-output --image=desipeli/dwk-log-output
```

## Exercise 1.02: Project v0.1

I created a simple web server with Go. The port can be set with environment variable PORT, default is 8000.

[https://hub.docker.com/repository/docker/desipeli/dwk-todo-app/general](https://hub.docker.com/repository/docker/desipeli/dwk-todo-app/general)


```
kubectl create deployment todo-app --image=desipeli/dwk-todo-app
```

## Exercise 1.03: Declarative approach

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: log-output-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: log-output
  template:
    metadata:
      labels:
        app: log-output
    spec:
      containers:
        - name: log-output
          image: desipeli/dwk-log-output:0.1
```

## Exercise 1.04: Project v0.2

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: todo-app-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: todo-app
  template:
    metadata:
      labels:
        app: todo-app
    spec:
      containers:
        - name: todo-app
          image: desipeli/dwk-todo-app:0.2
```

## Exercise 1.05: Project v0.3

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: todo-app-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: todo-app
  template:
    metadata:
      labels:
        app: todo-app
    spec:
      containers:
        - name: todo-app
          image: desipeli/dwk-todo-app:0.3
          env:
          - name: PORT
            value: "8000"
```

```bash
kubectl port-forward todo-app-dep-85d4cf558f-w5p2v 8001:8000
Forwarding from 127.0.0.1:8001 -> 8000
Forwarding from [::1]:8001 -> 8000
Handling connection for 8001
```

## Exercise 1.06: Project 0.4

```yaml
apiVersion: v1
kind: Service
metadata:
  name: todo-app-svc
spec:
  type: NodePort
  selector:
    app: todo-app
  ports:
    - name: http
      nodePort: 30080 
      protocol: TCP
      port: 1234
      targetPort: 8000
```

```bash
k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
kubectl apply -f todo-app/manifests/deployment.yml 
kubectl apply -f todo-app/manifests/service.yml 
```

## Exercise 1.07: External access with Ingress

Added http server to the logger

deployment.yml
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: log-output-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: log-output
  template:
    metadata:
      labels:
        app: log-output
    spec:
      containers:
        - name: log-output
          image: desipeli/dwk-log-output:0.2
          env:
          - name: PORT
            value: "8001"
```

service.yaml
```yaml
apiVersion: v1
kind: Service
metadata:
  name: log-output-svc
spec:
  type: ClusterIP
  selector:
    app: log-output
  ports:
    - port: 2345
      protocol: TCP
      targetPort: 8001
```

ingress.yaml
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: dwk-ingress
spec:
  rules:
  - http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: log-output-svc
            port:
              number: 2345
```

Commands
```bash
k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
kubectl apply -f log-output/manifests/deployment.yml
kubectl apply -f log-output/manifests/service.yaml
kubectl apply -f log-output/manifests/ingress.yaml
```

## Exercise 1.08: Project v0.5

deployment.yaml
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: todo-app-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: todo-app
  template:
    metadata:
      labels:
        app: todo-app
    spec:
      containers:
        - name: todo-app
          image: desipeli/dwk-todo-app:0.3
          env:
          - name: PORT
            value: "8000"
```

service.yaml
```yaml
apiVersion: v1
kind: Service
metadata:
  name: todo-app-svc
spec:
  type: ClusterIP
  selector:
    app: todo-app
  ports:
    - port: 2345
      protocol: TCP
      targetPort: 8000
```

ingress.yaml
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: dwk-ingress
spec:
  rules:
  - http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: todo-app-svc
            port:
              number: 2345
```

Commands
```bash
kubectl delete -f log-output/manifests/ingress.yaml
dessu@fleda2:~/koulu/dwk$ kubectl apply -f todo-app/manifests/deployment.yaml 
dessu@fleda2:~/koulu/dwk$ kubectl apply -f todo-app/manifests/service.yaml 
dessu@fleda2:~/koulu/dwk$ kubectl apply -f todo-app/manifests/ingress.yaml 
```

## Exercise 1.09: More services

### ping-pong

deployment.yaml
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ping-pong-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ping-pong
  template:
    metadata:
      labels:
        app: ping-pong
    spec:
      containers:
        - name: ping-pong
          image: desipeli/dwk-pingpong:0.1
          env:
          - name: PORT
            value: "8002"
```

service.yaml
```yaml
apiVersion: v1
kind: Service
metadata:
  name: ping-pong-svc
spec:
  type: ClusterIP
  selector:
    app: ping-pong
  ports:
    - port: 3456
      protocol: TCP
      targetPort: 8002
```

### log-output

ingress.yaml
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: dwk-ingress
spec:
  rules:
  - http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: log-output-svc
            port:
              number: 2345
      - path: /pingpong
        pathType: Prefix
        backend:
          service:
            name: ping-pong-svc
            port:
              number: 3456
```

## Exercise 1.10: Even more services

I splitted the log-output app to log-reader and log-writer

[manifests](e_1.10/)

## Exercise 1.11: Persisting data

Created a local path in the node: `docker exec k3d-k3s-default-agent-0 mkdir -p /tmp/kube`

Modified [log-reader](../log-output-splitted/log-reader/main.go) and [ping-pong](../ping-pong/main.go) to use persistent directory.

[persistentvolume, claim and ingress](e_1.11/)

Applied new configurations for ping-pong and log-output apps.

## Exercise 1.12: Project v0.6

Created a new directory for persistent data `docker exec k3d-k3s-default-agent-0 mkdir -p /tmp/todo`

Modified [todo app](../todo-app/main.go) to download and save image and the timestamp.

[Manifests](e_1.12/)

