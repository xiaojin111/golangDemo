package Test1

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestSomething(t *testing.T) {
	//assert.Equal(t, 123, 1234, "相等%s", "hello")
	assert.NotNil(t, nil)
	str := "222"
	log.Println()
}
