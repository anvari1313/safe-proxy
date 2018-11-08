package common

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	StatusCode int
	Headers map[string] []string
	Body string
}


func SerializeResponse(r Response) (stream string) {
	res, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error Ocuured")

	}

	stream = string(res)
	return
}

func DeserializeResponse(stream string) (object Response) {
	var output Response
	err := json.Unmarshal([]byte(stream), &output)
	if err != nil {
		fmt.Println("Error Ocuured")

	}
	return output
}