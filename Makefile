APP=car-user
ARCH?=amd64
REGISTRY?=
GIT_COMMIT:=$(shell git rev-parse --short HEAD)
ALL_ARCHITECTURES=amd64 arm arm64 ppc64le s390x
GOPROXY="https://goproxy.cn,direct"
all: build

# Build Rules
# -----------
src_deps=$(shell find -type f -name "*.go")
build: $(src_deps)
	GO111MODULE=on GOPROXY=$(GOPROXY) GOARCH=$(ARCH) go build -mod=vendor -o deploy/$(APP)

# Image Rules
# -----------
container: container-$(ARCH)

container-%:
	docker build -t $(REGISTRY)/$(APP):$(GIT_COMMIT) -f deploy/docker/Dockerfile deploy/

push:
	docker tag $(REGISTRY)/$(APP):$(GIT_COMMIT) $(REGISTRY)/$(APP):v1.0.0
	docker push $(REGISTRY)/$(APP):v1.0.0
