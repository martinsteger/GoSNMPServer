package GoSNMPServer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelper_oidLess(t *testing.T) {
	assert.True(t, oidLess("1.2", "1.2.3"))
	assert.False(t, oidLess("1.2.3", "1.2"))
	assert.True(t, oidLess("1.2.100000", "1.2.100001"))
	assert.False(t, oidLess("1.2.100000", "1.2.100000"))
	assert.False(t, oidLess("1.2.100001", "1.2.100000"))

}

func TestHelper_IsValidObjectIdentifier(t *testing.T) {
	assert.True(t, IsValidObjectIdentifier("1.2.3.4.5"))
	assert.True(t, IsValidObjectIdentifier(".1.2.3.4.5"))
	assert.False(t, IsValidObjectIdentifier(""))
	assert.False(t, IsValidObjectIdentifier("asdfdasf"))
	assert.False(t, IsValidObjectIdentifier("1..2.3.4.5"))
}
