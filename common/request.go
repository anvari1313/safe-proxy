package common

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Url string
	Method string
	Header map[string] []string
	Body string
}
/*
func Me() {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.

	// Encode (send) some values.
	err := enc.Encode(Request{"/r", "POST", "Column"})
	fmt.Println(network)
	if err != nil {
		log.Fatal("encode error:", err)
	}


	// Decode (receive) and print the values.
	var q Request
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%s: %s\n", q.Url, q.Method)
}

*/

//func Serialize(object Request, buffer *bytes.Buffer) {
//	enc := gob.NewEncoder(buffer)
//	err := enc.Encode(object)
//	if err != nil {
//		log.Fatal("encode error:", err)
//	}
//}
//
//func Deserialize(object *Request, buffer *bytes.Buffer) {
//	dec := gob.NewDecoder(buffer) // Will read from network.
//
//	err := dec.Decode(object)
//	if err != nil {
//		log.Fatal("decode error 1:", err)
//	}
//}

func SerializeRequest(r Request) (stream string) {
	res, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error Ocuured")

	}

	stream = string(res)
	return
}

func DeserializeRequest(stream string) (object Request) {
	var output Request
	err := json.Unmarshal([]byte(stream), &output)
	if err != nil {
		fmt.Println("Error Ocuured")

	}
	return output
}

//func ParseRequest(r *http.Request) Request {
//	return Request
//	}
//}