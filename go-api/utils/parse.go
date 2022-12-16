package utils

import "encoding/json"

func ParseObject[IN, OUT any](obj *IN) OUT {
	jsonObj, _ := json.Marshal(&obj)

	var result OUT
	json.Unmarshal(jsonObj, &result)
	return result
}
