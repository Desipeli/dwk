/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	dummydwkv1 "dummy.dwk/api/v1"
)

// DummySiteReconciler reconciles a DummySite object
type DummySiteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=dummy.dwk.dummy.dwk,resources=dummysites,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dummy.dwk.dummy.dwk,resources=dummysites/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=dummy.dwk.dummy.dwk,resources=dummysites/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DummySite object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *DummySiteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here

	test := &dummydwkv1.DummySite{}
	err := r.Get(ctx, req.NamespacedName, test)
	if err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "unable to fetch Test")
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	url := test.Spec.Url
	fmt.Println("Downloading URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Log.Error(err, "error fetching the site")
	}
	defer resp.Body.Close()

	fmt.Println("Site loaded")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Log.Error(err, "error reading the site")
	}

	fmt.Println("site read")

	os.WriteFile("/tmp/index.html", body, 0644)

	fmt.Println("Done")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DummySiteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	port := os.Getenv("SERVER_PORT")
	fmt.Printf("SERVER_PORT: %s", port)
	if port == "" {
		log.Log.Info("env SERVER_PORT required")
	}
	httpPort := fmt.Sprintf(":%s", port)

	go func() {

		mux := http.NewServeMux()
		mux.HandleFunc("/", handleRoot)

		fmt.Println("Listening...")
		if err := http.ListenAndServe(httpPort, mux); err != nil {
			log.Log.Error(err, "HTTP server fail")
		}
	}()

	return ctrl.NewControllerManagedBy(mgr).
		For(&dummydwkv1.DummySite{}).
		Named("dummysite").
		Complete(r)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {

	body, err := os.ReadFile("/tmp/index.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error reading html file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "text/html")
	w.Write(body)
}
