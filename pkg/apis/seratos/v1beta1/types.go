package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// Foo specifies an offered Foo with toppings.
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo objects.
type FooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Foo `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// Microservice Resource
type Microservice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec MicroserviceSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// MicroserviceSpec define the specification of the microservice object
type MicroserviceSpec struct {
	Container string `json:"container" protobuf:"bytes,1,opt,name=container"`
	Env       []Env  `json:"env" protobuf:"bytes,2,opt,name=env"`
}

// Env defines enviroment varialbles for the microservices
type Env struct {
	Name  string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

// MicroserviceTopology defines the topology for a service
type MicroserviceTopology struct {
	Proxys   []Proxy              `json:"proxys" protobuf:"bytes,1,opt,name=proxys"`
	Secutiry MicroserviceSecutiry `json:"security" protobuf:"bytes,2,opt,name=security"`
}

// Proxy defines some proxys for the service
type Proxy struct {
	Service string     `json:"service" protobuf:"bytes,1,opt,name=service"`
	Host    string     `json:"host" protobuf:"bytes,2,opt,name=host"`
	Ports   ProxyPorts `json:"ports" protobuf:"bytes,3,opt,name=ports"`
}

// ProxyPorts defines the proxy ports
type ProxyPorts struct {
	TCP   int32 `json:"tcp" protobuf:"bytes,1,opt,name=tcp"`
	HTTP2 int32 `json:"http2" protobuf:"bytes,1,opt,name=http2"`
}

// MicroserviceSecutiry defines microservice secutiry
type MicroserviceSecutiry struct {
	Policy string   `json:"policity" protobuf:"bytes,1,opt,name=policity"`
	Blocks []string `json:"blocks" protobuf:"bytes,1,opt,name=blocks"`
	Allow  []string `json:"allow" protobuf:"bytes,1,opt,name=allow"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MicroserviceList is a list of Microservices objects.
type MicroserviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Microservice `json:"items" protobuf:"bytes,2,rep,name=items"`
}
