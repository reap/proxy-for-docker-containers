package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handleProxyRequest(w http.ResponseWriter, r *http.Request) {
	service, url := SplitServiceURL(r.URL.Path)
	var buffer bytes.Buffer

	buffer.WriteString("http://")
	buffer.WriteString(service)
	buffer.WriteString("/")
	buffer.WriteString(url)

	fmt.Println("Going to call ", buffer.String())

	response, err := http.Get(buffer.String())
	if nil != err {
		log.Fatal(err)
	}
	defer response.Body.Close()

	resposeBuffer := new(bytes.Buffer)
	resposeBuffer.ReadFrom(response.Body)
	w.Write(resposeBuffer.Bytes())
}

func SplitServiceURL(s string) (string, string) {
	parts := strings.SplitN(strings.TrimPrefix(s, "/"), "/", 2)
	return parts[0], parts[1]

}

func main() {
	log.Println("Starting to listen for requests")
	http.HandleFunc("/", handleProxyRequest)
	http.ListenAndServe(":8080", nil)
}
