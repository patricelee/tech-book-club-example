package p39

import (
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestFooDatabaseByFunc(t *testing.T) {
	want := 1
	patch := monkey.Patch(getUserAgeFunc, func() int {
		return 1
	})
	defer patch.Restore()
	actual := fooDatabaseCaseByFunc()
	assert.Equal(t, want, actual)
}
