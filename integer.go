package easyconv

import (
	"strconv"
	"unsafe"
)

func IntToString(v int) string {
	return strconv.Itoa(v)
}

func Int8ToString(v int8) string {
	return formatInt(int64(v))
}

func Int16ToString(v int16) string {
	return formatInt(int64(v))
}

func Int32ToString(v int32) string {
	return formatInt(int64(v))
}

func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func formatInt(v int64) string {
	return strconv.FormatInt(v, 10)
}

func UIntToString(v uint) string {
	return formatUInt(uint64(v))
}

func UInt8ToString(v uint8) string {
	return formatUInt(uint64(v))
}

func UInt16ToString(v uint16) string {
	return formatUInt(uint64(v))
}

func UInt32ToString(v uint32) string {
	return formatUInt(uint64(v))
}

func UInt64ToString(v uint64) string {
	return formatUInt(v)
}

func formatUInt(v uint64) string {
	return strconv.FormatUint(v, 10)
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToInt8(s string) (int8, error) {
	x, err := parseInt(s, 8)
	if err != nil {
		return 0, err
	}
	return int8(x), nil
}

func StringToInt16(s string) (int16, error) {
	x, err := parseInt(s, 16)
	if err != nil {
		return 0, err
	}
	return int16(x), nil
}

func StringToInt32(s string) (int32, error) {
	x, err := parseInt(s, 32)
	if err != nil {
		return 0, err
	}
	return int32(x), nil
}

func StringToInt64(s string) (int64, error) {
	x, err := parseInt(s, 64)
	if err != nil {
		return 0, err
	}
	return x, nil
}

func parseInt(s string, sz int) (int64, error) {
	return strconv.ParseInt(s, 10, sz)
}

func StringToUInt(s string) (uint, error) {
	x, err := parseUInt(s, 8*(int(unsafe.Sizeof(uint(0)))))
	return uint(x), err
}

func StringToUInt8(s string) (uint8, error) {
	x, err := parseUInt(s, 8)
	if err != nil {
		return 0, err
	}
	return uint8(x), nil
}

func StringToUInt16(s string) (uint16, error) {
	x, err := parseUInt(s, 16)
	if err != nil {
		return 0, err
	}
	return uint16(x), nil
}

func StringToUInt32(s string) (uint32, error) {
	x, err := parseUInt(s, 32)
	if err != nil {
		return 0, err
	}
	return uint32(x), nil
}

func StringToUInt64(s string) (uint64, error) {
	x, err := parseUInt(s, 64)
	if err != nil {
		return 0, err
	}
	return x, nil
}

func parseUInt(s string, sz int) (uint64, error) {
	return strconv.ParseUint(s, 10, sz)
}

func XStringToInt(s string) int {
	x, err := StringToInt(s)
	if err != nil {
		panic(err)
	}
	return x
}

func XStringToInt8(s string) int8 {
	x, err := StringToInt8(s)
	if err != nil {
		panic(err)
	}
	return x
}

func XStringToInt16(s string) int16 {
	x, err := StringToInt16(s)
	if err != nil {
		panic(err)
	}
	return x
}

func XStringToInt32(s string) int32 {
	x, err := StringToInt32(s)
	if err != nil {
		panic(err)
	}
	return x
}

func XStringToInt64(s string) int64 {
	x, err := StringToInt64(s)
	if err != nil {
		panic(err)
	}
	return x
}

func XStringToUInt(s string) uint {
	x, err := StringToUInt(s)
	if err != nil {
		panic(err)
	}
	return x
}

func XStringToUInt8(s string) uint8 {
	x, err := StringToUInt8(s)
	if err != nil {
		panic(err)
	}
	return x
}

func XStringToUInt16(s string) uint16 {
	x, err := StringToUInt16(s)
	if err != nil {
		panic(err)
	}
	return x
}

func XStringToUInt32(s string) uint32 {
	x, err := StringToUInt32(s)
	if err != nil {
		panic(err)
	}
	return x
}

func XStringToUInt64(s string) uint64 {
	x, err := StringToUInt64(s)
	if err != nil {
		panic(err)
	}
	return x
}
