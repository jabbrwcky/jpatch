package main

import (
	"testing"
)

func TestDisjunctMerge(t *testing.T) {
	m1 := make(map[string]interface{})
	m1["foo"] = 1

	m2 := make(map[string]interface{})
	m2["bar"] = true

	result := merge(m1, m2)

	if result["foo"] != 1 {
		t.Fail()
	}
	if result["bar"] != true {
		t.Fail()
	}
}
