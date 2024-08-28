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