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