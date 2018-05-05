package envu

import (
	"os"
	"reflect"
	"strconv"
)

// Getenv environment or default
func Getenv(key string, def interface{}, to interface{}) error {

	val := reflect.ValueOf(to)
	if val.Type().Kind() != reflect.Ptr {
		panic("should be a pointer")
	}

	s := os.Getenv(key)

	switch v := to.(type) {
	case *int:
		if s == "" {
			*v = def.(int)
			return nil
		}
		envVal, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		*v = envVal
	case *string:
		if s == "" {
			*v = def.(string)
			return nil
		}
		*v = s
	default:
		panic("unknown type")
	}
	return nil
}
