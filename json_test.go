package util

import (
  "fmt"
)

func ExampleNewJSON() {
  type TestStruct struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
  }
  test1 := TestStruct{
    Name: "Akvicor",
    Age:  17,
  }
  res := NewJSON(&test1, false)
  if res.Error() == nil {
    fmt.Println(res.String())
    fmt.Println(res.Bytes())
    fmt.Println(res.Map())
    fmt.Println(res.Map(0))
    fmt.Println(res.Map(10))
    fmt.Println(res.MapArray())
    fmt.Println(NewJSON(res.Map(), false))
  }
  test2 := [...]TestStruct{{Name: "Akvicor", Age: 17}, {Name: "MIU", Age: 17}}
  res = NewJSON(&test2, true)
  if res.Error() == nil {
    fmt.Println(res.Map(1))
    fmt.Println(res.MapArray())
  }
  
  // Output:
  // {"name":"Akvicor","age":17}
  // [123 34 110 97 109 101 34 58 34 65 107 118 105 99 111 114 34 44 34 97 103 101 34 58 49 55 125]
  // map[age:17 name:Akvicor]
  // map[age:17 name:Akvicor]
  // map[age:17 name:Akvicor]
  // [map[age:17 name:Akvicor]]
  // {"age":17,"name":"Akvicor"}
  // map[age:17 name:MIU]
  // [map[age:17 name:Akvicor] map[age:17 name:MIU]]
}

func ExampleNewJSONResult() {
  fmt.Println(NewJSONResult([]byte(`{"name":"Akvicor","age":17}`)).Map())
  fmt.Println(NewJSONResult([]byte(`[{"name":"Akvicor","age":17}]`)).Map())
  
  // Output:
  // map[age:17 name:Akvicor]
  // map[age:17 name:Akvicor]
}
