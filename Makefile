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

.PHONY: image/push
image/push:
	docker push $(REG)/$(ORG)/$(IMAGE):$(TAG)

.PHONY: test/unit
test/unit:
	@go test -v -race -coverprofile=coverage.out ./pkg/...

.PHONY: test/e2e/local
test/e2e/local: image/build image/push
	operator-sdk test local ${TEST_FOLDER} --go-test-flags "-v" --image $(REG)/$(ORG)/$(IMAGE):$(TAG)

.PHONY: cluster/prepare
cluster/prepare:
	 @kubectl apply -f deploy/role.yaml  -f deploy/service_account.yaml -f deploy/role_binding.yaml  -n ${NS}
	 @kubectl apply -f deploy/crds/*crd.yaml

.PHONY: cluster/deploy/operator
cluster/deploy/operator:
	@kubectl apply -f deploy/operator.yaml -n ${NS}

.PHONY: cluster/delete/operator
cluster/delete/operator:
	@kubectl delete -f deploy/operator.yaml -n ${NS}

.PHONY: cluster/deploy/erd
cluster/deploy/erd:
	@kubectl apply -f deploy/erd-secret.yaml -n ${NS}
	@kubectl apply -f deploy/crds/erdemo_v1alpha1_emergencyresponsedemo_cr.yaml -n ${NS}

.PHONY: cluster/delete/erd
cluster/delete/erd:
	@kubectl delete -f deploy/erd-secret.yaml -n ${NS} || true
	@kubectl delete -f deploy/crds/erdemo_v1alpha1_emergencyresponsedemo_cr.yaml -n ${NS}

.PHONY: cluster/delete
cluster/delete: cluster/delete/erd cluster/delete/operator
