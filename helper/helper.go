// Package helper contains the helper functions.
//
// This package belongs to the Indicator project. Indicator is
// a Golang module that supplies a variety of technical
// indicators, strategies, and a backtesting framework
// for analysis.
//
// # License
//
//	Copyright (c) 2021-2024 Onur Cinar.
//	The source code is provided under GNU AGPLv3 License.
//	https://github.com/cinar/indicator/v2
//
// # Disclaimer
//
// The information provided on this project is strictly for
// informational purposes and is not to be construed as
// advice or solicitation to buy or sell any security.
package helper

// Integer refers to any integer type.
type Integer interface {
	int | int8 | int16 | int32 | int64
}

// Float refers to any float type.
type Float interface {
	float32 | float64
}

// Number refers to any numeric type.
type Number interface {
	Integer | Float
}
