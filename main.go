package parseForm

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Input struct {
	Id   int `form:"code"`
	Name string
	Cg   Change
}

type Change struct {
	Cname string
}

func parseForm(form url.Values, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	// 判断容器属性，如果不是指针类型或者指针下面不是一个结构体，返回错误
	if objT.Elem().Kind() != reflect.Struct || objT.Kind() != reflect.Ptr {
		return errors.New("obj not a ptrStruct")
	}
	objT, objV = objT.Elem(), objV.Elem()

	return formToStruct(form, objT, objV)
}

func formToStruct(form url.Values, objT reflect.Type, objV reflect.Value) error {
	for i := 0; i < objT.NumField(); i++ {
		fieldT := objT.Field(i)
		fieldV := objV.Field(i)
		// 如果字段是结构体
		if fieldV.Kind() == reflect.Struct {
			formToStruct(form, fieldT.Type, fieldV)
		}
		// 获取key值
		var key string

		tags := strings.Split(fieldT.Tag.Get("form"), ",")
		if len(tags) == 0 || tags[0] == "" {
			key = fieldT.Name
		} else if tags[0] == "-" {
			continue
		} else {
			key = tags[0]
		}
		// 根据key从form取值,如果form没有值，则跳过，进行下一个字段
		value := form.Get(key)
		if value == "" {
			continue
		}
		// 根据字段类型，使用不同的set方法
		switch fieldV.Kind() {
		case reflect.Int:
			x, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("parse int err:", err.Error())
				return err
			}
			fieldV.SetInt(int64(x))
		case reflect.String:
			fieldV.SetString(value)
		}
	}
	return nil
}
