package main

import (
	"os"
	"sync"

	kubeapiserver "k8s.io/apiserver/pkg/server"

	"github.com/Marcos30004347/seratos-api/pkg/cmd/server"
	"github.com/Marcos30004347/seratos-api/pkg/runtime"

	"k8s.io/klog"

	"github.com/Marcos30004347/seratos-api/pkg/controllers/controller"

	"github.com/Marcos30004347/seratos-api/pkg/utils"
)

func main() {
	flags := utils.ParseFlags()

	stopCh := kubeapiserver.SetupSignalHandler()

	kube, err := runtime.NewKubernetesRuntime(*flags.KubeConfig, *flags.Master, stopCh)
	if err != nil {
		klog.Fatalf("Error building Kube client: %s", err.Error())
	}

	seratosRuntime, err := runtime.NewSeratosRuntime(kube.RESTConfig)

	if err != nil {
		klog.Fatalf("Error building Seratos Runtime: %s", err.Error())
	}

	command := server.NewCommandStartCustomServer(
		server.NewCustomServerOptions(os.Stdout, os.Stderr),
		stopCh,
	)

	command.Flags().AddGoFlagSet(flags.CommandLine)

	msController := controller.NewController(
		kube,
		seratosRuntime,
	)

	kube.Run(stopCh)
	seratosRuntime.Run(stopCh)

	// controller.ReadEvents(kube)

	var wg sync.WaitGroup

	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		if err := command.Execute(); err != nil {
			klog.Fatal(err)
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		if err := msController.Run(2, stopCh); err != nil {
			klog.Fatalf("Error running controller: %s", err.Error())
		}
	}(&wg)

	wg.Wait()

	kube.Cleanup()
	seratosRuntime.Cleanup()
}
