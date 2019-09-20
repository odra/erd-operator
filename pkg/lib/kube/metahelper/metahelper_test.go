package metahelper

import (
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"testing"
)

func _createPod() *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: v1.ObjectMeta{
			Name:      "mypod",
			Namespace: "default",
		},
	}
}

func TestNew(t *testing.T) {
	cases := []struct {
		Name        string
		Obj         runtime.Object
		ExpectError bool
	}{
		{
			Name:        "Should instantiate metahelper",
			Obj:         _createPod(),
			ExpectError: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			mh, err := New(tc.Obj)

			if tc.ExpectError && err == nil {
				assert.NotNil(t, err)
			}

			if !tc.ExpectError && err != nil {
				assert.Nil(t, err)
			}

			assert.NotNil(t, mh.obj)
		})
	}
}

func TestMetaHelper_Reload(t *testing.T) {
	cases := []struct {
		Name        string
		Obj         runtime.Object
		ExpectError bool
		Validate    func(t *testing.T, mh *MetaHelper)
	}{
		{
			Name:        "Should reload metahelper",
			Obj:         _createPod(),
			ExpectError: false,
			Validate: func(t *testing.T, mh *MetaHelper) {
				assert.NotNil(t, mh)
				assert.NotNil(t, mh.obj)
				assert.IsType(t, &corev1.Pod{}, mh.obj)
			},
		},
		{
			Name:        "Should reload different kinds in metahelper",
			Obj:         _createPod(),
			ExpectError: false,
			Validate: func(t *testing.T, mh *MetaHelper) {
				assert.NotNil(t, mh)
				assert.NotNil(t, mh.obj)
				assert.IsType(t, &corev1.Pod{}, mh.obj)

				err := mh.Reload(&corev1.Namespace{})
				assert.Nil(t, err)

				assert.NotNil(t, mh)
				assert.NotNil(t, mh.obj)
				assert.IsType(t, &corev1.Namespace{}, mh.obj)
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			mh := &MetaHelper{}
			err := mh.Reload(tc.Obj)

			if tc.ExpectError && err == nil {
				assert.NotNil(t, err)
			}

			if !tc.ExpectError && err != nil {
				assert.Nil(t, err)
			}

			tc.Validate(t, mh)

		})
	}
}

func TestMetaHelper_Finalizers(t *testing.T) {
	cases := []struct {
		Name            string
		Helper          func(t *testing.T) *MetaHelper
		ExpectFinalizer []string
	}{
		{
			Name: "Should retrieve and assert finalizers",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()
				pod.Finalizers = []string{"f1", "f2"}

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectFinalizer: []string{"f1", "f2"},
		},
		{
			Name: "Should retrieve empty finalizer list",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()
				pod.Finalizers = []string{}

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectFinalizer: []string{},
		},
		{
			Name: "Should retrieve empty finalizer list if unset",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectFinalizer: make([]string, 0),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			helper := tc.Helper(t)

			assert.Equal(t, tc.ExpectFinalizer, helper.Finalizers())

		})
	}
}

func TestMetaHelper_HasFinalizer(t *testing.T) {
	cases := []struct {
		Name         string
		Finalizer    string
		Helper       func(t *testing.T) *MetaHelper
		ExpectResult bool
	}{
		{
			Name:      "Should find finalizer",
			Finalizer: "f2",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()
				pod.Finalizers = []string{"f1", "f2", " f3"}

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectResult: true,
		},
		{
			Name:      "Should not find finalizer",
			Finalizer: "f10",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()
				pod.Finalizers = []string{"f1", "f2", " f3"}

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectResult: false,
		},
		{
			Name:      "Should not find finalizer if unset",
			Finalizer: "f2",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectResult: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			helper := tc.Helper(t)

			assert.Equal(t, tc.ExpectResult, helper.HasFinalizer(tc.Finalizer))
		})
	}
}

func TestMetaHelper_AddFinalizer(t *testing.T) {
	cases := []struct {
		Name         string
		Finalizer    string
		Helper       func(t *testing.T) *MetaHelper
		ExpectResult []string
	}{
		{
			Name:      "Should add finalizer",
			Finalizer: "f3",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()
				pod.Finalizers = []string{"f1", "f5", "f2"}

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectResult: []string{"f1", "f5", "f2", "f3"},
		},
		{
			Name:      "Should not add existing finalizer",
			Finalizer: "f1",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()
				pod.Finalizers = []string{"f1", "f5", "f2"}

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectResult: []string{"f1", "f5", "f2"},
		},
		{
			Name:      "Should add finalizer in unset finalizer list",
			Finalizer: "f1",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectResult: []string{"f1"},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			helper := tc.Helper(t)
			helper.AddFinalizer(tc.Finalizer)

			assert.Equal(t, tc.ExpectResult, helper.obj.GetFinalizers())
		})
	}
}

func TestMetaHelper_RemoveFinalizer(t *testing.T) {
	cases := []struct {
		Name         string
		Finalizer    string
		Helper       func(t *testing.T) *MetaHelper
		ExpectResult []string
	}{
		{
			Name:      "Should remove finalizer",
			Finalizer: "f2",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()
				pod.Finalizers = []string{"f1", "f5", "f2"}

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectResult: []string{"f1", "f5"},
		},
		{
			Name:      "Should not fail if finalizer does not exist",
			Finalizer: "f50",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()
				pod.Finalizers = []string{"f1", "f5", "f2"}

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectResult: []string{"f1", "f5", "f2"},
		},
		{
			Name:      "Should not fail if finalizer list is not set",
			Finalizer: "f2",
			Helper: func(t *testing.T) *MetaHelper {
				pod := _createPod()

				accessor, err := meta.Accessor(pod)
				assert.Nil(t, err)

				mh := &MetaHelper{}
				mh.obj = accessor

				return mh
			},
			ExpectResult: nil,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			helper := tc.Helper(t)
			helper.RemoveFinalizer(tc.Finalizer)

			assert.Equal(t, tc.ExpectResult, helper.obj.GetFinalizers())
		})
	}
}
