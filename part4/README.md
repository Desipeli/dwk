# Part 4

## Exercise 4.01: Readiness Probe

When database is not applied

```bash
$ kubectl get po
NAME                            READY   STATUS    RESTARTS   AGE
logpod-dep-6456bbc548-2n65f     1/2     Running   0          22m
ping-pong-dep-b4fb67ff7-98f7x   0/1     Running   0          22m
```

```bash
$ kubectl describe pod ping-pong-dep-b4fb67ff7-98f7x
...
Events:
  Type     Reason     Age                   From               Message
  ----     ------     ----                  ----               -------
  Normal   Scheduled  23m                   default-scheduler  Successfully assigned default/ping-pong-dep-b4fb67ff7-98f7x to k3d-k3s-default-server-0
  Normal   Pulling    23m                   kubelet            Pulling image "desipeli/dwk-pingpong:4.01"
  Normal   Pulled     23m                   kubelet            Successfully pulled image "desipeli/dwk-pingpong:4.01" in 859ms (859ms including waiting). Image size: 10422254 bytes.
  Normal   Created    23m                   kubelet            Created container ping-pong
  Normal   Started    23m                   kubelet            Started container ping-pong
  Warning  Unhealthy  3m4s (x255 over 22m)  kubelet            Readiness probe failed: HTTP probe failed with statuscode: 500
```

After database is applied

```bash
$ kubectl get po
NAME                            READY   STATUS    RESTARTS   AGE
logpod-dep-6456bbc548-2n65f     2/2     Running   0          25m
ping-pong-dep-b4fb67ff7-98f7x   1/1     Running   0          25m
postgres-stset-0                1/1     Running   0          16s
```