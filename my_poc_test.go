package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OK(t *testing.T) {

	s1 := []string{}
	s2 := []string(nil)

	ms := myStruct{}
	s3 := ms.nums

	assert.Equal(t, s1, s2) // 不相同
	assert.Equal(t, s1, s3) // 不相同
	assert.Equal(t, s2, s3) // 相同，因為都是指向nil
}

// go test my_poc_test.go my_poc.go
