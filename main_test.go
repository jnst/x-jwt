package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tokenString := generate()
	c1 := parse(tokenString)
	fmt.Println(c1.IssuedAt)
	//fmt.Println(c1.IssuedAt.Location().String())
	c2 := parse(tokenString)
	//fmt.Println(c2.IssuedAt.Location().String())
	assert.Equal(t, c1.IssuedAt, c2.IssuedAt)
}
