package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

/**
Base Constraints for all Data Types
*/

func required(field field, param string) error {
	switch field.typ.Kind() {
	case reflect.String:
		c, err := convertBool(param)
		if err != nil {
			return err
		}
		if c == true {
			in, _ := field.value.Interface().(string)
			if in == "" {
				return ErrRequired
			}
		}
	case reflect.Bool:
	case reflect.Int:
	case reflect.Float32:
	case reflect.Uint:
	}
	return nil
}

func nillable(field field, param string) error {
	return nil
}

func def(field field, param string) error {
	return nil
}

/**
Numerical Type Constraints
*/

func min(field field, param string) error {
	return checkMin(field.value, field.typ, param, false)
}

func max(field field, param string) error {
	return checkMax(field.value, field.typ, param, false)
}

func exclusiveMin(field field, param string) error {
	return checkMin(field.value, field.typ, param, true)
}

func exclusiveMax(field field, param string) error {
	return checkMax(field.value, field.typ, param, true)
}

func multipleOf(field field, param string) error {
	// TODO : works only for int as of now
	valid := true
	in, _ := field.value.Interface().(int)
	c, err := convertInt(param, 0)
	cInt := int(c)
	if err != nil {
		return err
	}
	valid = in%cInt == 0
	if !valid {
		return ErrMultipleOf
	}
	return nil
}

/**
String Type Constraints
*/

func minLength(field field, param string) error {
	valid := true
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(field.value))
	valid = lv > lc
	if !valid {
		return ErrMinLength
	}
	return nil
}

func maxLength(field field, param string) error {
	valid := true
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(field.value))
	valid = lv < lc
	if !valid {
		return ErrMaxLength
	}
	return nil
}

func pattern(field field, param string) error {
	in, _ := field.value.Interface().(string)
	if field.value.Kind() != reflect.String {
		return ErrNotSupported
	}
	re, err := regexp.Compile(param)
	if err != nil {
		return ErrBadConstraint
	}
	if !re.MatchString(in) {
		return ErrPattern
	}
	return nil
}

func enum(field field, param string) error {
	flag := false
	switch field.value.Kind() {
	case reflect.Int:
		input := field.value.Interface().(int)
		flag = checkIfEnumExists(strconv.Itoa(input), param, ",")
	case reflect.String:
		input := field.value.String()
		flag = checkIfEnumExists(input, param, ",")
	}

	if flag == false {
		return ErrEnums
	}
	return nil
}
