// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	seratos "github.com/Marcos30004347/seratos-api/pkg/apis/seratos"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Enviroment)(nil), (*seratos.Enviroment)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Enviroment_To_seratos_Enviroment(a.(*Enviroment), b.(*seratos.Enviroment), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.Enviroment)(nil), (*Enviroment)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_Enviroment_To_v1beta1_Enviroment(a.(*seratos.Enviroment), b.(*Enviroment), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Microservice)(nil), (*seratos.Microservice)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Microservice_To_seratos_Microservice(a.(*Microservice), b.(*seratos.Microservice), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.Microservice)(nil), (*Microservice)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_Microservice_To_v1beta1_Microservice(a.(*seratos.Microservice), b.(*Microservice), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MicroserviceConnection)(nil), (*seratos.MicroserviceConnection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MicroserviceConnection_To_seratos_MicroserviceConnection(a.(*MicroserviceConnection), b.(*seratos.MicroserviceConnection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.MicroserviceConnection)(nil), (*MicroserviceConnection)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_MicroserviceConnection_To_v1beta1_MicroserviceConnection(a.(*seratos.MicroserviceConnection), b.(*MicroserviceConnection), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MicroserviceInfo)(nil), (*seratos.MicroserviceInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MicroserviceInfo_To_seratos_MicroserviceInfo(a.(*MicroserviceInfo), b.(*seratos.MicroserviceInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.MicroserviceInfo)(nil), (*MicroserviceInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_MicroserviceInfo_To_v1beta1_MicroserviceInfo(a.(*seratos.MicroserviceInfo), b.(*MicroserviceInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MicroserviceList)(nil), (*seratos.MicroserviceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MicroserviceList_To_seratos_MicroserviceList(a.(*MicroserviceList), b.(*seratos.MicroserviceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.MicroserviceList)(nil), (*MicroserviceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_MicroserviceList_To_v1beta1_MicroserviceList(a.(*seratos.MicroserviceList), b.(*MicroserviceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MicroserviceSecutiry)(nil), (*seratos.MicroserviceSecutiry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MicroserviceSecutiry_To_seratos_MicroserviceSecutiry(a.(*MicroserviceSecutiry), b.(*seratos.MicroserviceSecutiry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.MicroserviceSecutiry)(nil), (*MicroserviceSecutiry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_MicroserviceSecutiry_To_v1beta1_MicroserviceSecutiry(a.(*seratos.MicroserviceSecutiry), b.(*MicroserviceSecutiry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MicroserviceSpec)(nil), (*seratos.MicroserviceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MicroserviceSpec_To_seratos_MicroserviceSpec(a.(*MicroserviceSpec), b.(*seratos.MicroserviceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.MicroserviceSpec)(nil), (*MicroserviceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_MicroserviceSpec_To_v1beta1_MicroserviceSpec(a.(*seratos.MicroserviceSpec), b.(*MicroserviceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProxyPorts)(nil), (*seratos.ProxyPorts)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ProxyPorts_To_seratos_ProxyPorts(a.(*ProxyPorts), b.(*seratos.ProxyPorts), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.ProxyPorts)(nil), (*ProxyPorts)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_ProxyPorts_To_v1beta1_ProxyPorts(a.(*seratos.ProxyPorts), b.(*ProxyPorts), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Sidecar)(nil), (*seratos.Sidecar)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Sidecar_To_seratos_Sidecar(a.(*Sidecar), b.(*seratos.Sidecar), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.Sidecar)(nil), (*Sidecar)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_Sidecar_To_v1beta1_Sidecar(a.(*seratos.Sidecar), b.(*Sidecar), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SidecarFiles)(nil), (*seratos.SidecarFiles)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SidecarFiles_To_seratos_SidecarFiles(a.(*SidecarFiles), b.(*seratos.SidecarFiles), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.SidecarFiles)(nil), (*SidecarFiles)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_SidecarFiles_To_v1beta1_SidecarFiles(a.(*seratos.SidecarFiles), b.(*SidecarFiles), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SidecarList)(nil), (*seratos.SidecarList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SidecarList_To_seratos_SidecarList(a.(*SidecarList), b.(*seratos.SidecarList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.SidecarList)(nil), (*SidecarList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_SidecarList_To_v1beta1_SidecarList(a.(*seratos.SidecarList), b.(*SidecarList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SidecarSpec)(nil), (*seratos.SidecarSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SidecarSpec_To_seratos_SidecarSpec(a.(*SidecarSpec), b.(*seratos.SidecarSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seratos.SidecarSpec)(nil), (*SidecarSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seratos_SidecarSpec_To_v1beta1_SidecarSpec(a.(*seratos.SidecarSpec), b.(*SidecarSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_Enviroment_To_seratos_Enviroment(in *Enviroment, out *seratos.Enviroment, s conversion.Scope) error {
	out.Name = in.Name
	out.Value = in.Value
	return nil
}

// Convert_v1beta1_Enviroment_To_seratos_Enviroment is an autogenerated conversion function.
func Convert_v1beta1_Enviroment_To_seratos_Enviroment(in *Enviroment, out *seratos.Enviroment, s conversion.Scope) error {
	return autoConvert_v1beta1_Enviroment_To_seratos_Enviroment(in, out, s)
}

func autoConvert_seratos_Enviroment_To_v1beta1_Enviroment(in *seratos.Enviroment, out *Enviroment, s conversion.Scope) error {
	out.Name = in.Name
	out.Value = in.Value
	return nil
}

// Convert_seratos_Enviroment_To_v1beta1_Enviroment is an autogenerated conversion function.
func Convert_seratos_Enviroment_To_v1beta1_Enviroment(in *seratos.Enviroment, out *Enviroment, s conversion.Scope) error {
	return autoConvert_seratos_Enviroment_To_v1beta1_Enviroment(in, out, s)
}

func autoConvert_v1beta1_Microservice_To_seratos_Microservice(in *Microservice, out *seratos.Microservice, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_MicroserviceSpec_To_seratos_MicroserviceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Microservice_To_seratos_Microservice is an autogenerated conversion function.
func Convert_v1beta1_Microservice_To_seratos_Microservice(in *Microservice, out *seratos.Microservice, s conversion.Scope) error {
	return autoConvert_v1beta1_Microservice_To_seratos_Microservice(in, out, s)
}

func autoConvert_seratos_Microservice_To_v1beta1_Microservice(in *seratos.Microservice, out *Microservice, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_seratos_MicroserviceSpec_To_v1beta1_MicroserviceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_seratos_Microservice_To_v1beta1_Microservice is an autogenerated conversion function.
func Convert_seratos_Microservice_To_v1beta1_Microservice(in *seratos.Microservice, out *Microservice, s conversion.Scope) error {
	return autoConvert_seratos_Microservice_To_v1beta1_Microservice(in, out, s)
}

func autoConvert_v1beta1_MicroserviceConnection_To_seratos_MicroserviceConnection(in *MicroserviceConnection, out *seratos.MicroserviceConnection, s conversion.Scope) error {
	out.Service = in.Service
	out.Host = in.Host
	if err := Convert_v1beta1_ProxyPorts_To_seratos_ProxyPorts(&in.Ports, &out.Ports, s); err != nil {
		return err
	}
	out.URL = in.URL
	return nil
}

// Convert_v1beta1_MicroserviceConnection_To_seratos_MicroserviceConnection is an autogenerated conversion function.
func Convert_v1beta1_MicroserviceConnection_To_seratos_MicroserviceConnection(in *MicroserviceConnection, out *seratos.MicroserviceConnection, s conversion.Scope) error {
	return autoConvert_v1beta1_MicroserviceConnection_To_seratos_MicroserviceConnection(in, out, s)
}

func autoConvert_seratos_MicroserviceConnection_To_v1beta1_MicroserviceConnection(in *seratos.MicroserviceConnection, out *MicroserviceConnection, s conversion.Scope) error {
	out.Service = in.Service
	out.Host = in.Host
	if err := Convert_seratos_ProxyPorts_To_v1beta1_ProxyPorts(&in.Ports, &out.Ports, s); err != nil {
		return err
	}
	out.URL = in.URL
	return nil
}

// Convert_seratos_MicroserviceConnection_To_v1beta1_MicroserviceConnection is an autogenerated conversion function.
func Convert_seratos_MicroserviceConnection_To_v1beta1_MicroserviceConnection(in *seratos.MicroserviceConnection, out *MicroserviceConnection, s conversion.Scope) error {
	return autoConvert_seratos_MicroserviceConnection_To_v1beta1_MicroserviceConnection(in, out, s)
}

func autoConvert_v1beta1_MicroserviceInfo_To_seratos_MicroserviceInfo(in *MicroserviceInfo, out *seratos.MicroserviceInfo, s conversion.Scope) error {
	out.Connections = *(*[]seratos.MicroserviceConnection)(unsafe.Pointer(&in.Connections))
	if err := Convert_v1beta1_MicroserviceSecutiry_To_seratos_MicroserviceSecutiry(&in.Secutiry, &out.Secutiry, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_MicroserviceInfo_To_seratos_MicroserviceInfo is an autogenerated conversion function.
func Convert_v1beta1_MicroserviceInfo_To_seratos_MicroserviceInfo(in *MicroserviceInfo, out *seratos.MicroserviceInfo, s conversion.Scope) error {
	return autoConvert_v1beta1_MicroserviceInfo_To_seratos_MicroserviceInfo(in, out, s)
}

func autoConvert_seratos_MicroserviceInfo_To_v1beta1_MicroserviceInfo(in *seratos.MicroserviceInfo, out *MicroserviceInfo, s conversion.Scope) error {
	out.Connections = *(*[]MicroserviceConnection)(unsafe.Pointer(&in.Connections))
	if err := Convert_seratos_MicroserviceSecutiry_To_v1beta1_MicroserviceSecutiry(&in.Secutiry, &out.Secutiry, s); err != nil {
		return err
	}
	return nil
}

// Convert_seratos_MicroserviceInfo_To_v1beta1_MicroserviceInfo is an autogenerated conversion function.
func Convert_seratos_MicroserviceInfo_To_v1beta1_MicroserviceInfo(in *seratos.MicroserviceInfo, out *MicroserviceInfo, s conversion.Scope) error {
	return autoConvert_seratos_MicroserviceInfo_To_v1beta1_MicroserviceInfo(in, out, s)
}

func autoConvert_v1beta1_MicroserviceList_To_seratos_MicroserviceList(in *MicroserviceList, out *seratos.MicroserviceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]seratos.Microservice)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_MicroserviceList_To_seratos_MicroserviceList is an autogenerated conversion function.
func Convert_v1beta1_MicroserviceList_To_seratos_MicroserviceList(in *MicroserviceList, out *seratos.MicroserviceList, s conversion.Scope) error {
	return autoConvert_v1beta1_MicroserviceList_To_seratos_MicroserviceList(in, out, s)
}

func autoConvert_seratos_MicroserviceList_To_v1beta1_MicroserviceList(in *seratos.MicroserviceList, out *MicroserviceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Microservice)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_seratos_MicroserviceList_To_v1beta1_MicroserviceList is an autogenerated conversion function.
func Convert_seratos_MicroserviceList_To_v1beta1_MicroserviceList(in *seratos.MicroserviceList, out *MicroserviceList, s conversion.Scope) error {
	return autoConvert_seratos_MicroserviceList_To_v1beta1_MicroserviceList(in, out, s)
}

func autoConvert_v1beta1_MicroserviceSecutiry_To_seratos_MicroserviceSecutiry(in *MicroserviceSecutiry, out *seratos.MicroserviceSecutiry, s conversion.Scope) error {
	out.Policy = in.Policy
	out.Blocks = *(*[]string)(unsafe.Pointer(&in.Blocks))
	out.Allow = *(*[]string)(unsafe.Pointer(&in.Allow))
	return nil
}

// Convert_v1beta1_MicroserviceSecutiry_To_seratos_MicroserviceSecutiry is an autogenerated conversion function.
func Convert_v1beta1_MicroserviceSecutiry_To_seratos_MicroserviceSecutiry(in *MicroserviceSecutiry, out *seratos.MicroserviceSecutiry, s conversion.Scope) error {
	return autoConvert_v1beta1_MicroserviceSecutiry_To_seratos_MicroserviceSecutiry(in, out, s)
}

func autoConvert_seratos_MicroserviceSecutiry_To_v1beta1_MicroserviceSecutiry(in *seratos.MicroserviceSecutiry, out *MicroserviceSecutiry, s conversion.Scope) error {
	out.Policy = in.Policy
	out.Blocks = *(*[]string)(unsafe.Pointer(&in.Blocks))
	out.Allow = *(*[]string)(unsafe.Pointer(&in.Allow))
	return nil
}

// Convert_seratos_MicroserviceSecutiry_To_v1beta1_MicroserviceSecutiry is an autogenerated conversion function.
func Convert_seratos_MicroserviceSecutiry_To_v1beta1_MicroserviceSecutiry(in *seratos.MicroserviceSecutiry, out *MicroserviceSecutiry, s conversion.Scope) error {
	return autoConvert_seratos_MicroserviceSecutiry_To_v1beta1_MicroserviceSecutiry(in, out, s)
}

func autoConvert_v1beta1_MicroserviceSpec_To_seratos_MicroserviceSpec(in *MicroserviceSpec, out *seratos.MicroserviceSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.Sidecars = *(*[]string)(unsafe.Pointer(&in.Sidecars))
	out.Replicas = in.Replicas
	out.Enviroment = *(*[]seratos.Enviroment)(unsafe.Pointer(&in.Enviroment))
	if err := Convert_v1beta1_MicroserviceInfo_To_seratos_MicroserviceInfo(&in.Info, &out.Info, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_MicroserviceSpec_To_seratos_MicroserviceSpec is an autogenerated conversion function.
func Convert_v1beta1_MicroserviceSpec_To_seratos_MicroserviceSpec(in *MicroserviceSpec, out *seratos.MicroserviceSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_MicroserviceSpec_To_seratos_MicroserviceSpec(in, out, s)
}

func autoConvert_seratos_MicroserviceSpec_To_v1beta1_MicroserviceSpec(in *seratos.MicroserviceSpec, out *MicroserviceSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.Sidecars = *(*[]string)(unsafe.Pointer(&in.Sidecars))
	out.Replicas = in.Replicas
	out.Enviroment = *(*[]Enviroment)(unsafe.Pointer(&in.Enviroment))
	if err := Convert_seratos_MicroserviceInfo_To_v1beta1_MicroserviceInfo(&in.Info, &out.Info, s); err != nil {
		return err
	}
	return nil
}

// Convert_seratos_MicroserviceSpec_To_v1beta1_MicroserviceSpec is an autogenerated conversion function.
func Convert_seratos_MicroserviceSpec_To_v1beta1_MicroserviceSpec(in *seratos.MicroserviceSpec, out *MicroserviceSpec, s conversion.Scope) error {
	return autoConvert_seratos_MicroserviceSpec_To_v1beta1_MicroserviceSpec(in, out, s)
}

func autoConvert_v1beta1_ProxyPorts_To_seratos_ProxyPorts(in *ProxyPorts, out *seratos.ProxyPorts, s conversion.Scope) error {
	out.TCP = in.TCP
	out.HTTP2 = in.HTTP2
	return nil
}

// Convert_v1beta1_ProxyPorts_To_seratos_ProxyPorts is an autogenerated conversion function.
func Convert_v1beta1_ProxyPorts_To_seratos_ProxyPorts(in *ProxyPorts, out *seratos.ProxyPorts, s conversion.Scope) error {
	return autoConvert_v1beta1_ProxyPorts_To_seratos_ProxyPorts(in, out, s)
}

func autoConvert_seratos_ProxyPorts_To_v1beta1_ProxyPorts(in *seratos.ProxyPorts, out *ProxyPorts, s conversion.Scope) error {
	out.TCP = in.TCP
	out.HTTP2 = in.HTTP2
	return nil
}

// Convert_seratos_ProxyPorts_To_v1beta1_ProxyPorts is an autogenerated conversion function.
func Convert_seratos_ProxyPorts_To_v1beta1_ProxyPorts(in *seratos.ProxyPorts, out *ProxyPorts, s conversion.Scope) error {
	return autoConvert_seratos_ProxyPorts_To_v1beta1_ProxyPorts(in, out, s)
}

func autoConvert_v1beta1_Sidecar_To_seratos_Sidecar(in *Sidecar, out *seratos.Sidecar, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_SidecarSpec_To_seratos_SidecarSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Sidecar_To_seratos_Sidecar is an autogenerated conversion function.
func Convert_v1beta1_Sidecar_To_seratos_Sidecar(in *Sidecar, out *seratos.Sidecar, s conversion.Scope) error {
	return autoConvert_v1beta1_Sidecar_To_seratos_Sidecar(in, out, s)
}

func autoConvert_seratos_Sidecar_To_v1beta1_Sidecar(in *seratos.Sidecar, out *Sidecar, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_seratos_SidecarSpec_To_v1beta1_SidecarSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_seratos_Sidecar_To_v1beta1_Sidecar is an autogenerated conversion function.
func Convert_seratos_Sidecar_To_v1beta1_Sidecar(in *seratos.Sidecar, out *Sidecar, s conversion.Scope) error {
	return autoConvert_seratos_Sidecar_To_v1beta1_Sidecar(in, out, s)
}

func autoConvert_v1beta1_SidecarFiles_To_seratos_SidecarFiles(in *SidecarFiles, out *seratos.SidecarFiles, s conversion.Scope) error {
	out.Filepath = in.Filepath
	out.Description = in.Description
	out.Template = in.Template
	return nil
}

// Convert_v1beta1_SidecarFiles_To_seratos_SidecarFiles is an autogenerated conversion function.
func Convert_v1beta1_SidecarFiles_To_seratos_SidecarFiles(in *SidecarFiles, out *seratos.SidecarFiles, s conversion.Scope) error {
	return autoConvert_v1beta1_SidecarFiles_To_seratos_SidecarFiles(in, out, s)
}

func autoConvert_seratos_SidecarFiles_To_v1beta1_SidecarFiles(in *seratos.SidecarFiles, out *SidecarFiles, s conversion.Scope) error {
	out.Filepath = in.Filepath
	out.Description = in.Description
	out.Template = in.Template
	return nil
}

// Convert_seratos_SidecarFiles_To_v1beta1_SidecarFiles is an autogenerated conversion function.
func Convert_seratos_SidecarFiles_To_v1beta1_SidecarFiles(in *seratos.SidecarFiles, out *SidecarFiles, s conversion.Scope) error {
	return autoConvert_seratos_SidecarFiles_To_v1beta1_SidecarFiles(in, out, s)
}

func autoConvert_v1beta1_SidecarList_To_seratos_SidecarList(in *SidecarList, out *seratos.SidecarList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]seratos.Sidecar)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_SidecarList_To_seratos_SidecarList is an autogenerated conversion function.
func Convert_v1beta1_SidecarList_To_seratos_SidecarList(in *SidecarList, out *seratos.SidecarList, s conversion.Scope) error {
	return autoConvert_v1beta1_SidecarList_To_seratos_SidecarList(in, out, s)
}

func autoConvert_seratos_SidecarList_To_v1beta1_SidecarList(in *seratos.SidecarList, out *SidecarList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Sidecar)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_seratos_SidecarList_To_v1beta1_SidecarList is an autogenerated conversion function.
func Convert_seratos_SidecarList_To_v1beta1_SidecarList(in *seratos.SidecarList, out *SidecarList, s conversion.Scope) error {
	return autoConvert_seratos_SidecarList_To_v1beta1_SidecarList(in, out, s)
}

func autoConvert_v1beta1_SidecarSpec_To_seratos_SidecarSpec(in *SidecarSpec, out *seratos.SidecarSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.Files = *(*[]seratos.SidecarFiles)(unsafe.Pointer(&in.Files))
	out.Enviroment = *(*[]seratos.Enviroment)(unsafe.Pointer(&in.Enviroment))
	return nil
}

// Convert_v1beta1_SidecarSpec_To_seratos_SidecarSpec is an autogenerated conversion function.
func Convert_v1beta1_SidecarSpec_To_seratos_SidecarSpec(in *SidecarSpec, out *seratos.SidecarSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_SidecarSpec_To_seratos_SidecarSpec(in, out, s)
}

func autoConvert_seratos_SidecarSpec_To_v1beta1_SidecarSpec(in *seratos.SidecarSpec, out *SidecarSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.Files = *(*[]SidecarFiles)(unsafe.Pointer(&in.Files))
	out.Enviroment = *(*[]Enviroment)(unsafe.Pointer(&in.Enviroment))
	return nil
}

// Convert_seratos_SidecarSpec_To_v1beta1_SidecarSpec is an autogenerated conversion function.
func Convert_seratos_SidecarSpec_To_v1beta1_SidecarSpec(in *seratos.SidecarSpec, out *SidecarSpec, s conversion.Scope) error {
	return autoConvert_seratos_SidecarSpec_To_v1beta1_SidecarSpec(in, out, s)
}
