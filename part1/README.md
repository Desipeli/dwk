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

```
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

```
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

```
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

```
kubectl port-forward todo-app-dep-85d4cf558f-w5p2v 8001:8000
Forwarding from 127.0.0.1:8001 -> 8000
Forwarding from [::1]:8001 -> 8000
Handling connection for 8001
```