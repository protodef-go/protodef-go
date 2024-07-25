package datatypes

func GetNativeType(name string) *Type {
	switch name {
	case "i8":
		return Int8
	case "u8":
		return UInt8
	case "i16":
		return Int16
	case "u16":
		return UInt16
	case "i32":
		return Int32
	case "u32":
		return UInt32
	case "f32":
		return fInt32
	case "f64":
		return fInt64
	case "i64":
		return Int64
	case "u64":
		return UInt64
	}

	return nil
}
