package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJson2Bson(t *testing.T) {
	in := []byte(`{
	"beuha": "aussi",
	"age": 42

}`)
	fmt.Println(in)
	out, err := Json2Bson(in)
	assert.NoError(t, err)
	fmt.Println(out)
	ioutil.WriteFile("debug.bson", out, 0640)
	assert.Fail(t, "oups")
}
