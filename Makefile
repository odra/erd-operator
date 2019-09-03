SHELL = /bin/bash
REG = quay.io
ORG = emergencyresponsedemo
IMAGE = erd-operator
TAG = latest
TEST_FOLDER = ./test/e2e
NS = erd

.PHONY: code/check
code/check:
	@diff -u <(echo -n) <(gofmt -d `find . -type f -name '*.go' -not -path "./vendor/*"`)

.PHONY: code/fix
code/fix:
	@gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*"`

.PHONY: image/build
image/build:
	@operator-sdk build $(REG)/$(ORG)/$(IMAGE):$(TAG)

.PHONY: image/build/test
image/build/test:
	operator-sdk build --enable-tests $(REG)/$(ORG)/$(IMAGE):$(TAG)

.PHONY: image/push
image/push:
	docker push $(REG)/$(ORG)/$(IMAGE):$(TAG)

.PHONY: test/unit
test/unit:
	@go test -v -race -coverprofile=coverage.out ./pkg/...

.PHONY: test/e2e/local
test/e2e:
	@operator-sdk test local ${TEST_FOLDER} --go-test-flags "-v"

.PHONY: cluster/prepare
cluster/prepare:
	@kubectl apply -f deploy/role.yaml -n ${NS}
	@kubectl apply -f deploy/role_binding.yaml -n ${NS}
	@kubectl apply -f deploy/service_account.yaml -n ${NS}
	@kubectl apply -f deploy/org_v1alpha1_graphback_crd.yaml -n ${NS}

.PHONY: cluster/deploy
cluster/deploy:
	@kubectl apply -f deploy/operator.yaml -n ${NS}