package fake

import (
	"encoding/json"
	"fmt"
)

func ReturnJson() ([]byte,error) {
	data := map[string]interface{}{
		"author":       "name",
		"uuid":         1,
		"Youtube_link": "name",
	}
	jsonData, err := json.Marshal(data)
	if err!=nil{
		fmt.Print("could not marshal json",err)
	}
	return jsonData,nil
}