// authors: wangoo
// created: 2018-07-11

package main

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
)

var (
	content = `http://golang.org/a/b
http://golang.org/a/c
http://golang.org/a/d
http://golang.org/a/e
`
)

func TestPatternBatchProcess(t *testing.T) {
	regex := "http://golang.org(?P<uri>[/\\w\\.]+)"
	PatternBatchProcess([]byte(content), regex, []byte("$uri"), func(s string) {
		fmt.Println(s)
		assert.True(t, strings.HasPrefix(s, "/a/"))
	})
}

func TestUriBatchProcess(t *testing.T) {
	prefix := "http://golang.org"
	UriBatchProcess(prefix, []byte(content), func(s string) {
		fmt.Println(s)
		assert.True(t, strings.HasPrefix(s, "a/"))
	})
}
