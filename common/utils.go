package common

import "reflect"

// AnyPtr 将任意类型的值转换为指针
// 如果值已经是指针，则直接返回
// 如果值是 nil，则返回 nil
// 否则创建一个新的指针指向该值
func AnyPtr(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		return v
	}

	ptr := reflect.New(rv.Type())
	ptr.Elem().Set(rv)
	return ptr.Interface()
}
