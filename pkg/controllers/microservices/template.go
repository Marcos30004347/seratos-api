package controller

import (
	"github.com/Marcos30004347/seratos-api/pkg/apis/seratos/v1beta1"
	"github.com/flosch/pongo2/v4"
)

func parseFileContent(raw string, microservice *v1beta1.Microservice) string {
	tpl, err := pongo2.FromString(raw)

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	out, err := tpl.Execute(pongo2.Context{"Microservice": microservice})

	if err != nil {
		panic(err)
	}

	return out
}
