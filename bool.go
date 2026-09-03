package easyconv

func BoolToString(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func BoolToStringOmitEmpty(v bool) string {
	if v {
		return "true"
	}
	return ""
}

func BoolToUInt8(v bool) uint8 {
	if v {
		return 1
	}
	return 0
}

func BoolToUInt16(v bool) uint16 {
	if v {
		return 1
	}
	return 0
}

func BoolToUInt32(v bool) uint32 {
	if v {
		return 1
	}
	return 0
}

func BoolToUInt64(v bool) uint64 {
	if v {
		return 1
	}
	return 0
}

func BoolToIntString(v bool) string {
	if v {
		return "1"
	}
	return "0"
}

func BoolToIntStringOmitEmpty(v bool) string {
	if v {
		return "1"
	}
	return ""
}
