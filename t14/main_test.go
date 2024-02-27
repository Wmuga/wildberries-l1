package main

import "testing"

type testCase struct {
	val      interface{}
	expected Type
}

func testRunner(cases []testCase, t *testing.T) {
	for i, testCase := range cases {
		caseType := getType(testCase.val)
		if testCase.expected != caseType {
			t.Errorf("Wrong case %d. Type of %#v: got %#v expected %#v\n", i, testCase.val, caseType, testCase.expected)
		}
	}
}

func TestNormalTypes(t *testing.T) {
	cases := []testCase{{"123", StringT}, {123, IntT}, {true, BoolT}, {make(chan int), ChanT}}
	testRunner(cases, t)
}

func TestPointerTypes(t *testing.T) {
	data1 := "123"
	data2 := 123
	data3 := true
	data4 := make(chan int)
	cases := []testCase{{&data1, StringPtrT}, {&data2, IntPtrT}, {&data3, BoolPtrT}, {&data4, ChanPtrT}}
	testRunner(cases, t)
}

func TestUnknownTypes(t *testing.T) {
	data1 := []int{1}
	data2 := struct{}{}
	cases := []testCase{{data1, Unknown}, {data2, Unknown}, {&data1, UnknownPtr}, {&data2, UnknownPtr}}
	testRunner(cases, t)
}
