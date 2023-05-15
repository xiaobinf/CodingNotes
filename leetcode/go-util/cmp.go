package go_util

import (
	"reflect"
)

type CMPRES int8

const (
	INCMP CMPRES = iota - 2
	LT
	EQ
	GT
)

func Compare(lhs, rhs interface{}) CMPRES {
	if !isCompareable(lhs, rhs) {
		return INCMP
	}
	lhsVal := reflect.ValueOf(lhs)
	rhsVal := reflect.ValueOf(rhs)

	switch lhsVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch {
		case lhsVal.Int() < rhsVal.Int():
			return LT
		case lhsVal.Int() == rhsVal.Int():
			return EQ
		case lhsVal.Int() > rhsVal.Int():
			return GT
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch {
		case lhsVal.Uint() < rhsVal.Uint():
			return LT
		case lhsVal.Uint() == rhsVal.Uint():
			return EQ
		case lhsVal.Uint() > rhsVal.Uint():
			return GT
		}
	case reflect.Float32, reflect.Float64:
		switch {
		case lhsVal.Float() < rhsVal.Float():
			return LT
		case lhsVal.Float() == rhsVal.Float():
			return EQ
		case lhsVal.Float() > rhsVal.Float():
			return GT
		}
	case reflect.String:
		switch {
		case lhsVal.String() < rhsVal.String():
			return LT
		case lhsVal.String() == rhsVal.String():
			return EQ
		case lhsVal.String() > rhsVal.String():
			return GT
		}
	}
	return INCMP
}

func isCompareable(lhs, rhs interface{}) bool {
	return reflect.ValueOf(lhs).Kind() == reflect.ValueOf(rhs).Kind()
}

func CompareEQ(lhs, rhs interface{}) bool {
	return Compare(lhs, rhs) == EQ
}

func CompareLT(lhs, rhs interface{}) bool {
	return Compare(lhs, rhs) == LT
}

func CompareGT(lhs, rhs interface{}) bool {
	return Compare(lhs, rhs) == GT
}

func CompareGE(lhs, rhs interface{}) bool {
	res := Compare(lhs, rhs)
	return res == EQ || res == GT
}

func CompareLE(lhs, rhs interface{}) bool {
	res := Compare(lhs, rhs)
	return res == EQ || res == LT
}

/*
如果switch后面什么都不带，则每个case分支的表达式结果只要是boolena类型就可以
如果switch后面带了某个变量，比如int类型的，则每个case分支的表达式结果必须跟这个变量的类型保持一致，这样switch case才能对条件做校验比较
总结起来就是：落脚到每个case 语句都是boolean类型，true或者false
*/
