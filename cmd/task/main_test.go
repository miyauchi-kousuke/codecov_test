package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAchievingTargetExecute(t *testing.T) {
	assert.Equal(t, "hoge","hoge")
}

func TestAchievingTargetExecute2(t *testing.T) {

	main()

	assert.NotEqual(t, "hoge","fuga")
}