package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/boyzhujian/supercmd/controller/osapi"
	"github.com/gin-gonic/gin"
)

var (
	// Initialization of the working directory. Needed to load asset files.
	filePath = determineWorkingDirectory()
	// File names for the HTTPS certificate
	certFilename = filepath.Join(filePath, "cert.pem")
	keyFilename  = filepath.Join(filePath, "key.pem")
)

func init() {

}

func determineWorkingDirectory() string {
	var certPath string
	// Check if a custom path has been provided by the user.
	flag.StringVar(&certPath, "certpath", "", "Specify a cert path to the serverkey and servercrt files")
	flag.Parse()
	// Get the absolute path this executable is located in.
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatal("Error: Couldn't determine working directory: " + err.Error())
	}
	// Set the working directory to the path the executable is located in.
	os.Chdir(executablePath)
	// Return the user-specified path. Empty string if no path was provided.
	return certPath
}

func main() {
	fmt.Println(certFilename)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "this is gin verions supercmd") })
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/server/fileexist", gin.WrapF(osapi.FileexistHandler))
	r.GET("/server/gethostname", gin.WrapF(osapi.GethostnameHandler))

	//r.Run() // listen and serve on 0.0.0.0:8080
	r.RunTLS(":10000", certFilename, keyFilename)
}
