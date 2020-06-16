PKG=github.com/capitalonline/cloud-controller-manager
IMAGE?=registry-bj.capitalonline.net/cck/cloud-controller-manager
VERSION=test15
CCM_DEPLOY_PATH=./deploy
CCM_KUSTOMIZATION_RELEASE_PATH=${CCM_DEPLOY_PATH}/overlays/release
CCM_KUSTOMIZATION_RELEASE_FILE=${CCM_KUSTOMIZATION_RELEASE_PATH}/kustomization.yaml
CCM_KUSTOMIZATION_TEST_PATH=${CCM_DEPLOY_PATH}/overlays/test
CCM_KUSTOMIZATION_TEST_FILE=${CCM_KUSTOMIZATION_RELEASE_PATH}/kustomization.yaml
.EXPORT_ALL_VARIABLES:

.PHONY: build
build:
	mkdir -p bin
	CGO_ENABLED=0 go build -o bin/cds-ccm ./cmd/

.PHONY: container-binary
container-binary:
	CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -mod vendor -o /cds-ccm ./cmd/

.PHONY: image-release
image-release:
	docker build -t $(IMAGE):$(VERSION) .

.PHONY: image-build
image:
	docker build -t $(IMAGE):latest .

.PHONY: release
release: image-release
	docker push $(IMAGE):$(VERSION)

.PHONY: sync-version
sync-version:
	sed -i.bak 's/newTag: .*/newTag: '${VERSION}'/g' ${CCM_KUSTOMIZATION_RELEASE_FILE} && rm ${CCM_KUSTOMIZATION_RELEASE_FILE}.bak

.PHONY: kustomize
kustomize:sync-version
	kubectl kustomize ${CCM_KUSTOMIZATION_RELEASE_PATH} > ${CCM_DEPLOY_PATH}/deploy.yaml

.PHONY: unit-test
unit-test:
	@echo "**************************** running unit test ****************************"
	go test -v -race ./pkg/ccm/...