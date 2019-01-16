package curl

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

//filename   ,target  two querystring,using curl to upload file
func UploadfileHandler(w http.ResponseWriter, r *http.Request) {
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
