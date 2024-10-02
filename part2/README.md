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

## Exercise 2.03: Keep them separated

New namespace for log-output and ping-pong apps.

```bash
$ kubectl create namespace dwk-apps
namespace/dwk-apps created

$ kubens dwk-apps 
Context "k3d-k3s-default" modified.
Active namespace is "dwk-apps".

$ kubectl apply -f part2/e_2.01/
deployment.apps/logpod-dep created
deployment.apps/ping-pong-dep created
ingress.networking.k8s.io/dwk-ingress created
persistentvolume/pings-pv created
persistentvolumeclaim/pings-claim created
service/logpod-svc created
service/ping-pong-svc created
```

## Exercise 2.04: Project v1.1

```bash
kubectl delete -f part2/e_2.02/
deployment.apps "todo-app-dep" deleted
ingress.networking.k8s.io "dwk-todo-ingress" deleted
persistentvolume "todo-pv" deleted
persistentvolumeclaim "todo-claim" deleted
service "todo-app-svc" deleted
service "todo-backend-svc" deleted


$ kubectl create namespace dwk-project
namespace/dwk-project created

$ kubens dwk-project 
Context "k3d-k3s-default" modified.
Active namespace is "dwk-project".

kubectl apply -f part2/e_2.02/
deployment.apps/todo-app-dep created
ingress.networking.k8s.io/dwk-todo-ingress created
persistentvolume/todo-pv created
persistentvolumeclaim/todo-claim created
service/todo-app-svc created
service/todo-backend-svc created

```

## Exercise 2.06: Documentation and ConfigMaps

[manifests](e_2.06/)

Created [configmap](e_2.06/configmap_reader.yaml) and added it to [deployment.yaml](e_2.06/deployment.yaml)

```bash
$ kubectl exec -it my-busybox -- wget -qO - http://logpod-svc:2345
file content: This is text from file
env variable: MESSAGE=Hello Wörld
2024-09-25 09:29:52.113 +0000 7a989879-d230-4376-a48e-74b78e95dbfc
Ping / Pongs: 39
```

## Exercise 2.07: Stateful applications

[Manifests](e_2.07/)

Changed pingpong app to save and read counter from postgres database. Set initialisazion [script](e_2.07/init-sql.yaml) for database using ConfigMap.

Created [secret](e_2.07/secret.enc.yaml) and encrypted it with command:

```bash
sops --encrypt --age age1mcxnf7uc4vwame00e36s9u5znc4lztxccg67fp223nrkk8yntq8qn96n3t --encrypted-regex '^(data)$' part2/e_2.07/secret.yaml > part2/e_2.07/secret.enc.yaml
```

Secret can be encrypted with private key after the key is exported:

```bash
$ export SOPS_AGE_KEY_FILE=key.txt

$ sops --decrypt secret.enc.yaml > secret.yaml
```

## Exercise 2.08: Poject v1.2

Changed the todo-backend to save and read todos to a postgres database in a separate pod. The database set up is similar to previous exercise.