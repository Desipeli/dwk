# Part 5

## Exercise 5.01: DIY CRD & Controller

- Installed [kubebuilder](https://book.kubebuilder.io/)

```bash
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && sudo mv kubebuilder /usr/local/bin/
```

- Created new Go project and ran kubebuilder

```bash
$ kubectl apply -f resourcedefenition.yaml
$ kubectl apply -f dummysite.yaml

$ mkdir dummy-controller
$ cd dummy-controller
$ go mod init dummy.dwk

$ kubebuilder init --domain dummy.dwk



$ kubebuilder create api --group dummy.dwk --version v1 --kind DummySite
INFO Create Resource [y/n]                        
y
INFO Create Controller [y/n]                      
y
```

- Added Url string `json:"url"` to [dummysite_controller.go](e_5.01/dummy-controller/api/v1/dummysite_types.go)
- Controller logic in [dummysite-controller.go](e_5.01/dummy-controller/internal/controller/dummysite_controller.go)
- changed port from the cmd/main.go from 8081 to 8084

```bash
$ make manifests
$ make install

$ make run SERVER_PORT=8085
```

- Changed the api of dummysite.yaml from `apiVersion: dummy.dwk/v1` to `apiVersion: dummy.dwk.dummy.dwk/v1`, because kubebuilder added domain.
- `$ make docker-build docker-push IMG=desipeli/dwk-dummy:1.3`

Run with: `kubectl apply -f ./dummy-manifests`

![image](e_5.01/wikipedia.png)

## Exercise 5.02: Project, the Service Mesh Edition

- Had to add `config.linkerd.io/skip-outbound-ports: "4222"` annotation to backend and broadcaster.

[manifests](e_5.02/manifests/production/)

![linkerd image](e_5.02/linkerd.png)

## Exercise 5.03: Learn from external material

[script](e_5.03/script)

## Exercise 5.04: Wikipedia with init and sidecar

[manifests](e_5.04/)

- For some reason nginx displays 404 sometimes...

![image](e_5.04/5_04.png)


## Exercise 5.05: Platform comparison, Rancher and OpenShift

### OpenShift

- Self or fully managed service
- Includes Red Hats enhancements to Kubernetes
- Includes everything needed for development lifecycle: ci/cd, environments...
- Preinstalled and configured monitoring stack
- Security monitoring everywhere
- Cli and browser based web console to manage and visualize
- Run and manage also virtual machine workloads
- Central authentication and authorization
- Built-in container image registry
- Marketplace with add-ons 

Openshift seems to be more opinionated, and is probably easier to get running because of that. However, there is a risk of vendor lock in.

### Rancher

- Open source
- Centralized authentication, access control and observability
- For bare metal, private and public clouds
- app library
- audit logging and rate-limiting
- Simplifies Kubernetes
- Can manage clusters even if they are not created with Rancher
- No risk of vendor lock in

100% open source and no risk of vendor lock in (based on their site). Integrates other open source projects to enhance Kubernetes.

### Which is better?

For small organization Rancher, because it is open source, community version is free and no risk of vendor lock in.