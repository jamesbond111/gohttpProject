package main

import (
	"fmt"
	"net/http"
)


func sayHello(w http.ResponseWriter,r *http.Request)  {
	n, err := fmt.Fprintln(w, "<h1>hello james</h1>")
	if err != nil {
		return
	}
	fmt.Print(n)

}

/**
	使用原生的net/http包开启http server
 */
func main() {
	http.HandleFunc("/hello",sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		return
	}
}
