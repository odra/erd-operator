package metahelper

import (
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type MetaHelper struct {
	obj metav1.Object
}

func New(obj runtime.Object) (*MetaHelper, error) {
	mh := &MetaHelper{}

	err := mh.Reload(obj)
	if err != nil {
		return nil, err
	}

	return mh, nil
}

func (mh *MetaHelper) Reload(obj runtime.Object) error {
	metaobj, err := meta.Accessor(obj)
	if err != nil {
		return err
	}

	mh.obj = metaobj

	return nil
}

func (mh *MetaHelper) Finalizers() []string {
	return mh.obj.GetFinalizers()
}

func (mh *MetaHelper) HasFinalizer(name string) bool {
	for _, finalizer := range mh.Finalizers() {
		if finalizer == name {
			return true
		}
	}

	return false
}

func (mh *MetaHelper) AddFinalizer(name string) {
	if mh.HasFinalizer(name) {
		return
	}

	finalizers := append(mh.Finalizers(), name)
	mh.obj.SetFinalizers(finalizers)
}

func (mh *MetaHelper) RemoveFinalizer(name string) {
	if !mh.HasFinalizer(name) {
		return
	}

	finalizers := make([]string, 0)
	for _, finalizer := range mh.Finalizers() {
		if finalizer == name {
			continue
		}
		finalizers = append(finalizers, finalizer)
	}

	mh.obj.SetFinalizers(finalizers)
}


