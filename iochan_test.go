package iochan

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestDelimReader(t *testing.T) {
	buf := new(bytes.Buffer)
	buf.WriteString("foo\nbar\nbaz")

	ch := DelimReader(buf, '\n')

	results := make([]string, 0, 3)
	expected := []string{"foo\n", "bar\n", "baz"}
	for v := range ch {
		results = append(results, v)
	}

	if !reflect.DeepEqual(results, expected) {
		t.Fatalf("unexpected results: %#v", results)
	}
}

func TestLineReader(t *testing.T) {

	data := []string{"foo", "bar", "baz"}

	buf := new(bytes.Buffer)
	buf.WriteString(strings.Join(data, "\n") + "\n")

	ch := LineReader(buf)

	var result []string
	expected := data
	for v := range ch {
		result = append(result, v)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("unexpected results: %#v", result)
	}
}
