# Part 1

## Exercise 1.01: Gettings started

I created a simple program with go, that stores a random uuid to memory and prints it and timestamp every 5 seconds. The app can be found on [https://hub.docker.com/repository/docker/desipeli/dwk-log-output/general](https://hub.docker.com/repository/docker/desipeli/dwk-log-output/general)

Deploy to cluster:

```
kubectl create deployment log-output --image=desipeli/dwk-log-output
```