package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWord1(t *testing.T) {
	result := isChangeOnce("telkom", "telecom")
	assert.Equal(t, false, result)
}

func TestWord2(t *testing.T) {
	result := isChangeOnce("telkom", "tlkom")
	assert.Equal(t, true, result)
}
