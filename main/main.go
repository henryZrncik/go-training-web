package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)
func rootHandler(w http.ResponseWriter, r *http.Request) {
fmt.Fprint(w, "<h1>  hello world</h1>")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>  hello </h1>")
	// parsed from url
	var msg2 = r.Form.Get("message")
	log.Printf("message is: %s", msg2)
}

func postHandler(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w,"bad format of of form, please provide 'message' data key")
	}
	// TODO update parse
	var msg string = r.PostForm.Get("message")
	fmt.Fprintf(w, "you posted msg: '%s'", msg )
}

func delayedCodedHandler(w http.ResponseWriter, r *http.Request) {

	delay := r.URL.Query().Get("delay")
	log.Printf("original delay: '%s'",delay)
	delayInt, _ := strconv.Atoi(delay)
	log.Println(delayInt)

	errCode := r.URL.Query().Get("code")
	log.Printf("original errCode: '%s'",errCode)
	errCodeInt, _ := strconv.Atoi(errCode)
	log.Println(errCodeInt)

	time.Sleep(time.Duration(delayInt) * time.Second)
	w.WriteHeader(int(errCodeInt))
}

func onlyPostHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		// must be before any writing
		w.Header().Set("IAllowMethods", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main(){

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET", "POST")
	r.HandleFunc("/home", homeHandler)
	r.HandleFunc("/post", postHandler)
	r.HandleFunc("/onlyPost", onlyPostHandler)
	r.HandleFunc("/delayedCode", delayedCodedHandler)
	http.ListenAndServe(":3000", r)
	//
	//r.HandleFunc("/data", datasController.Handler)
}

