package reflect_tools

import (
	"reflect"
)

/*
 * 使用 struct 给 out struct(model) 初始化
 * @Param out interface{} 要修改的结构体
 * @Param in interace{} 有数据的结构体
 */
func StructAssign(out interface{}, in interface{}) {
	outVal := reflect.ValueOf(out).Elem() //获取reflect.Type类型
	inVal := reflect.ValueOf(in).Elem()   //获取reflect.Type类型
	vTypeOfT := inVal.Type()
	for i := 0; i < inVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := outVal.FieldByName(name).IsValid(); ok {
			// log.Debug(name)
			// log.Debug(inVal.Field(i).Interface())
			outVal.FieldByName(name).Set(reflect.ValueOf(inVal.Field(i).Interface()))
		}
	}
}
