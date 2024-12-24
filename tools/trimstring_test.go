package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimstring(t *testing.T) {
	s := "An Efficient and Secure Dynamic Auditing Protocol for Data Storage in Cloud Computing"
	assert.Equal(t, trimString(s), "AnEfficientandSecureDynamicAuditingProtocolforDataStorageinCloudComputing")
}
