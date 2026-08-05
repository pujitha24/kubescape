package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"
)

func TestNewSecurityExceptionEventRecorderWithClient_NilClient(t *testing.T) {
	recorder, shutdown := newSecurityExceptionEventRecorderWithClient(nil)
	assert.Nil(t, recorder)
	assert.Nil(t, shutdown)
}

func TestNewSecurityExceptionEventRecorderWithClient_ReturnsShutdown(t *testing.T) {
	client := fake.NewSimpleClientset()

	recorder, shutdown := newSecurityExceptionEventRecorderWithClient(client)

	assert.NotNil(t, recorder)
	if assert.NotNil(t, shutdown) {
		// The caller must be able to stop the broadcaster's background
		// event-watcher goroutine; calling the closure must not panic.
		assert.NotPanics(t, shutdown)
	}
}
