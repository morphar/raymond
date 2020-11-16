package raymond

import (
	"reflect"
	"strconv"
)

// Float returns a float representation of the provided type value
// Note that is only able to parse strings or convert actual numbers
// It defaults to 0.0 if the value cannot be properly represented as a float
func Float(value interface{}) float64 {
	return floatValue(reflect.ValueOf(value))
}

// floatValue returns the float64 representation of a reflect.Value
func floatValue(value reflect.Value) float64 {
	result := float64(0.0)

	switch value.Kind() {
	case reflect.String:
		result, _ = strconv.ParseFloat(value.String(), 64)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		result = float64(value.Int())
	case reflect.Float32, reflect.Float64:
		result = value.Float()
	}

	return result
}
