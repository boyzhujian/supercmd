package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "supercmd turns your sever into one giant api server ")
}

//filename   ,target  two querystring
func curluploadfileHandler(w http.ResponseWriter, r *http.Request) {
	path, err := exec.LookPath("curl")
	if err != nil {
		log.Fatal("curl is not installed on you system ,can not upload")
	}
	fmt.Printf("curl is available at %s\n", path)

	filename := r.URL.Query().Get("filename")
	target := r.URL.Query().Get("target")
	s := "file=@" + filename
	fmt.Println(s)
	fmt.Println(target)
	cmd := exec.Command(path, "-F", s, target)

	err = cmd.Run()
	fmt.Fprint(w, err)
	o, errr := cmd.Output()
	fmt.Fprint(w, o)
	fmt.Fprint(w, errr)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", mainHandler)
	router.HandleFunc("/curl/upload", curluploadfileHandler)
	http.ListenAndServe(":8443", router)
	//http.ListenAndServeTLS(":8443", "../config/newcert.pem", "../config/privkey.pem", router)

}
