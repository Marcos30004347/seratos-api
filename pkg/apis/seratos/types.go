package seratos

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
	Container string
	Replicas  int32
	Env       []Env
}

// Env defines enviroment varialbles for the microservices
type Env struct {
	Name  string
	Value string
}

// MicroserviceTopology defines the topology for a service
type MicroserviceTopology struct {
	Proxys   []MicroserviceProxy
	Secutiry MicroserviceSecutiry
}

// MicroserviceProxy defines some proxys for the service
type MicroserviceProxy struct {
	Service string
	Host    string
	Ports   MicroserviceProxyPorts
}

// MicroserviceProxyPorts defines the proxy ports
type MicroserviceProxyPorts struct {
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
