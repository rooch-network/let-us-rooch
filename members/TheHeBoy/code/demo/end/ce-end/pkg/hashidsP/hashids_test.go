package hashidsP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitHashIds(t *testing.T) {
	InitHashIds()
	a := assert.New(t)
	encode, err := HashID.Encode([]int{1})
	if err != nil {
		return
	}
	a.NotEqual(len(encode), 0, "HashID.Encode failed")
}
