package util

import (
	"reflect"
)

type basisKind int

const (
	invalidKind basisKind = iota
	nullKind
	intKind
	uintKind
	floatKind
	complexKind
	stringKind
	mapKind
	structKind
	iterableKind
	boolKind
	pointerKind
)

var (
	anyType = reflect.ValueOf((*any)(nil)).Type().Elem()
)

func basicKind(v reflect.Value) basisKind {
	if !v.IsValid() || IsNil(v.Interface()) {
		return nullKind
	}

	switch v.Kind() {
	case reflect.Bool:
		return boolKind
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intKind
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintKind
	case reflect.Float32, reflect.Float64:
		return floatKind
	case reflect.Complex64, reflect.Complex128:
		return complexKind
	case reflect.String:
		return stringKind
	case reflect.Map:
		return mapKind
	case reflect.Struct:
		return structKind
	case reflect.Slice, reflect.Array:
		return iterableKind
	case reflect.Pointer:
		return pointerKind
	}
	return invalidKind
}

func rawValue(x reflect.Value) reflect.Value {
	if !x.IsValid() {
		return x
	}

	return reflect.ValueOf(x.Interface())
}
