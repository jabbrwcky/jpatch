/*
 *  Copyright 2014 Jens Hausherr
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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

	if _, ok := result["foo"]; !ok {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})

	m1["foo"] = "foo"
	m1["bar"] = 42
	m1["baz"] = true

	m2["foo"] = nil
	m2["baz"] = nil

	result := merge(m1, m2)

	if len(result) != 1 {
		t.Fail()
	}

	if result["bar"] != 42 {
		t.Fail()
	}
}

func TestMergeNested(t *testing.T) {
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m1a := make(map[string]interface{})
	m2a := make(map[string]interface{})

	m1["foo"] = 29
	m1a["baz"] = true
	m1["bar"] = m1a

	m2a["baz"] = false
	m2a["boo"] = "here"
	m2["bar"] = m2a

	result := merge(m1, m2)

	var nested map[string]interface{}
	var ok bool

	if nested, ok = result["bar"].(map[string]interface{}); !ok {
		t.Fail()
	}

	if nested["baz"] != false {
		t.Fail()
	}

	if nested["boo"] == "here" {
		t.Fail()
	}
}

func TestReplaceArray(t *testing.T) {
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})

	a1 := []int{1, 2, 3, 4}
	a2 := []int{1}

	m1["foo"] = a1
	m2["foo"] = a2

	result := merge(m1, m2)

	if len(result["foo"].([]int)) != len(a2) {
		t.Fail()
	}
}
