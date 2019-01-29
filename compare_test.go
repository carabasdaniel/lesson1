package main

import (
	"testing"

	v "github.com/carabasdaniel/lesson1/validator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMockCompare(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockValidator := v.NewMockValidator(ctrl)
	mockValidator.EXPECT().Compare(gomock.Any(), gomock.Any()).Return(true)

	assert.True(t, funnelCompare(mockValidator, "test", "tedjsahkjdhasst"))
}

func TestCompareSteps(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockValidator := v.NewMockValidator(ctrl)
	mockValidator.EXPECT().Compare("ake", "est").Return(false)
	mockValidator.EXPECT().Compare("fke", "est").Return(false)
	mockValidator.EXPECT().Compare("fae", "est").Return(false)
	mockValidator.EXPECT().Compare("fak", "est").Return(false)

	assert.False(t, funnelCompare(mockValidator, "fake", "est"))
}
