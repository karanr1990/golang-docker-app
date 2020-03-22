package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"<h1>hello world</h1>")

}
func check(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"<h1>health check</h1>")

}
func main()  {

	http.HandleFunc("/",index)
	http.HandleFunc("/HealthCheck",check)

	fmt.Println("server starting.....")

	http.ListenAndServe(":8090",nil)
}
