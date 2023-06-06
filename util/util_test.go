package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueID(t *testing.T) {
	uniqueID := UniqueID()
	t.Logf("\n----> uniqueID: \"%s\"\n", uniqueID)
	assert.NotEmpty(t, uniqueID)
}

func TestCheckPhone(t *testing.T) {
	phone := "+6512345678"
	t.Logf("\n----> check phone number: \"%s\"\n", phone)
	assert.True(t, CheckPhone(phone))
}