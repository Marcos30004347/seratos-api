package seratos

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
