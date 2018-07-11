// authors: wangoo
// created: 2018-07-11

package main

import (
	"regexp"
	"strings"
)

type PatternBatchProcessor func(string)

func PatternBatchProcess(content []byte, regex string, regexTemplate []byte, processor PatternBatchProcessor) {
	pattern := regexp.MustCompile(regex)
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		result := []byte{}
		result = pattern.Expand(result, regexTemplate, content, submatches)
		processor(string(result))
	}
}

func UriBatchProcess(batchUrlPrefix string, bytes []byte, processor PatternBatchProcessor) {
	regex := batchUrlPrefix
	if !strings.HasSuffix(regex, "/") {
		regex = regex + "/"
	}
	regex = regex + "(?P<uri>[/\\w\\.\\-\\_]+)"
	PatternBatchProcess(bytes, regex, []byte("$uri"), processor)
}
