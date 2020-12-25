package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//ini 文件解析器

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	DBName   string `ini:"dbname"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Pro      bool   `ini:"pro"`
}
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 参数校验 传进来的data必须是指针类型（因为要在函数中对其赋值）
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		// err = fmt.Errorf("") //格式化输出一个error类型
		err = errors.New("data param should be a pointer")
		return err
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")
		return err
	}
	//读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//字节转为字符串
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)

	var structName string
	//一行一行读数据
	for idx, line := range lineSlice {
		//去掉字符串首尾的空格
		line := strings.TrimSpace(line)
		// 注释 或者 空行 跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		// 以"["开头的
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line: %d  syntax error", idx+1)
				return err
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d  syntax error", idx+1)
				return err
			}
			// 根据字符串sectionName 去data里面根据反射找到对应的结构体
			//v :=reflect.ValueOf(data)
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的结构体%s\n", sectionName, structName)

				}
			}
		} else {
			// 	不以"["开头的 就是=分割的键值对
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") || strings.HasSuffix(line, "=") {
				err = fmt.Errorf("line:%d syntax error\n", idx+1)
				return err
			}
			//	1、以等号分割这一行，左边是key，右边是value
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//	2、根据structName 去data 里面把对应嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("便利data中的%v字段,data应该是一个结构体\n", structName)
				return err
			}
			var fieldName string
			//	3、遍历嵌套结构体的每一个字段，判断tag是否等于key，
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i)
				if filed.Tag.Get("ini") == key {
					// 找到了字段
					fieldName = filed.Name
					break
				}
			}
			//	4 如果key等于tag，给这个字段赋值
			fileObj := sValue.FieldByName(fieldName)
			fmt.Println(fieldName, fileObj.Kind())
			switch fileObj.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int16:
				var valueInt int64
				// ParseInt
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d syntax error\n", idx+1)
					return err
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d syntax error\n", idx+1)
					return err
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d syntax error\n", idx+1)
					return err
				}
				fileObj.SetFloat(valueFloat)

			}
		}

	}
	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Println("load ini file failed,err:", err)
	}
	//fmt.Printf("%v\n %v\n %v\n %v\n",cfg.Address,cfg.Port,cfg.Username,cfg.Password)
	fmt.Printf("%#v", cfg)
}
