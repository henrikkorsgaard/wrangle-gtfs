package api


import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func init(){
	fmt.Println("Running rest tests")
}

func TestGetStops(t *testing.T) {
	Serve()
	assert.Equal(t, true, false)
}