#!/bin/bash -e

CURRENT_DIR=$(echo "$(pwd)/$line")
GEN_DIR=""
REPO_DIR="$CURRENT_DIR"

PROJECT_MODULE="github.com/Marcos30004347/k8s-custom-API-Server"
IMAGE_NAME="kubernetes-codegen:latest"

CUSTOM_RESOURCE_NAME="baz"
CUSTOM_RESOURCE_VERSION="v1beta1,v1beta1"

echo "Building codegen Docker image..."
docker build -f "${CURRENT_DIR}/hack/docker/codegen.dockerfile" \
             -t "${IMAGE_NAME}" \
             "${REPO_DIR}"


cmd0="./generate-groups.sh all \
    "$PROJECT_MODULE/pkg/generated" \
    "$PROJECT_MODULE/pkg/apis" \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

cmd1="./generate-internal-groups.sh "deepcopy,defaulter,conversion,informer,listers,client" \
    "$PROJECT_MODULE/pkg/generated" \
    "$PROJECT_MODULE/pkg/apis" \
    "$PROJECT_MODULE/pkg/apis" \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

echo "Generating client codes..."
echo "${REPO_DIR}"
docker run --rm \
           -v "${REPO_DIR}:/go/src/github.com/Marcos30004347/k8s-custom-API-Server" \
           "${IMAGE_NAME}" ls /go/src/github.com/Marcos30004347/k8s-custom-API-Server/pkg

docker run --rm \
           -v "${REPO_DIR}:/go/src/github.com/Marcos30004347/k8s-custom-API-Server" \
           "${IMAGE_NAME}" $cmd0

docker run --rm \
           -v "${REPO_DIR}:/go/src/github.com/Marcos30004347/k8s-custom-API-Server" \
           "${IMAGE_NAME}" $cmd1

# sudo chown $USER:$USER -R $REPO_DIR/pkg
