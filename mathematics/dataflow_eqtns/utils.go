// https://stackoverflow.com/questions/40165842/function-that-accepts-all-numeric-types-int-float-and-adds-them
package main

func get64(x interface{}) interface{} {
	switch x := x.(type) {
	case uint8:
		return int64(x)
	case int8:
		return int64(x)
	case uint16:
		return int64(x)
	case int16:
		return int64(x)
	case uint32:
		return int64(x)
	case int32:
		return int64(x)
	case uint64:
		return int64(x)
	case int64:
		return int64(x)
	case int:
		return int64(x)
	case float32:
		return float64(x)
	case float64:
		return float64(x)
	}
	panic("invalid input")
}



func add(x, y interface{}) interface{} {
	switch x := get64(x).(type) {
	case int64:
		switch y := get64(y).(type) {
		case int64:
			return x + y
		case float64:
			return float64(x) + y
		}
	case float64:
		switch y := get64(y).(type) {
		case int64:
			return x + float64(y)
		case float64:
			return x + y
		}
	}

	panic("invlaid input")
}
