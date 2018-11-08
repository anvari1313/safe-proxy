package common

import (
	"net/http"
	"fmt"
)

func PrintHeaders(r http.Request) {
	for k, v := range r.Header {
		fmt.Printf("%s : %s\n", k, v)
	}
}
