# Part 3

- Started free trial with Google Cloud, started new project and installed gcloud
- Created new cluster: 

```bash 
$ gcloud container clusters create dwk-cluster --zone=europe-north1-b --cluster-version=1.30

$ sudo apt install google-cloud-cli-gke-gcloud-auth-plugin
```

Delete the cluster with:

```bash
gcloud container clusters delete dwk-cluster --zone=europe-north1-b 
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

## Exercise 3.02: Back to Ingress

[Manifests](e_3.02/)

- Deployed log and pingpong apps to GKE
- Changed the communication between log and pingpong to use json
- Services are exposed with Ingress

```bash
$ curl http://34.144.230.54/
file content: This is text from file
<br>env variable: MESSAGE=Hello Wörld<br>2024-10-08 10:43:23.665 +0000 51ac1a7b-46e6-45a9-a2f1-3f9843084874<br>Ping / Pongs: 35

$ curl http://34.144.230.54/pingpong
pong 36
```

## Exercise 3.03:

- Created service account here: [https://console.cloud.google.com/iam-admin/serviceaccounts/](https://console.cloud.google.com/iam-admin/serviceaccounts/) with required roles
- Created key and added it to GitHub secrets
- Created Docker repository Google Cloud