package main

import (
	"os"
	"sync"

	"github.com/Marcos30004347/seratos-api/pkg/cmd/server"
	controller "github.com/Marcos30004347/seratos-api/pkg/controllers/microservices"
	"github.com/Marcos30004347/seratos-api/pkg/runtime"
	"github.com/Marcos30004347/seratos-api/pkg/utils"

	kubeserver "k8s.io/apiserver/pkg/server"

	"k8s.io/klog"
)

func main() {
	flags := utils.ParseFlags()

	stopCh := kubeserver.SetupSignalHandler()

	kubeRuntime, err := runtime.NewKubernetesRuntime(*flags.KubeConfig, *flags.Master)

	if err != nil {
		klog.Fatalf("Error building Kube client: %s", err.Error())
	}

	seratosRuntime, err := runtime.NewSeratosRuntime(*flags.KubeConfig, *flags.Master)

	if err != nil {
		klog.Fatalf("Error building Seratos Runtime: %s", err.Error())
	}

	command := server.NewCommandStartCustomServer(
		server.NewCustomServerOptions(os.Stdout, os.Stderr),
		stopCh,
	)

	command.Flags().AddGoFlagSet(flags.CommandLine)

	msController := controller.NewController(
		kubeRuntime,
		seratosRuntime,
	)

	kubeRuntime.Run(stopCh)
	seratosRuntime.Run(stopCh)

	var wg sync.WaitGroup

	// Run the API and Controllers
	wg.Add(2)

	// API Server
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if err := command.Execute(); err != nil {
			klog.Fatal(err)
		}
	}(&wg)

	// Microservices Controller
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if err := msController.Run(2, stopCh); err != nil {
			klog.Fatalf("Error running controller: %s", err.Error())
		}
	}(&wg)

	wg.Wait()

	kubeRuntime.Cleanup()
	seratosRuntime.Cleanup()
}
