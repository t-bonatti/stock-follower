package composition

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGetComposition(t *testing.T) {
	items := GetComposition()

	assert := assert.New(t)
	assert.Equal(len(items), 13)
}
