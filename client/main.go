package main

import (
	"net/http"
	"bytes"
	"github.com/anvari1313/safe-proxy/common"
	"fmt"
	"strings"
	"io/ioutil"
	"io"
)


func main() {
	//a := common.Request{Body:"This", Method:"POST", Url:"Where"}
	//fmt.Println(common.Serialize(a))
	//var q common.Request
	//q = common.Deserialize(common.Serialize(a))
	//fmt.Println(q)
	addr := ":2580"
	http.HandleFunc("/", requestHandler)
	fmt.Printf("Safe Proxy Client Side is started at %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func standardizeRequest(request http.Request) common.Request {
	requestBodyBuffer := new(bytes.Buffer)
	requestBodyBuffer.ReadFrom(request.Body)
	headers := make(map[string] []string)
	for k, v := range request.Header  {
		headers[k] = v
	}

	return common.Request{
		Url: request.RequestURI,
		Method: request.Method,
		Body: requestBodyBuffer.String(),
		Header: headers,
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	common.LogNow()
	var requestObject = standardizeRequest(*r)
	response := sendToProxyServer(requestObject)
	fmt.Println(response)
	//send(common.SerializeRequest(requestObject))
	for k, v  := range response.Headers {
		w.Header().Set(k, v[0])
	}
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")


	w.WriteHeader(response.StatusCode)		// Send http status code
	io.WriteString(w, response.Body)
}

func sendToProxyServer(standardizeRequest common.Request) common.Response {
	serializedRequest := common.SerializeRequest(standardizeRequest)
	rawResponse, _ := http.Post("http://localhost:1580/st", "application/json; charset=utf-8", strings.NewReader(serializedRequest))
	readRawResponse, _ := ioutil.ReadAll(rawResponse.Body)
	return common.DeserializeResponse(string(readRawResponse))
}

//func send(r common.Request) {
func send(stream string) {
	//r := common.Request{Url:"/s", Method:"POST"}

	/*b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)*/
	fmt.Println(stream)
	res1, _ := http.Post("http://localhost:1580/st", "application/json; charset=utf-8", strings.NewReader(stream))
	//io.Copy(os.Stdout, res2.Body)
	stringBody, _ := ioutil.ReadAll(res1.Body)

	fmt.Println(string(stringBody))
}

//
//
//package main
//
//import (
//	"context"
//	"flag"
//	"fmt"
//	"log"
//	"net/http"
//	"os"
//	"os/signal"
//	"sync/atomic"
//	"time"
//)
//
//type key int
//
//const (
//	requestIDKey key = 0
//)
//
//var (
//	listenAddr string
//	healthy    int32
//)
//
//func main() {
//	flag.StringVar(&listenAddr, "listen-addr", ":5000", "server listen address")
//	flag.Parse()
//
//	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
//	logger.Println("Server is starting...")
//
//	router := http.NewServeMux()
//	router.Handle("/", index())
//	router.Handle("/healthz", healthz())
//
//	nextRequestID := func() string {
//		return fmt.Sprintf("%d", time.Now().UnixNano())
//	}
//
//	server := &http.Server{
//		Addr:         listenAddr,
//		Handler:      tracing(nextRequestID)(logging(logger)(router)),
//		ErrorLog:     logger,
//		ReadTimeout:  5 * time.Second,
//		WriteTimeout: 10 * time.Second,
//		IdleTimeout:  15 * time.Second,
//	}
//
//	done := make(chan bool)
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, os.Interrupt)
//
//	go func() {
//		<-quit
//		logger.Println("Server is shutting down...")
//		atomic.StoreInt32(&healthy, 0)
//
//		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//		defer cancel()
//
//		server.SetKeepAlivesEnabled(false)
//		if err := server.Shutdown(ctx); err != nil {
//			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
//		}
//		close(done)
//	}()
//
//	logger.Println("Server is ready to handle requests at", listenAddr)
//	atomic.StoreInt32(&healthy, 1)
//	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
//	}
//
//	<-done
//	logger.Println("Server stopped")
//}
//
//func index() http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.URL.Path != "/" {
//			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json; charset=utf-8")
//		w.Header().Set("X-Content-Type-Options", "nosniff")
//		w.WriteHeader(http.StatusOK)
//		fmt.Fprintln(w, "{\"KEY\":\"VALUE\"}")
//	})
//}
//
//func healthz() http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if atomic.LoadInt32(&healthy) == 1 {
//			w.WriteHeader(http.StatusNoContent)
//			return
//		}
//		w.WriteHeader(http.StatusServiceUnavailable)
//	})
//}
//
//func logging(logger *log.Logger) func(http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			defer func() {
//				requestID, ok := r.Context().Value(requestIDKey).(string)
//				if !ok {
//					requestID = "unknown"
//				}
//				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
//			}()
//			next.ServeHTTP(w, r)
//		})
//	}
//}
//
//func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			requestID := r.Header.Get("X-Request-Id")
//			if requestID == "" {
//				requestID = nextRequestID()
//			}
//			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
//			w.Header().Set("X-Request-Id", requestID)
//			next.ServeHTTP(w, r.WithContext(ctx))
//		})
//	}
//}