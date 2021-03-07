package seratos

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo specifies an offered pizza with toppings.
type Foo struct {
	metav1.TypeMeta
	metav1.ObjectMeta
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo objects.
type FooList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Foo
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
	Container string
	Env       []Env
}

// Env defines enviroment varialbles for the microservices
type Env struct {
	name  string
	value string
}

// MicroserviceTopology defines the topology for a service
type MicroserviceTopology struct {
	Proxys   []Proxy
	Secutiry MicroserviceSecutiry
}

// Proxy defines some proxys for the service
type Proxy struct {
	service string
	host    string
	ports   ProxyPorts
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
