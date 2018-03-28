package fec

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var FIELD_SEP = string([]byte{0x1c})
const RECORD_SEP = "\n"

type FECMarshaler interface {
	MarshalFECField(*[]string) error
}

type FECUnmarshaler interface {
	UnmarshalFECField(*[]string) error
}

type FECRecord interface {
	Name() string
	Record() []byte
	Id() string
	Schedule() string
}

func hasTag(f reflect.StructField, key string) bool {
	tagStr := f.Tag.Get("fec")
	if tagStr == "" {
		return false
	}
	tags := strings.Split(tagStr, ",")
	for _, tag := range tags {
		if tag == key {
			return true
		}
	}
	return false
}

func Marshal(s interface{}, fields *[]string) error {
	mer, isa := s.(FECMarshaler)
	if isa {
		return mer.MarshalFECField(fields)
	}
	rv := reflect.ValueOf(s)
	switch rv.Kind() {
	case reflect.String:
		sv, isa := rv.Interface().(string)
		if !isa && rv.Type().ConvertibleTo(reflect.TypeOf("")) {
			sv = rv.String()
		}
		*fields = append(*fields, sv)
	case reflect.Bool:
		v := rv.Interface().(bool)
		if v {
			*fields = append(*fields, "X")
		} else {
			*fields = append(*fields, "")
		}
	case reflect.Int:
		*fields = append(*fields, strconv.FormatInt(int64(rv.Interface().(int)), 10))
	case reflect.Int8:
		*fields = append(*fields, strconv.FormatInt(int64(rv.Interface().(int8)), 10))
	case reflect.Int16:
		*fields = append(*fields, strconv.FormatInt(int64(rv.Interface().(int16)), 10))
	case reflect.Int32:
		*fields = append(*fields, strconv.FormatInt(int64(rv.Interface().(int32)), 10))
	case reflect.Int64:
		*fields = append(*fields, strconv.FormatInt(rv.Interface().(int64), 10))
	case reflect.Uint:
		*fields = append(*fields, strconv.FormatUint(uint64(rv.Interface().(uint)), 10))
	case reflect.Uint8:
		*fields = append(*fields, strconv.FormatUint(uint64(rv.Interface().(uint8)), 10))
	case reflect.Uint16:
		*fields = append(*fields, strconv.FormatUint(uint64(rv.Interface().(uint16)), 10))
	case reflect.Uint32:
		*fields = append(*fields, strconv.FormatUint(uint64(rv.Interface().(uint32)), 10))
	case reflect.Uint64:
		*fields = append(*fields, strconv.FormatUint(rv.Interface().(uint64), 10))
	case reflect.Float32, reflect.Float64:
		*fields = append(*fields, fmt.Sprintf("%.2f", rv.Interface()))
	case reflect.Struct:
		var err error
		rt := rv.Type()
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			if hasTag(f, "skip") {
				continue
			}
			if strings.HasPrefix(f.Name, "XXX") {
				methName := strings.TrimPrefix(f.Name, "XXX")
				methType, ok := rt.MethodByName(methName)
				if ok {
					if methType.Type.NumIn() == 0 && methType.Type.NumOut() == 1 {
						method := rv.MethodByName(methName)
						rf := method.Call([]reflect.Value{})[0]
						err = Marshal(rf.Interface(), fields)
						if err != nil {
							fmt.Println("error marshaling field", f.Name, err)
							return err
						}
						continue
					}
				}
			}
			rf := rv.Field(i)
			err = Marshal(rf.Interface(), fields)
			if err != nil {
				fmt.Println("error marshaling field", f.Name, err)
				return err
			}
		}
	case reflect.Ptr:
		if !rv.IsNil() {
			return Marshal(rv.Elem().Interface(), fields)
		}
		if rv.Elem().Kind() == reflect.Struct {
			return Marshal(reflect.New(rv.Type().Elem()), fields)
		} else {
			*fields = append(*fields, "")
		}
	default:
		fmt.Println("unknown field type:", rv.Kind())
		return UnknownFieldTypeError
	}
	return nil
}

func Unmarshal(s interface{}, fields *[]string) error {
	if len(*fields) < 1 {
		return IncompleteFieldError
	}
	unmer, isa := s.(FECUnmarshaler)
	if (isa) {
		fmt.Printf("custom unmarshaler for %t\n", s)
		return unmer.UnmarshalFECField(fields)
	}
	rv := reflect.ValueOf(s)
	if rv.Kind() != reflect.Ptr {
		return errors.New("Can't unmarshal into non-pointer")
	}
	switch rv.Elem().Type().Kind() {
	case reflect.String:
		rv.Elem().SetString((*fields)[0])
		*fields = (*fields)[1:]
	case reflect.Bool:
		if strings.TrimSpace((*fields)[0]) == "" {
			rv.Elem().SetBool(false)
		} else {
			rv.Elem().SetBool(true)
		}
		*fields = (*fields)[1:]
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if (*fields)[0] != "" {
			i, err := strconv.ParseInt((*fields)[0], 10, 64)
			if err != nil {
				return err
			}
			rv.Elem().SetInt(i)
		}
		*fields = (*fields)[1:]
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if (*fields)[0] != "" {
			u, err := strconv.ParseUint((*fields)[0], 10, 64)
			if err != nil {
				return err
			}
			rv.Elem().SetUint(u)
		}
		*fields = (*fields)[1:]
	case reflect.Float32, reflect.Float64:
		if (*fields)[0] != "" {
			f, err := strconv.ParseFloat((*fields)[0], 64)
			if err != nil {
				return err
			}
			rv.Elem().SetFloat(f)
		}
		*fields = (*fields)[1:]
	case reflect.Struct:
		rvv := reflect.New(rv.Type().Elem())
		rt := rvv.Elem().Type()
		optional := false
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			if hasTag(f, "skip") {
				continue
			}
			if hasTag(f, "optional") {
				optional = true
			}
			if optional && len(*fields) == 0 {
				break
			}
			fmt.Println("unmarshaling field", f.Name)
			rf := rvv.Elem().Field(i)
			if rf.Kind() == reflect.Ptr {
				if rf.Type().Elem().Kind() != reflect.Struct && (*fields)[0] == "" {
					fmt.Println("skipping empty field", rf.Type().Elem().Kind())
					*fields = (*fields)[1:]
					continue
				}
				rfp := reflect.New(rf.Type().Elem())
				err := Unmarshal(rfp.Interface(), fields)
				if err != nil {
					//fmt.Printf("err in pointer field %s: %s", f.Name, err)
					return err
				}
				rf.Set(rfp)
			} else {
				err := Unmarshal(rf.Addr().Interface(), fields)
				if err != nil {
					//fmt.Printf("err in plain field %s: %s", f.Name, err)
					return err
				}
			}
		}
		rv.Elem().Set(rvv.Elem())
	default:
		fmt.Println("unknown field type:", rv.Type().Kind())
		return UnknownFieldTypeError
	}
	return nil
}


