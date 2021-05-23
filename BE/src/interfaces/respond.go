package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response struct {
	Status int               `json:"status"`
	Data   map[string]string `json:"data"`
}

func convertMapToJsonString(src interface{}) string {
	bytes, err := json.Marshal(src)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return ""
	}
	log.Println(string(bytes))
	return string(bytes)
}
