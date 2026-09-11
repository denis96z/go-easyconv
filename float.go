package easyconv

import (
	"strconv"
)

func Float32ToString(x float32) string {
	return strconv.FormatFloat(float64(x), 'f', -1, 32)
}

func Float32ToStringOmitEmpty(v float32) string {
	if v == 0 {
		return ""
	}
	return Float32ToString(v)
}

func Float64ToString(x float64) string {
	return strconv.FormatFloat(x, 'f', -1, 64)
}

func Float64ToStringOmitEmpty(v float64) string {
	if v == 0 {
		return ""
	}
	return Float64ToString(v)
}

func StringToFloat32(s string) (float32, error) {
	x, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}
	return float32(x), nil
}

func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func XStringToFloat32(s string) float32 {
	x, err := StringToFloat32(s)
	if err != nil {
		panic(err)
	}
	return x
}

func XStringToFloat64(s string) float64 {
	x, err := StringToFloat64(s)
	if err != nil {
		panic(err)
	}
	return x
}
