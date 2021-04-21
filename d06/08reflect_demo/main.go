package main

import (
	"fmt"
	"reflect"
)

type cat struct {
	name string
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v,type kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64,value is %f\n", v.Float())
	case reflect.Struct:
		fmt.Printf("type is struct,value is %v\n", v)
	}
}
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	//反射中使用Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func reflectCheck() {
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	fmt.Println("nil IsNil:", reflect.ValueOf(nil).IsValid())

	b := struct{}{}
	fmt.Println(reflect.ValueOf(b).FieldByName("abd").IsValid())  //结构体字段查找字段的信息
	fmt.Println(reflect.ValueOf(b).MethodByName("abd").IsValid()) //结构体字段查找方法的信息

	c := map[string]int{}
	fmt.Println(reflect.ValueOf(c).MapIndex(reflect.ValueOf("伺机待发")).IsValid())
}

func f1() {
	var a float32 = 3.14
	reflectType(a)
	reflectValue(a)

	var b int64 = 100
	reflectType(b)
	reflectValue(b)

	c := cat{
		name: "mimi",
	}
	reflectType(c)
	reflectValue(c)

	//设置值
	reflectSetValue1(&b)
	fmt.Println(b)
	reflectSetValue2(&b)
	fmt.Println(b)
}

func main() {
	f1()
	reflectCheck()
}
