# Part 2

## Exercise 2.01: Connecting pods

- Removed the shared volume
- Added route to ping pong service "/" to get the pongs
- changed the log-reader to use api.

busybox debug:

``` bash
kubectl exec -it my-busybox -- wget -qO - http://ping-pong-svc:3456
36

kubectl exec -it my-busybox -- wget -qO - http://10.42.0.25:8002
36
```

## Exercise 2.02: Project v1.0

[manifests](e_2.02/)

I created [todo-backend](../todo-backend/) for saving and fetching todos. Address of the backend is passed through a environment variable to the todo-app in the [deployment.yaml](e_2.02/deployment.yaml). New todos can be created using a form in the website, and the list of todos is updated when a new todo is created.