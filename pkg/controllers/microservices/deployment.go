package controller

import (
	"fmt"
	"strings"

	v1beta1 "github.com/Marcos30004347/seratos-api/pkg/apis/seratos/v1beta1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// createDeployment creates a new Deployment for a Microservice resource. It also sets
// the appropriate OwnerReferences on the resource so handleObject can discover
// the Microservice resource that 'owns' it.
func createDeployment(microservice *v1beta1.Microservice, sidecars []*v1beta1.Sidecar) *appsv1.Deployment {
	initContainers := []corev1.Container{}
	volumes := []corev1.Volume{}

	// Build needed Volumes and mounts
	for _, sidecar := range sidecars {
		used := make(map[string]bool)
		it := 0
		for _, file := range sidecar.Spec.Files {
			path := strings.Split(file.Filepath, "/")
			path = path[:len(path)-1]
			filepath := strings.Join(path, "/")

			if !used[filepath] {
				used[filepath] = true

				volumes = append(volumes, corev1.Volume{
					Name: "volume" + sidecar.Name + fmt.Sprint(it),
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				})
				it++
			}
		}
	}

	// Built the init Containers
	for _, sidecar := range sidecars {
		volumeMounts := []corev1.VolumeMount{}
		echos := ""
		used := make(map[string]bool)
		it := 0

		for _, file := range sidecar.Spec.Files {
			path := strings.Split(file.Filepath, "/")
			path = path[:len(path)-1]
			filepath := strings.Join(path, "/")

			filecontent := parseFileContent(file.Template, microservice)
			echos += "echo '" + filecontent + "' >> " + file.Filepath + ";"

			if !used[filepath] {
				used[filepath] = true

				volumeMounts = append(volumeMounts, corev1.VolumeMount{
					Name:      "volume" + sidecar.Name + fmt.Sprint(it),
					MountPath: filepath,
				})
				it++
			}
		}

		initContainers = append(initContainers, corev1.Container{
			Name:         "init-" + sidecar.Name,
			Image:        "alpine:latest",
			Command:      []string{"/bin/sh", "-c"},
			Args:         []string{echos},
			VolumeMounts: volumeMounts,
		})
	}

	// Get microservice container
	containers := []corev1.Container{
		{
			Name:  microservice.Name,
			Image: microservice.Spec.Image,
		},
	}

	// include sidecars
	for _, sidecar := range sidecars {
		volumeMounts := []corev1.VolumeMount{}
		used := make(map[string]bool)
		it := 0
		for _, file := range sidecar.Spec.Files {
			path := strings.Split(file.Filepath, "/")
			path = path[:len(path)-1]
			filepath := strings.Join(path, "/")

			if !used[filepath] {
				used[filepath] = true

				volumeMounts = append(volumeMounts, corev1.VolumeMount{
					Name:      "volume" + sidecar.Name + fmt.Sprint(it),
					MountPath: filepath,
				})
				it++
			}
		}

		containers = append(containers, corev1.Container{
			Name:         sidecar.Name,
			Image:        sidecar.Spec.Image,
			VolumeMounts: volumeMounts,
			Args:         sidecar.Spec.Args,
			Command:      sidecar.Spec.Command,
		})
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      microservice.Name,
			Namespace: microservice.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(microservice, v1beta1.SchemeGroupVersion.WithKind("Microservice")),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &microservice.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: microservice.Labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: microservice.Labels,
				},
				Spec: corev1.PodSpec{
					Volumes:        volumes,
					InitContainers: initContainers,
					Containers:     containers,
				},
			},
		},
	}
}
