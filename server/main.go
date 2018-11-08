package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/anvari1313/safe-proxy/common"
	"strings"
)

func main() {
	//var r = common.Request{
	//	Url: "this",
	//	Method: "POST",
	//}
	//var s = common.Serialize(r)
	//fmt.Println(s)
	//var r2 = common.Request{}
	//r2 = common.Deserialize(s)
	//fmt.Println(r2)
///////////////////////////////////////////////////////////////////////////////
	//common.Me()
	//var byteStream bytes.Buffer
	//var r = common.Request{Method:"POST", Url:"/r"}
	//common.Serialize(r, &byteStream)
	//fmt.Println(byteStream)
	//var q = common.Request{}
	//common.Deserialize(&q, &byteStream)
	//fmt.Println(q)
	//response, _ := http.Post("http://localhost:3000/?q=12", "text/plain", &byteStream)
	//fmt.Println("Result")
	//fmt.Println(response.Body)
	////////////////////////////////////////////
	var addr = ":1580"
	http.HandleFunc("/", requestHandler)
	fmt.Printf("Safe Proxy Server started at %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func sendToEndServer(standardRequest common.Request) common.Response {
	fmt.Println(standardRequest)
	client := &http.Client{}
	req, err := http.NewRequest(standardRequest.Method, standardRequest.Url, strings.NewReader(standardRequest.Body))
	for k, v := range standardRequest.Header {
		req.Header.Add(k, v[0])
	}
	if err != nil {
		fmt.Println("Error Occured")
	}
	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("Error Occured")
	}
	rawResponse, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println("Error Occured")
	}

	return common.Response{
		Headers: resp.Header,
		StatusCode: resp.StatusCode,
		Body: string(rawResponse),
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	parsedBody, _ := ioutil.ReadAll(r.Body)
	body := string(parsedBody)
	common.LogNow()
	fmt.Println(body)
	var parsedRequestBody common.Request
	json.Unmarshal(parsedBody, &parsedRequestBody)
	serverResponse := sendToEndServer(parsedRequestBody)

	jData, _ := json.Marshal(serverResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

	//response, err := http.Get("http://localhost:3000/?q=12")
	//fmt.Println(response.Body.Close())
	//if err != nil {
	//	fmt.Println("Error")
	//	fmt.Println(err)
	//} else {
	//	response := common.Response{
	//		Body: "SomeBody",
	//		StatusCode: 200,
	//		Headers: nil,
	//	}
	//
	//	//io.WriteString(w, "Successful response " + string(time.Now().String()))
	//}


	//io.WriteString(w, response.Body.Read())
}