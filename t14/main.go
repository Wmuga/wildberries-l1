package main

import (
	"fmt"
	"reflect"
)

/*
 * Разработать программу, которая в рантайме
 * способна определить тип переменной: int, string, bool, channel
 * из переменной типа interface{}.
 */

type Type int

const (
	Unknown Type = 0
	IntT    Type = iota
	StringT
	BoolT
	ChanT
)

const (
	UnknownPtr Type = iota + 0x10
	IntPtrT
	StringPtrT
	BoolPtrT
	ChanPtrT
)

func getType(data interface{}) Type {
	dataVal := reflect.ValueOf(data)
	// "Разыменовывание" указателя. Чтобы знать что в нем
	if dataVal.Kind() == reflect.Pointer {
		return UnknownPtr | getType(dataVal.Elem().Interface())
	}
	// Смотрим какой тип данных
	switch dataVal.Kind() {
	case reflect.Int:
		return IntT
	case reflect.String:
		return StringT
	case reflect.Bool:
		return BoolT
	case reflect.Chan:
		return ChanT
	}
	return Unknown
}

func main() {
	v := 3456
	data := []any{"test", 123, &v, true, struct{}{}}
	for _, elem := range data {
		typeName := "Unknown"
		switch getType(elem) {
		case IntT:
			typeName = "int"
		case StringT:
			typeName = "string"
		case BoolT:
			typeName = "bool"
		case ChanT:
			typeName = "channel"
		case UnknownPtr:
			typeName = "Unknown pointer"
		case IntPtrT:
			typeName = "int pointer"
		case StringPtrT:
			typeName = "string pointer"
		case BoolPtrT:
			typeName = "bool pointer"
		case ChanPtrT:
			typeName = "channel pointer"
		}
		fmt.Printf("%#v - %s\n", elem, typeName)
	}
}
