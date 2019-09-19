package e2e

import (
	testCtx "github.com/integr8ly/erd-operator/test/api/ctx"
	"github.com/integr8ly/erd-operator/test/api/meta"
	"testing"
)

func TestSecretNotFound(t *testing.T) {
	//var err error

	waitOpts := meta.WaitOpts{
		RetryInterval: meta.DefaultRetryInterval,
		Timeout:       meta.DefaultTimeout,
	}

	ctx := testCtx.PrepareContext(t, waitOpts)
	defer ctx.Cleanup()
}
