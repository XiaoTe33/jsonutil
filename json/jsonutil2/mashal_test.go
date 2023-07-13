package jsonutil2_test

import (
	"fmt"
	"jsonutil/json/jsonutil2"
	"reflect"
	"testing"
)

type Pointer struct {
	p int
}
type S struct {
	Id   int `json:"id"`
	Info struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	} `json:"info"`
	Score map[string]int `json:"-"`
	Habit map[string]any `json:"habit"`
	Class []interface{}  `json:"class"`
	P     uintptr        `json:""`
	Ptr   *Pointer       `json:"ptr"`
	Nil   []string       `json:"nil"`
}

func TestMarshal(t *testing.T) {
	a := S{
		Id: 2022,
		Info: struct {
			Name string `json:"name"`
			Age  int64  `json:"age"`
		}{
			Name: "xiaoming",
			Age:  18,
		},
		Score: map[string]int{
			"math":    63,
			"english": 75,
		},
		Habit: map[string]any{
			"exercise": map[string]bool{
				"football":   true,
				"basketball": false,
			},
		},
		Class: []interface{}{
			1,
			"2",
			true,
			[]any{
				1,
				"2",
				true,
				map[string]any{
					"k1": "v1",
					"k2": "v2",
				},
			},
		},
		P:   0x3,
		Ptr: nil,
		Nil: nil,
	}
	marshal, err := jsonutil2.Marshal(a)
	fmt.Println(string(marshal), err)
}

func TestNumMarshal(t *testing.T) {
	var n uint64 = 18446744073709551615
	fmt.Println(uint64Marshal(&n))
}

func TestA(t *testing.T) {
	a := []any{"1", 2, true}
	aT := reflect.TypeOf(a)
	aV := reflect.ValueOf(a)
	fmt.Printf("reflect.TypeOf(a):%v\n", aT)
	fmt.Printf("reflect.ValueOf(a):%v\n", aV)
	aPtrT := reflect.TypeOf(&a)
	aPtrV := reflect.ValueOf(&a)
	fmt.Printf("reflect.TypeOf(&a):%v\n", aPtrT)
	fmt.Printf("reflect.ValueOf(&a):%v\n", aPtrV)

	fmt.Printf("aV.Interface():%v\n", aV.Interface())
	fmt.Printf("aPtrV.Interface()%v\n", aPtrV.Interface())

	fmt.Println(aV.Index(2))
	//fmt.Println(aT.Key())
	//fmt.Println(aT.Size())
	//fmt.Println(aT.Kind())
	//fmt.Println(aT.Align())
	//fmt.Println(aT.NumField())
	//var iSlice []int
	//for i := 0; i < aT.NumField(); i++ {
	//	iSlice = append(iSlice, i)
	//}
	//fmt.Println()
	//fmt.Println(aT.FieldByIndex(iSlice))
	//for i := 0; i < aT.NumField(); i++ {
	//	fmt.Println(reflect.ValueOf(aV.Field(i).Interface()).Kind())
	//}
}

func boolMarshal(v any) string {
	vV := reflect.ValueOf(v)
	b := vV.Interface().(*bool)
	if *b {
		return "\"true\""
	} else {
		return "\"false\""
	}
}

func uint64Marshal(v any) string {
	vV := reflect.ValueOf(v)
	s := vV.Interface().(*uint64)
	return fmt.Sprintf("%v", *s)
}
