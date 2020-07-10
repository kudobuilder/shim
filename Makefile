
.PHONY: build-crd-controller
build-crd-controller:
	 docker build . -f Dockerfile.crd -t kudobuilder/crd-controller:0.0.1-alpha

.PHONY: push-crd-controller
push-crd-controller: build-crd-controller
	docker push kudobuilder/crd-controller:0.0.1-alpha

.PHONY: build-shim-controller
build-shim-controller:
	 docker build . -f Dockerfile.shim -t kudobuilder/shim-controller:0.0.1-alpha

.PHONY: push-shim-controller
push-shim-controller: build-shim-controller
	docker push kudobuilder/shim-controller:0.0.1-alpha

.PHONY: build.all
build.all: push-crd-controller build-shim-controller

.PHONY: push.all
push.all: push-crd-controller push-shim-controller

generate:
ifneq ($(shell go list -f '{{.Version}}' -m sigs.k8s.io/controller-tools), $(shell controller-gen --version 2>/dev/null | cut -b 10-))
	@echo "(Re-)installing controller-gen. Current version:  $(controller-gen --version 2>/dev/null | cut -b 10-). Need $(go list -f '{{.Version}}' -m sigs.k8s.io/controller-tools)"
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@$$(go list -f '{{.Version}}' -m sigs.k8s.io/controller-tools)
endif
	controller-gen crd paths=./shim-controller/pkg/apis/... output:crd:dir=config/crds output:stdout
ifeq (, $(shell which go-bindata))
	go get github.com/go-bindata/go-bindata/go-bindata@$$(go list -f '{{.Version}}' -m github.com/go-bindata/go-bindata)
endif
	./hack/update_codegen.sh
