# Part 3

- Started free trial with Google Cloud, started new project and installed gcloud
- Created new cluster: 

```bash 
$ gcloud container clusters create dwk-cluster --zone=europe-north1-b --cluster-version=1.30

$ sudo apt install google-cloud-cli-gke-gcloud-auth-plugin
```

## Exercise 3.01: Pingpong GKE

[Manifests](e_3.01/)

- Removed storageClass from the volumeClaimTemplate
- Added subPath to configuration
- Using LoadBalancer to expose the service

```bash
$ kubectl get svc --watch
NAME            TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
kubernetes      ClusterIP      34.118.224.1     <none>        443/TCP        4d19h
ping-pong-svc   LoadBalancer   34.118.235.199   <pending>     80:30115/TCP   31s
postgres-svc    ClusterIP      None             <none>        5432/TCP       73s
ping-pong-svc   LoadBalancer   34.118.235.199   34.88.127.25   80:30115/TCP   38s

$ curl 34.88.127.25/pingpong
pong 10
```