package utils

import "flag"

type Flags struct {
	KubeConfig               *string
	Master                   *string
	CertificatesDirectory    *string
	EtcdServers              *string
	AuthenticationKubeConfig *string
	AuthorizationKubeConfig  *string
	SecurePort               *string
	CommandLine              *flag.FlagSet
}

func ParseFlags() *Flags {
	flags := &Flags{
		KubeConfig:               flag.String("kubeconfig", "", ""),
		Master:                   flag.String("master", "", ""),
		CertificatesDirectory:    flag.String("cert-dir", "", ""),
		EtcdServers:              flag.String("etcd-servers", "", ""),
		AuthenticationKubeConfig: flag.String("authentication-kubeconfig", "", ""),
		AuthorizationKubeConfig:  flag.String("authorization-kubeconfig", "", ""),
		SecurePort:               flag.String("secure-port", "", ""),
		CommandLine:              flag.CommandLine,
	}

	flag.Parse()
	return flags
}
