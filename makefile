deps:
	GO111MODULE=on go mod tidy  

start:
	chmod 777 ./hack/scripts/minikube-startup.sh
	./hack/scripts/minikube-startup.sh

compile: deps
	CGO_ENABLED=0 GOOS=linux go build .

stop:
	./hack/scripts/minikube-shutdown.sh

build:
	docker build -f ./hack/docker/seratos-api.dockerfile -t marcos30004347/seratos-api .

codegen:
	./hack/scripts/codegen.sh

deploy:
	kubectl apply -f ./manifests/deploy/ns.yaml
	kubectl apply -f ./manifests/deploy/

undeploy:
	kubectl delete -f ./manifests/deploy/

deploy-secure:
	kubectl apply -f ./manifests/deploy/secure/ns.yaml
	kubectl apply -f ./manifests/deploy/secure/

undeploy-secure:
	kubectl delete -f ./manifests/deploy/secure/

run:
	go run . --etcd-servers localhost:2379 --authentication-kubeconfig ~/.kube/config --authorization-kubeconfig ~/.kube/config --kubeconfig ~/.kube/config

sidecar:
	kubectl apply -f ./manifests/example/sidecar.yaml

microservice:
	kubectl apply -f ./manifests/example/microservice.yaml

secure:
	./hack/scripts/secure.sh
