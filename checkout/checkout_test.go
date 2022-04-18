package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	testService := NewService()
	status, err := testService.Verify(10, "random details")
	assert.Equal(t, Successful, status)
	assert.NoError(t, err)
}
