package jsonutil2

import (
	"fmt"
	"reflect"
	"strings"
)

func Marshal(v any) ([]byte, error) {
	vV := reflect.ValueOf(v)
	return []byte(getMarshalFunc(vV)(vV)), nil
}

func getMarshalFunc(vV reflect.Value) func(vV reflect.Value) string {
	switch vV.Kind() {
	case reflect.Bool:
		return boolMarshal
	case reflect.Int:
		return intMarshal
	case reflect.Int8:
		return int8Marshal
	case reflect.Int16:
		return int16Marshal
	case reflect.Int32:
		return int32Marshal
	case reflect.Int64:
		return int64Marshal
	case reflect.Uint:
		return uintMarshal
	case reflect.Uint8:
		return uint8Marshal
	case reflect.Uint16:
		return uint16Marshal
	case reflect.Uint32:
		return uint32Marshal
	case reflect.Uint64:
		return uint64Marshal
	case reflect.Uintptr:
		return uintptrMarshal
	case reflect.Float32, reflect.Float64:
		return intMarshal
	case reflect.String:
		return stringMarshal
	case reflect.Struct:
		return structMarshal
	case reflect.Interface:
		return interfaceMarshal
	case reflect.Map:
		return mapMarshal
	case reflect.Slice:
		return sliceMarshal
	case reflect.Array:
		return arrayMarshal
	case reflect.Pointer:
		return pointerMarshal
	default:
		panic(fmt.Sprintf("unsupported kind %v", vV.Kind()))
	}
}

func boolMarshal(vV reflect.Value) string {
	b := vV.Interface().(bool)
	if b {
		return "\"true\""
	} else {
		return "\"false\""
	}
}

func stringMarshal(vV reflect.Value) string {
	s := vV.Interface().(string)
	return "\"" + s + "\""
}

func int32Marshal(vV reflect.Value) string {
	s := vV.Interface().(int32)
	return fmt.Sprintf("%v", s)
}

func intMarshal(vV reflect.Value) string {
	s := vV.Interface().(int)
	return fmt.Sprintf("%v", s)
}

func int16Marshal(vV reflect.Value) string {
	s := vV.Interface().(int16)
	return fmt.Sprintf("%v", s)
}

func int8Marshal(vV reflect.Value) string {
	s := vV.Interface().(int8)
	return fmt.Sprintf("%v", s)
}

func int64Marshal(vV reflect.Value) string {
	s := vV.Interface().(int64)
	return fmt.Sprintf("%v", s)
}

func uint8Marshal(vV reflect.Value) string {
	s := vV.Interface().(uint8)
	return fmt.Sprintf("%v", s)
}

func uint16Marshal(vV reflect.Value) string {
	s := vV.Interface().(uint16)
	return fmt.Sprintf("%v", s)
}

func uint32Marshal(vV reflect.Value) string {
	s := vV.Interface().(uint32)
	return fmt.Sprintf("%v", s)
}

func uintMarshal(vV reflect.Value) string {
	s := vV.Interface().(uint)
	return fmt.Sprintf("%v", s)
}

func uint64Marshal(vV reflect.Value) string {
	s := vV.Interface().(uint64)
	return fmt.Sprintf("%v", s)
}

func interfaceMarshal(vV reflect.Value) string {
	iV := vV.Elem()
	return getMarshalFunc(iV)(iV)
}

func uintptrMarshal(vV reflect.Value) string {
	s := vV.Interface().(uintptr)
	return fmt.Sprintf("%v", s)
}

func structMarshal(vV reflect.Value) string {
	s := "{"
	n := vV.NumField()
	for i := 0; i < n; i++ {
		tag := vV.Type().Field(i).Tag.Get("json")
		if tag == "-" {
			continue
		}
		if tag == "" {
			continue
		}
		vVnext := vV.Field(i)
		s += "\"" + tag + "\":" + getMarshalFunc(vVnext)(vVnext)
		if i < n-1 {
			if nexTag := vV.Type().Field(i).Tag.Get("json"); nexTag != "" && nexTag != "-" {
				s += ","
			}
		}
	}
	s += "}"
	return s
}

func mapMarshal(vV reflect.Value) string {
	if vV.IsNil() {
		return "<nil>"
	}
	s := "{"
	keys := vV.MapKeys()
	for i := 0; i < len(keys); i++ {
		vVnext := vV.MapIndex(keys[i])
		key := getMarshalFunc(keys[i])(keys[i])
		if !strings.HasPrefix(key, "\"") || !strings.HasSuffix(key, "\"") {
			key = "\"" + key + "\""
		}
		s += key + ":" + getMarshalFunc(vVnext)(vVnext)
		if i < len(keys)-1 {
			s += ","
		}
	}
	s += "}"
	return s
}

func sliceMarshal(vV reflect.Value) string {
	s := "["
	for i := 0; i < vV.Len(); i++ {
		vVIdx := vV.Index(i)
		s += getMarshalFunc(vVIdx)(vVIdx)
		if i < vV.Len()-1 {
			s += ","
		}
	}
	s += "]"
	return s
}

func arrayMarshal(vV reflect.Value) string {
	s := "["
	for i := 0; i < vV.Len(); i++ {
		vVIdx := vV.Index(i)
		s += getMarshalFunc(vVIdx)(vVIdx)
		if i < vV.Len()-1 {
			s += ","
		}
	}
	s += "]"
	return s
}

func pointerMarshal(vV reflect.Value) string {
	if vV.IsNil() {
		return "<nil>"
	}
	s := vV.Pointer()
	return fmt.Sprintf("0x%x", s)
}
