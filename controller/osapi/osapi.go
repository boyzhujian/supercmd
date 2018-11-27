package osapi

import (
	"fmt"
	"net/http"
	"os"
)

//FileexistHandler receive one filename  test if that file exist,return okokok if exist nonono if not exist
func FileexistHandler(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Query().Get("filename")
	fmt.Println(f)
	finfo, err := os.Stat(f)
	fmt.Println(finfo)
	fmt.Println(err)
	if err != nil {
		fmt.Println("nonono")
		fmt.Fprintln(w, "nonono")

	} else {
		fmt.Println("okokok")
		fmt.Fprintln(w, "okokok")
	}

}

//Gethostname return serverhostname
func GethostnameHandler(w http.ResponseWriter, r *http.Request) {
	servername, _ := os.Hostname()
	fmt.Fprint(w, servername)
}
