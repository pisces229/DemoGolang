package app

import (
	"fmt"
	"reflect"
)

func demoReflect() {
	value := FirstStruct{
		"a1",
		SecondStruct{"a2"},
	}
	reflectType := reflect.TypeOf(value)
	reflectValue := reflect.ValueOf(value)
	fmt.Println(reflectType, reflectValue, reflectType.NumField())
	elem := reflect.ValueOf(&value).Elem()
	for i := 0; i < reflectType.NumField(); i++ {
		fmt.Println("...")
		{
			field := reflectType.Field(i)

			fmt.Println(field.Index, field.Name, field.Type.Kind())
			fmt.Println(elem.Field(i))
		}
		//fmt.Println("...")
		//{
		//	elem := reflect.ValueOf(&value).Elem()
		//	field := elem.Field(i)
		//	field.SetString("b1")
		//}
		fmt.Println("...")
	}

	fmt.Println(value)
}

type FirstStruct struct {
	StringValue string
	Second      SecondStruct
}
type SecondStruct struct {
	StringValue string
}
