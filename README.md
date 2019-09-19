# Emergency Response Demo Operator

This operator deploys and manages Emergency Response Demo instances.

## Components

High level components provisioned/managed by this operator:

* kafka
    * cluster
    * topics
* pgadmin4
* postgresql
* infinispan
* emergency response
    * incident process
    * incident service
    * process viewer
    * incident priority service
    * mission service
    * process service
    * responder service
    * incident priority service
    * mission service
    * process service
    * responder service
    * incident priority service
    * responder simulator service
    * disaster simulator service
    * emergency console

## Dependencies
    
### Operators

Third party operators this operator depends on:

* infinispan: https://github.com/infinispan/infinispan-operator
* integreatly: https://github.com/integr8ly/integreatly-operator
* kafka: https://github.com/strimzi/strimzi-kafka-operator
* monitoring: https://github.com/integr8ly/application-monitoring-operator
* postgresql: https://github.com/zalando/postgres-operator

## Deployment

### Standalone

Create a namespace for the operator to be deployed:

```bash
kubectl create ns erd
```

Apply the required pre-deployment resources:

```bash
kubectl apply \
-f deploy/role.yaml \
-f deploy/role_binding.yaml \
-f deploy/service_account.yaml \
-f deploy/crds/org_v1alpha1_graphback_crd.yaml \
-n erd
```

Deploy the operator and wait for its readiness:

```bash
kubectl apply -f deploy/operator.yaml -n erd
```

Deploy an erd operator secret, see [deploy/erd-secret.yaml](./deploy/erd-secret.yaml) for reference:

```bash
kubectl apply -f erd-secret.yaml-n erd
```

Deploy an erd custom resource, see [deploy/crds/org_v1alpha1_graphback_crd.yaml](.deploy/crds/org_v1alpha1_graphback_crd.yaml) for reference:

```bash
kubctl apply -f erd.yaml -n erd
```

You can now check your erd resource status by running:

```bash
kubectl get erd/demo  -o jsonpath='{.status.type}'-n erd
```

Status should be `Ready` once everything is properly deployed and configured.

### OLM

TBD

## LICENSE

Licensed under apache license, see [LICENSE](./LICENSE) for more details.
