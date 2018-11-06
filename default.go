package file

import (
	"reflect"
	"strconv"
)

// SetDefaultValue 構造体を解析し、デフォルト値を設定します.
func SetDefaultValue(out interface{}) error {
	outType := getType(reflect.TypeOf(out))
	outRef := getValue(reflect.ValueOf(out))

	for i := 0; i < outType.NumField(); i++ {
		f := outType.Field(i)
		dv := f.Tag.Get("default")
		// デフォルトが設定されてなければ対象外
		if dv == "" {
			continue
		}

		v := outRef.FieldByName(f.Name)
		switch v.Kind() {
		case reflect.Bool:
			if val, err := strconv.ParseBool(dv); err != nil {
				return err
			} else {
				v.SetBool(val)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val, err := strconv.ParseInt(dv, 10, 64); err != nil {
				return err
			} else {
				v.SetInt(val)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if val, err := strconv.ParseUint(dv, 10, 64); err != nil {
				return err
			} else {
				v.SetUint(val)
			}
		case reflect.Float32, reflect.Float64:
			if val, err := strconv.ParseFloat(dv, 64); err != nil {
				return err
			} else {
				v.SetFloat(val)
			}

		default:
			v.Set(reflect.ValueOf(dv))
		}
	}
	return nil
}

func getType(target reflect.Type) reflect.Type {
	if target.Kind() != reflect.Ptr {
		return target
	}

	return getType(target.Elem())
}

func getValue(target reflect.Value) reflect.Value {
	if target.Kind() != reflect.Ptr {
		return target
	}

	return getValue(target.Elem())
}
