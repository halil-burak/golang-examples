package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	resp1 := response1{
		Page:   1,
		Fruits: []string{"pear", "apple", "peach"}}

	resp1Json, _ := json.Marshal(resp1)
	fmt.Println(string(resp1Json))

	resp2 := response2{
		Page:   2,
		Fruits: []string{"watermelon", "pear", "banana"}}
	resp2Json, _ := json.Marshal(resp2)
	fmt.Println(string(resp2Json))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}
