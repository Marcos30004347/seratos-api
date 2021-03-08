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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Event struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec EventSpec
}

type EventSpec struct {
	Kind        string
	SubjectName string
	Event       string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EventList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Event
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EventHandler struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec EventHandlerSpec
}

type EventHandlerSpec struct {
	SubjectKind string
	Command     EventHandlerCommand
	Selector    metav1.LabelSelector
}

type EventHandlerCommand struct {
	Exec string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EventHandlerList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []EventHandler
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EventBinding struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec EventBindingSpec
}

type EventBindingSpec struct {
	On       []string
	Handler  string
	Subjects []EventBindingSubjects
}

type EventBindingSubjects struct {
	Kind     string
	Selector metav1.LabelSelector
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EventBindingList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []EventBinding
}
