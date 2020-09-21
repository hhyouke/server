package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var invCode *InvCode

func init() {
	invCode = &InvCode{
		base:    "HVE8S2DZX9C7P5IK3MJUAR4WYLTN6BGQ",
		decimal: 32,
		pad:     "F",
		len:     6,
	}
}

func TestIDToCode(t *testing.T) {
	id := int64(37232459628150785)
	code := invCode.IDToCode(id)
	t.Logf("%v\n", code)
}

func TestCodeToID(t *testing.T) {
	assert := assert.New(t)
	code := "VHHHS2DWDEVV"
	id := invCode.CodeToID(code)
	assert.Equal(id, int64(37232459628150785))
	// t.Logf("%v\n", id)
}

func TestIDToCustomCode(t *testing.T) {
	inst := NewCustomInvCode("F6BGQ2DZX9C7P5IK3MJUAR4WYLTNVE8S", "H")
	id := int64(37232459628150785)
	code := inst.IDToCode(id)
	t.Logf("%v\n", code)
}

func TestCustomCodeToID(t *testing.T) {
	assert := assert.New(t)
	inst := NewCustomInvCode("F6BGQ2DZX9C7P5IK3MJUAR4WYLTNVE8S", "H")
	code := "6FFFQ2DWDB66"
	id := inst.CodeToID(code)
	assert.Equal(id, int64(37232459628150785))
}
