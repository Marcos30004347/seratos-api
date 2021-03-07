package main

import (
	"flag"
	"os"
	"sync"

	"github.com/spf13/cobra"

	genericapiserver "k8s.io/apiserver/pkg/server"

	"github.com/Marcos30004347/seratos-api/pkg/cmd/server"

	"k8s.io/component-base/logs"
	"k8s.io/klog"
)

var (
	stopCh  <-chan struct{}
	command *cobra.Command
	wg      sync.WaitGroup
)

func runServerAPI(wg *sync.WaitGroup) {
	if err := command.Execute(); err != nil {
		klog.Fatal(err)
	}
}

func main() {
	// Read command line flags
	// Init logs
	flag.String("kubeconfig", "", "a string")
	flag.String("master", "", "a string")
	flag.String("cert-dir", "", "a string")
	flag.String("etcd-servers", "", "a string")
	flag.String("authentication-kubeconfig", "", "a string")
	flag.String("authorization-kubeconfig", "", "a string")
	flag.String("secure-port", "", "a string")

	flag.Parse()

	logs.InitLogs()
	defer logs.FlushLogs()

	stopCh = genericapiserver.SetupSignalHandler()
	options := server.NewCustomServerOptions(os.Stdout, os.Stderr)

	command = server.NewCommandStartCustomServer(options, stopCh)
	command.Flags().AddGoFlagSet(flag.CommandLine)

	wg.Add(1)

	go runServerAPI(&wg)

	wg.Wait()
}
