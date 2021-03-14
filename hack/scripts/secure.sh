# !/bin/bash

openssl genrsa -out tls.key 4096
openssl req -new -sha256 -out tls.csr -key tls.key -config hack/config/ssl.conf 
openssl x509 -req -sha256 -days 3650 -in tls.csr -signkey tls.key -out tls.crt -extensions req_ext -extfile hack/config/ssl.conf 

# Get template cpmtemt
secret_manifest=$(cat ./manifests/templates/cert-secret.yaml.template)
apiservice_manifest=$(cat ./manifests/templates/v1beta1-apiservice.yaml.template)

# Get cert and key values
cert=$(cat ./tls.crt | base64)
key=$(cat ./tls.key | base64)

# Get secret manifest content
secret_manifest=$(echo "${secret_manifest//CERT/$cert}")
secret_manifest=$(echo "${secret_manifest//KEY/$key}")

# Get apiservice manifest content
apiservice_manifest=$(echo "${apiservice_manifest//CERT/$cert}")
apiservice_manifest=$(echo "${apiservice_manifest//KEY/$key}")

if [ -d ./manifests/deploy/secure ]; then rm -rf ./manifests/deploy/secure; fi;

mkdir -p ./manifests/deploy/secure

echo "$secret_manifest" >> ./manifests/deploy/secure/cert-secret.yaml
echo "$apiservice_manifest" >> ./manifests/deploy/secure/v1beta1-apiservice.yaml


cp ./manifests/templates/auth-delegator.yaml ./manifests/deploy/secure
cp ./manifests/templates/ns.yaml ./manifests/deploy/secure
cp ./manifests/templates/auth-delegator.yaml ./manifests/deploy/secure
cp ./manifests/templates/auth-reader.yaml ./manifests/deploy/secure
cp ./manifests/templates/deployment.yaml ./manifests/deploy/secure
cp ./manifests/templates/rbac-bind.yaml ./manifests/deploy/secure
cp ./manifests/templates/rbac.yaml ./manifests/deploy/secure
cp ./manifests/templates/sa.yaml ./manifests/deploy/secure
cp ./manifests/templates/service.yaml ./manifests/deploy/secure