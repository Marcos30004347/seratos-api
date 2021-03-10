package apiserver

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Marcos30004347/seratos-api/pkg/apis/seratos"
	"github.com/Marcos30004347/seratos-api/pkg/apis/seratos/install"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/registry/rest"

	customregistry "github.com/Marcos30004347/seratos-api/pkg/registry"

	// storages
	microservicesstorage "github.com/Marcos30004347/seratos-api/pkg/registry/seratos/microservices"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

var (
	// Scheme defines methods for serializing and deserializing API objects.
	Scheme = runtime.NewScheme()

	// Codecs provides methods for retrieving codecs and serializers for specific
	// versions and content types.
	Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
	install.Install(Scheme)

	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})

	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

type ExtraConfig struct {
	// Place your custom config here.
}

type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}

// CustomServer contains state for a Kubernetes custom api server.
type CustomServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}

type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of
	// this package.
	*completedConfig
}

func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{
		cfg.GenericConfig.Complete(),
		&cfg.ExtraConfig,
	}

	c.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	return CompletedConfig{&c}
}

func (c CompletedConfig) New() (*CustomServer, error) {
	genericServer, err := c.GenericConfig.New("seratos-api", genericapiserver.NewEmptyDelegate())

	if err != nil {
		return nil, err
	}

	s := &CustomServer{
		GenericAPIServer: genericServer,
	}

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(seratos.GroupName, Scheme, metav1.ParameterCodec, Codecs)

	// NewREST from the registry/etcd.go
	v1beta1storage := map[string]rest.Storage{}
	v1beta1storage["microservices"] = customregistry.RESTInPeace(microservicesstorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter))

	apiGroupInfo.VersionedResourcesStorageMap["v1beta1"] = v1beta1storage

	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}

	return s, nil
}
