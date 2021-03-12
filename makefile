deps:
	GO111MODULE=on go mod tidy  

start:
	chmod 777 ./hack/scripts/minikube-startup.sh
	./hack/scripts/minikube-startup.sh

compile: deps
	CGO_ENABLED=0 GOOS=linux go build .

stop:
	chmod 777 ./hack/scripts/minikube-shutdown.sh
	./hack/scripts/minikube-shutdown.sh

build:
	docker build -f ./hack/docker/seratos-api.dockerfile -t marcos30004347/seratos-api .

codegen:
	chmod 777 ./hack/scripts/codegen.sh
	./hack/scripts/codegen.sh

deploy:
	kubectl apply -f ./artifacts/deploy/ns.yaml
	kubectl apply -f ./artifacts/deploy/

undeploy:
	kubectl delete -f ./artifacts/deploy/

run:
	go run . --etcd-servers localhost:2379 --authentication-kubeconfig ~/.kube/config --authorization-kubeconfig ~/.kube/config --kubeconfig ~/.kube/config

sidecar:
	kubectl apply -f ./artifacts/example/sidecar.yaml

microservice:
	kubectl apply -f ./artifacts/example/microservice.yaml