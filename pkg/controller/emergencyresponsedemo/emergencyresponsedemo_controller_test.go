package emergencyresponsedemo

import (
	"github.com/integr8ly/erd-operator/pkg/apis/erdemo/v1alpha1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

func _assertStatus(t *testing.T, expected *v1alpha1.EmergencyResponseDemoStatus, current *v1alpha1.EmergencyResponseDemoStatus) {
	assert.Equal(t, expected.Status, current.Status)
	assert.Equal(t, expected.Type, current.Type)
	assert.Equal(t, *expected.Message, *current.Message)
	assert.Equal(t, *expected.Reason, *current.Reason)
	assert.NotNil(t, current.LastHeartbeatTime)
	assert.NotNil(t, current.LastTransitionTime)
}

func _strPointer(v string) *string {
	s := new(string)
	*s = v

	return s
}

func TestReconcileEmergencyResponseDemo_getSecret(t *testing.T) {
	cases := []struct{
		Name string
		Instance *v1alpha1.EmergencyResponseDemo
		Secret *v1.Secret
		ExpectedError bool
		Client func(objs []runtime.Object) client.Client
		Validate func(t *testing.T, instance *v1.Secret)
	}{
		{
			Name: "Should retrieve secret",
			Instance: &v1alpha1.EmergencyResponseDemo{
				ObjectMeta: v12.ObjectMeta{
					Name:                       "erd",
					Namespace:                  "default",
				},
				Spec:v1alpha1.EmergencyResponseDemoSpec{
					SecretName:      "erd-credentials",
				},
			},
			Secret: &v1.Secret{
				ObjectMeta: v12.ObjectMeta{
					Name:                       "erd-credentials",
					Namespace:                  "default",
				},
			},
			ExpectedError: false,
			Client: func(objs []runtime.Object) client.Client {
				cl := fake.NewFakeClient(objs...)

				return cl
			},
			Validate: func(t *testing.T, instance *v1.Secret) {
				assert.Equal(t, instance.Name, "erd-credentials")
				assert.Equal(t, instance.Namespace, "default")
			},
		},
		{
			Name: "Should fail to retrieve secret",
			Instance: &v1alpha1.EmergencyResponseDemo{
				ObjectMeta: v12.ObjectMeta{
					Name:                       "erd",
					Namespace:                  "default",
				},
				Spec:v1alpha1.EmergencyResponseDemoSpec{
					SecretName:      "erd-credentials",
				},
			},
			Secret: &v1.Secret{
				ObjectMeta: v12.ObjectMeta{
					Name:                       "erd-credentials-404",
					Namespace:                  "default",
				},
			},
			ExpectedError: true,
			Client: func(objs []runtime.Object) client.Client {
				cl := fake.NewFakeClient(objs...)

				return cl
			},
			Validate: func(t *testing.T, instance *v1.Secret) {
				assert.Equal(t, instance.Name, "")
				assert.Equal(t, instance.Namespace, "")
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			//setup objects
			instance := tc.Instance
			objs := []runtime.Object{instance, tc.Secret}
			//scheme setup
			s := scheme.Scheme
			s.AddKnownTypes(v1alpha1.SchemeGroupVersion, instance)
			//client
			cl := tc.Client(objs)
			//reconcile
			r := &ReconcileEmergencyResponseDemo{
				client: cl,
				scheme: s,
			}

			secret, err := r.getSecret(instance)
			if tc.ExpectedError && err == nil {
				assert.NotNil(t, err)
			}
			if !tc.ExpectedError && err != nil {
				assert.Nil(t, err)
			}

			tc.Validate(t, secret)
		})
	}
}

func TestReconcileEmergencyResponseDemo_setReadyStatus(t *testing.T) {
	cases := []struct{
		Name string
		Instance *v1alpha1.EmergencyResponseDemo
		ExpectedStatus *v1alpha1.EmergencyResponseDemoStatus
		Client func(objs []runtime.Object) client.Client
	}{
		{
			Name: "Status should assert as ready",
			Instance: &v1alpha1.EmergencyResponseDemo{},
			ExpectedStatus:&v1alpha1.EmergencyResponseDemoStatus{
				Type:               v1alpha1.EmergencyResponseDemoReady,
				Status:             v1.ConditionTrue,
				Reason:             _strPointer("Ready"),
				Message:            _strPointer("Emergency Response Demo is ready"),
			},
			Client: func(objs []runtime.Object) client.Client {
				cl := fake.NewFakeClient(objs...)

				return cl
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			//setup objects
			instance := tc.Instance
			objs := []runtime.Object{instance}
			//scheme setup
			s := scheme.Scheme
			s.AddKnownTypes(v1alpha1.SchemeGroupVersion, instance)
			//client
			cl := tc.Client(objs)
			//reconcile
			r := &ReconcileEmergencyResponseDemo{
				client: cl,
				scheme: s,
			}
			err := r.setReadyStatus(instance)
			assert.Nil(t, err)

			_assertStatus(t, tc.ExpectedStatus, &instance.Status)
		})
	}
}

func TestReconcileEmergencyResponseDemo_setSecretErrorStatus(t *testing.T) {
	cases := []struct{
		Name string
		Instance *v1alpha1.EmergencyResponseDemo
		ExpectedStatus *v1alpha1.EmergencyResponseDemoStatus
		Client func(objs []runtime.Object) client.Client
	}{
		{
			Name: "Status should assert as secret not found error status",
			Instance: &v1alpha1.EmergencyResponseDemo{},
			ExpectedStatus:&v1alpha1.EmergencyResponseDemoStatus{
				Type:               v1alpha1.EmergencyResponseDemoError,
				Status:             v1.ConditionTrue,
				Reason:             _strPointer("ERDSecretNotFound"),
				Message:            _strPointer("Could not find an ERD secret in this namespace"),
			},
			Client: func(objs []runtime.Object) client.Client {
				cl := fake.NewFakeClient(objs...)

				return cl
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			//setup objects
			instance := tc.Instance
			objs := []runtime.Object{instance}
			//scheme setup
			s := scheme.Scheme
			s.AddKnownTypes(v1alpha1.SchemeGroupVersion, instance)
			//client
			cl := tc.Client(objs)
			//reconcile
			r := &ReconcileEmergencyResponseDemo{
				client: cl,
				scheme: s,
			}
			err := r.setSecretErrorStatus(instance)
			assert.Nil(t, err)

			_assertStatus(t, tc.ExpectedStatus, &instance.Status)
		})
	}
}

//func TestReconcileEmergencyResponseDemo_Reconcile(t *testing.T) {
//	cases := []struct{
//		Name string
//		Instance *v1alpha1.EmergencyResponseDemo
//		Client func(objs []runtime.Object) client.Client
//		ReconcileError string
//		ReconcileResult reconcile.Result
//		Status func() *v1alpha1.EmergencyResponseDemoStatus
//	}{
//		{
//			Name: "Should stop and reconcile at secret error status",
//			Instance: &v1alpha1.EmergencyResponseDemo{
//				ObjectMeta: v12.ObjectMeta{
//					Name:      "erd",
//					Namespace: "default",
//				},
//				Spec:       v1alpha1.EmergencyResponseDemoSpec{
//					SecretName:      "erd-credentials",
//					SelfSignedCerts: true,
//					SubDomain:       "domain.com",
//					MasterUrl:       "domain.com",
//				},
//			},
//			Client: func(objs []runtime.Object) client.Client {
//				cl := fake.NewFakeClient(objs...)
//
//				return cl
//			},
//			ReconcileError: "secrets \"erd-credentials\" not found",
//			ReconcileResult:reconcile.Result{
//				Requeue:      true,
//				RequeueAfter: 0,
//			},
//			Status: func() *v1alpha1.EmergencyResponseDemoStatus {
//				reason := new(string)
//				*reason = "ERDSecretNotFound"
//
//				message := new(string)
//				*message = "Could not find an ERD secret in this namespace"
//
//				return &v1alpha1.EmergencyResponseDemoStatus{
//					Type:               v1alpha1.EmergencyResponseDemoError,
//					Status:             v1.ConditionTrue,
//					Reason:             reason,
//					Message:            message,
//				}
//			},
//		},
//		{
//			Name: "Should finish reconcile loop with success",
//			Instance: &v1alpha1.EmergencyResponseDemo{
//				ObjectMeta: v12.ObjectMeta{
//					Name:      "erd",
//					Namespace: "default",
//				},
//				Spec:       v1alpha1.EmergencyResponseDemoSpec{
//					SecretName:      "erd-credentials",
//					SelfSignedCerts: true,
//					SubDomain:       "domain.com",
//					MasterUrl:       "domain.com",
//				},
//			},
//			Client: func(objs []runtime.Object) client.Client {
//				secret := &v1.Secret{
//					ObjectMeta: v12.ObjectMeta{
//						Name:                       "erd-credentials",
//						Namespace:                  "default",
//					},
//					Data:       nil,
//					StringData: nil,
//					Type:       v1.SecretTypeOpaque,
//				}
//				objs = append(objs, secret)
//
//				cl := fake.NewFakeClient(objs...)
//
//				return cl
//			},
//			ReconcileError: "",
//			ReconcileResult:reconcile.Result{
//				Requeue:      false,
//				RequeueAfter: 0,
//			},
//			Status: func() *v1alpha1.EmergencyResponseDemoStatus {
//				reason := new(string)
//				*reason = "Ready"
//
//				message := new(string)
//				*message = "Emergency Response Demo is ready"
//
//				return &v1alpha1.EmergencyResponseDemoStatus{
//					Type:               v1alpha1.EmergencyResponseDemoReady,
//					Status:             v1.ConditionTrue,
//					Reason:             reason,
//					Message:            message,
//				}
//			},
//		},
//	}
//
//	for _, tc := range cases {
//		tc := tc
//		t.Run(tc.Name, func(t *testing.T) {
//			//setup objects
//			instance := tc.Instance
//			objs := []runtime.Object{instance}
//			//scheme setup
//			s := scheme.Scheme
//			s.AddKnownTypes(v1alpha1.SchemeGroupVersion, instance)
//			//client setup
//			cl := tc.Client(objs)
//			//reconcile setup
//			r := &ReconcileEmergencyResponseDemo{
//				client: cl,
//				scheme: s,
//			}
//			req := reconcile.Request{
//				NamespacedName: types.NamespacedName{
//					Namespace: "default",
//					Name:      "erd",
//				},
//			}
//			res, err := r.Reconcile(req)
//			if tc.ReconcileError != "" {
//				assert.EqualError(t, err, tc.ReconcileError)
//			}
//			assert.Equal(t, tc.ReconcileResult, res)
//
//			//read updated erd cr
//			readInstance :=  &v1alpha1.EmergencyResponseDemo{}
//			err = r.client.Get(context.TODO(), req.NamespacedName, readInstance)
//			assert.Nil(t, err)
//
//			//status validation
//			expectedStatus := tc.Status()
//			_assertStatus(t, expectedStatus, &readInstance.Status)
//		})
//	}
//}
