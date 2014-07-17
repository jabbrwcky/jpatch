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

func TestNoModMerge(t *testing.T) {

	m1 := make(map[string]interface{})
	m1["foo"] = 1
	m1["bar"] = true

	m2 := make(map[string]interface{})

	result := merge(m1, m2)

	if len(result) != len(m1) {
		t.Fail()
	}

	if result["foo"] != m1["foo"] {
		t.Fail()
	}

	if result["bar"] != m1["bar"] {
		t.Fail()
	}
}

func TestEmptySourceObject(t *testing.T) {
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m2["foo"] = 42

	result := merge(m1, m2)

	if v, ok = result["foo"]; !ok {
		t.Fail()
	}
}
