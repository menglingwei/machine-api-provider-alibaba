/*
Copyright 2021 The Kubernetes Authors.

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

package main

import (
	"flag"
	"os"
	"time"

	"github.com/AliyunContainerService/cluster-api-provider-alibabacloud/pkg/version"

	"github.com/openshift/machine-api-operator/pkg/metrics"

	"sigs.k8s.io/controller-runtime/pkg/cache"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"

	actuator "github.com/AliyunContainerService/cluster-api-provider-alibabacloud/pkg/actuators/machine"
	alibabacloudClient "github.com/AliyunContainerService/cluster-api-provider-alibabacloud/pkg/client"

	"sigs.k8s.io/controller-runtime/pkg/controller"

	machineactuator "github.com/AliyunContainerService/cluster-api-provider-alibabacloud/pkg/actuators/machine"
	machinesetcontroller "github.com/AliyunContainerService/cluster-api-provider-alibabacloud/pkg/actuators/machineset"
	"github.com/AliyunContainerService/cluster-api-provider-alibabacloud/pkg/apis"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	"github.com/openshift/machine-api-operator/pkg/controller/machine"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// The default durations for the leader electrion operations.
var (
	leaseDuration    = 120 * time.Second
	renewDealine     = 110 * time.Second
	retryPeriod      = 20 * time.Second
	leaderElectionID = "cluster-api-provider-alibabacloud-leader"
)

func main() {
	healthAddr := flag.String(
		"health-addr",
		":9440",
		"The address for health checking.",
	)

	metricsAddress := flag.String(
		"metrics-bind-address",
		metrics.DefaultMachineMetricsAddress,
		"Address for hosting metrics",
	)

	leaderElectResourceNamespace := flag.String(
		"leader-elect-resource-namespace",
		"",
		"The namespace of resource object that is used for locking during leader election. If unspecified and running in cluster, defaults to the service account namespace for the controller. Required for leader-election outside of a cluster.",
	)

	leaderElect := flag.Bool(
		"leader-elect",
		false,
		"Start a leader election client and gain leadership before executing the main loop. Enable this when running replicated components for high availability.",
	)

	leaderElectLeaseDuration := flag.Duration(
		"leader-elect-lease-duration",
		leaseDuration,
		"The duration that non-leader candidates will wait after observing a leadership renewal until attempting to acquire leadership of a led but unrenewed leader slot. This is effectively the maximum duration that a leader can be stopped before it is replaced by another candidate. This is only applicable if leader election is enabled.",
	)

	watchNamespace := flag.String(
		"namespace",
		"",
		"Namespace that the controller watches to reconcile machine-api objects. If unspecified, the controller watches for machine-api objects across all namespaces.",
	)

	enableMetrics := flag.Bool(
		"enable-metrics",
		true,
		"Whether to enable metrics, Default value true. If you test in local, you can disable it",
	)

	printVersion := flag.Bool(
		"enable-print-version",
		true,
		"Whether to print release version, Default value true.",
	)

	klog.InitFlags(nil)
	flag.Set("logtostderr", "true")
	flag.Parse()

	// print release version
	if printVersion != nil && *printVersion {
		klog.Infof("The cluster-api-provider-alibabacloud version {%s}", version.PrintVerboseVersionInfo())
	}

	cfg := config.GetConfigOrDie()
	syncPeriod := 10 * time.Minute

	opts := manager.Options{
		LeaderElection:          *leaderElect,
		LeaderElectionNamespace: *leaderElectResourceNamespace,
		LeaderElectionID:        leaderElectionID,
		LeaseDuration:           leaderElectLeaseDuration,
		HealthProbeBindAddress:  *healthAddr,
		SyncPeriod:              &syncPeriod,
		// Slow the default retry and renew election rate to reduce etcd writes at idle: BZ 1858400
		RetryPeriod:   &retryPeriod,
		RenewDeadline: &renewDealine,
	}

	if enableMetrics != nil && *enableMetrics {
		opts.MetricsBindAddress = *metricsAddress
	}

	if *watchNamespace != "" {
		opts.Namespace = *watchNamespace
		klog.Infof("Watching machine-api objects only in namespace %q for reconciliation.", opts.Namespace)
	}

	mgr, err := manager.New(cfg, opts)
	if err != nil {
		klog.Fatalf("Error creating manager: %v", err)
	}

	// Setup Scheme for all resources
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		klog.Fatal(err)
	}

	if err := v1beta1.AddToScheme(mgr.GetScheme()); err != nil {
		klog.Fatal(err)
	}

	if err := configv1.AddToScheme(mgr.GetScheme()); err != nil {
		klog.Fatal(err)
	}

	configManagedClient, startCache, err := newConfigManagedClient(mgr)
	if err != nil {
		klog.Fatal(err)
	}
	mgr.Add(startCache)

	// Initialize machine actuator.
	machineActuator := machineactuator.NewActuator(machineactuator.ActuatorParams{
		Client:                    mgr.GetClient(),
		EventRecorder:             mgr.GetEventRecorderFor("alibabacloud-controller"),
		AlibabaCloudClientBuilder: alibabacloudClient.NewClient,
		ConfigManagedClient:       configManagedClient,
		ReconcilerBuilder:         actuator.NewReconciler,
	})

	if err := machine.AddWithActuator(mgr, machineActuator); err != nil {
		klog.Fatal(err)
	}

	ctrl.SetLogger(klogr.New())
	setupLog := ctrl.Log.WithName("setup")
	if err = (&machinesetcontroller.Reconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("MachineSet"),
	}).SetupWithManager(mgr, controller.Options{}); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MachineSet")
		os.Exit(1)
	}

	if err := mgr.AddReadyzCheck("ping", healthz.Ping); err != nil {
		klog.Fatal(err)
	}

	if err := mgr.AddHealthzCheck("ping", healthz.Ping); err != nil {
		klog.Fatal(err)
	}

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		klog.Fatalf("Failed to run manager: %v", err)
	}
}

// newConfigManagedClient returns a controller-runtime client that can be used to access the openshift-config-managed
// namespace.
func newConfigManagedClient(mgr manager.Manager) (runtimeclient.Client, manager.Runnable, error) {
	cacheOpts := cache.Options{
		Scheme:    mgr.GetScheme(),
		Mapper:    mgr.GetRESTMapper(),
		Namespace: alibabacloudClient.KubeCloudConfigNamespace,
	}

	c, err := cache.New(mgr.GetConfig(), cacheOpts)
	if err != nil {
		return nil, nil, err
	}

	clientOpts := runtimeclient.Options{
		Scheme: mgr.GetScheme(),
		Mapper: mgr.GetRESTMapper(),
	}

	cachedClient, err := cluster.DefaultNewClient(c, config.GetConfigOrDie(), clientOpts)
	if err != nil {
		return nil, nil, err
	}

	return cachedClient, c, nil
}
