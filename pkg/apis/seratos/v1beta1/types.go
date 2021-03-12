package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Enviroment defines enviroment varialbles for the microservices
type Enviroment struct {
	Name  string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
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
	Image      string           `json:"image,omitempty" protobuf:"bytes,1,opt,name=image"`
	Sidecars   []string         `json:"sidecars,omitempty" protobuf:"bytes,2,opt,name=sidecars"`
	Replicas   int32            `json:"replicas,omitempty" protobuf:"bytes,3,opt,name=replicas"`
	Enviroment []Enviroment     `json:"enviroment,omitempty" protobuf:"bytes,4,opt,name=enviroment"`
	Info       MicroserviceInfo `json:"info,omitempty" protobuf:"bytes,5,opt,name=info"`
}

// MicroserviceInfo defines the topology for a service
type MicroserviceInfo struct {
	Connections []MicroserviceConnection `json:"connections,omitempty" protobuf:"bytes,1,opt,name=connections"`
	Secutiry    MicroserviceSecutiry     `json:"security,omitempty" protobuf:"bytes,2,opt,name=security"`
}

// MicroserviceConnection defines some proxys for the service
type MicroserviceConnection struct {
	Service string     `json:"service,omitempty" protobuf:"bytes,1,opt,name=service"`
	Host    string     `json:"host,omitempty" protobuf:"bytes,2,opt,name=host"`
	Ports   ProxyPorts `json:"ports,omitempty" protobuf:"bytes,3,opt,name=ports"`
	URL     string     `json:"url,omitempty" protobuf:"bytes,4,opt,name=url"`
}

// ProxyPorts defines the proxy ports
type ProxyPorts struct {
	TCP   int32 `json:"tcp,omitempty" protobuf:"bytes,1,opt,name=tcp"`
	HTTP2 int32 `json:"http2,omitempty" protobuf:"bytes,1,opt,name=http2"`
}

// MicroserviceSecutiry defines microservice secutiry
type MicroserviceSecutiry struct {
	Policy string   `json:"policity,omitempty" protobuf:"bytes,1,opt,name=policity"`
	Blocks []string `json:"blocks,omitempty" protobuf:"bytes,1,opt,name=blocks"`
	Allow  []string `json:"allow,omitempty" protobuf:"bytes,1,opt,name=allow"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MicroserviceList is a list of Microservices objects.
type MicroserviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Microservice `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Sidecar Resource
type Sidecar struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec SidecarSpec `json:"spec,omitempty" protobuf:"bytes,2,rep,name=spec"`
}

// SidecarSpec Specification
type SidecarSpec struct {
	Image      string         `json:"image,omitempty" protobuf:"bytes,1,rep,name=image"`
	Command    []string       `json:"command,omitempty" protobuf:"bytes,2,rep,name=command"`
	Args       []string       `json:"args,omitempty" protobuf:"bytes,3,rep,name=args"`
	Files      []SidecarFiles `json:"files,omitempty" protobuf:"bytes,4,rep,name=files"`
	Enviroment []Enviroment   `json:"enviroment,omitempty" protobuf:"bytes,5,rep,name=enviroment"`
}

// SidecarFiles is the files created by the sidecar
type SidecarFiles struct {
	Filepath    string `json:"filepath,omitempty" protobuf:"bytes,1,rep,name=filepath"`
	Description string `json:"description,omitempty" protobuf:"bytes,2,rep,name=description"`
	Template    string `json:"template,omitempty" protobuf:"bytes,3,rep,name=template"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SidecarList is a list of Sidecar objects.
type SidecarList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Sidecar `json:"items" protobuf:"bytes,2,rep,name=items"`
}
