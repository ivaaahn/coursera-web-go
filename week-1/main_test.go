package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
4
5`

var testNotSorted = `
5
4
3
2
1`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := unique(in, out)

	if err != nil {
		t.Errorf("test for OK was failed - error")
	}

	result := out.String()

	if result != testOkResult {
		t.Errorf("test for OK was failed - results doesn't match \n%#v\n%#v", result, testOkResult)
	}
}

func TestNotSortedError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testNotSorted))
	out := new(bytes.Buffer)
	err := unique(in, out)

	if err == nil {
		t.Errorf("test for not sorted input was failed")
	}
}
