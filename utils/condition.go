package utils

import (
	"reflect"
	"strings"
)

func IfAssigment[T any](cond bool, first, second T) T {
	if cond {
		return first
	}
	return second
}

// IsNil check
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func IsNilArgs(args ...interface{}) bool {
	allOfThemNil := true
	for _, arg := range args {
		if !IsNil(arg) {
			allOfThemNil = false
		}
	}
	return allOfThemNil
}

// IsEmail check
func IsEmail(i string) bool {
	return strings.Contains(i, "@")
}

func IsEmptyStr(value *string) bool {
	return value == nil || len(*value) == 0
}

func IsContains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
