package main

import (
	"fmt"
	"net/http"

	"github.com/boyzhujian/supercmd/controller/curl"
	"github.com/boyzhujian/supercmd/controller/osapi"
	"github.com/gorilla/mux"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "supercmd turns your sever into one giant api server ")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", mainHandler)
	router.HandleFunc("/server/fileexist", osapi.FileexistHandler)
	router.HandleFunc("/server/gethostname", osapi.GethostnameHandler)
	router.HandleFunc("/curl/upload", curl.UploadfileHandler)
	// err := http.ListenAndServe(":10001", router)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// s, err := fs.New()
	// f, err := s.Open("../config/newcert.pem")
	http.ListenAndServeTLS(":8443", "../config/newcert.pem", "../config/privkey.pem", router)
	//	http.ListenAndServeTLS(":8443", f, "../config/privkey.pem", router)

}
