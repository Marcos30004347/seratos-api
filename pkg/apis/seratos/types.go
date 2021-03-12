package seratos

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Enviroment defines enviroment varialbles for the microservices
type Enviroment struct {
	Name  string
	Value string
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Microservice Resource
type Microservice struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec MicroserviceSpec
}

// MicroserviceSpec define the specification of the microservice object
type MicroserviceSpec struct {
	Image      string
	Sidecars   []string
	Replicas   int32
	Enviroment []Enviroment
	Info       MicroserviceInfo
}

// MicroserviceInfo defines the topology for a service
type MicroserviceInfo struct {
	Connections []MicroserviceConnection
	Secutiry    MicroserviceSecutiry
}

// MicroserviceConnection defines some proxys for the service
type MicroserviceConnection struct {
	Service string
	Host    string
	Ports   ProxyPorts
	URL     string
}

// ProxyPorts defines the proxy ports
type ProxyPorts struct {
	TCP   int32
	HTTP2 int32
}

// MicroserviceSecutiry defines microservice secutiry parameters
type MicroserviceSecutiry struct {
	Policy string
	Blocks []string
	Allow  []string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MicroserviceList is a list of Microservices objects.
type MicroserviceList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Microservice
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Sidecar Resource
type Sidecar struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec SidecarSpec
}

// SidecarSpec Specification
type SidecarSpec struct {
	Image      string
	Command    []string
	Args       []string
	Files      []SidecarFiles
	Enviroment []Enviroment
}

// SidecarFiles is the files created by the sidecar
type SidecarFiles struct {
	Filepath    string
	Description string
	Template    string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SidecarList is a list of Sidecar objects.
type SidecarList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Sidecar
}
