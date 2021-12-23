package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/goinggo/mapstructure"
)

type People struct {
	Name string `json:"name_title"`
	Age  int    `json:"age_size"`
}

func JsonToStructDemo() {
	jsonStr := `
        {
                "name_title": "jqw"
                "age_size":12
        }
        `
	var people People
	json.Unmarshal([]byte(jsonStr), &people)
	fmt.Println(people)
}

func StructToJsonDemo() {
	p := People{
		Name: "jqw",
		Age:  18,
	}

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}

func JsonToMapDemo() {
	jsonStr := `
	{
			"name": "jqw",
			"age": 18
	}
	`
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	fmt.Println(mapResult)
}

func MapToJsonDemo1() {
	mapInstances := []map[string]interface{}{}
	instance_1 := map[string]interface{}{"name": "John", "age": 10}
	instance_2 := map[string]interface{}{"name": "Alex", "age": 12}
	mapInstances = append(mapInstances, instance_1, instance_2)

	jsonStr, err := json.Marshal(mapInstances)

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
}
func MapToJsonDemo2() {
	b, _ := json.Marshal(map[string]int{"test": 1, "try": 2})
	fmt.Println(string(b))
}

func MapToStructDemo() {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "jqw"
	mapInstance["Age"] = 18

	var people People
	err := mapstructure.Decode(mapInstance, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}

func StructToMapDemo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// func TestStructToMap() {
// 	student := Student{10, "jqw", 18}
// 	data := StructToMapDemo(student)
// 	fmt.Println(data)
// }

func main() {
	// JsonToStructDemo()
	// StructToJsonDemo()
	JsonToMapDemo()
	// MapToJsonDemo1()
	// MapToJsonDemo2()
	// MapToStructDemo()
	// StructToMapDemo()
	// TestStructToMap()
}
