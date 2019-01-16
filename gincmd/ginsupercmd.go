package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/boyzhujian/supercmd/controller/curl"
	"github.com/boyzhujian/supercmd/controller/osapi"
	"github.com/gin-gonic/gin"
)

var (
	// Initialization of the working directory. Needed to load asset files.
	filePath = determineWorkingDirectory()
	// File names for the HTTPS certificate
	certFilename = filepath.Join(filePath, "cert.pem")
	keyFilename  = filepath.Join(filePath, "key.pem")
	port         string
)

func init() {

}

func determineWorkingDirectory() string {
	var certPath string
	// Check if a custom path has been provided by the user.
	flag.StringVar(&certPath, "certpath", "", "Specify a cert path to the serverkey and servercrt files")
	flag.StringVar(&port, "port", ":8180", "input the port to listen")
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

//maybe have multiple exec if run in gorutine,if prog is daemon ,it stuck in goroutin
func runprog(prog string) {
	cmd := exec.Command(prog, "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

func main() {
	//fmt.Println(certFilename)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "this is gin verions supercmd") })
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/printrequest", func(c *gin.Context) {
		requestDump, _ := httputil.DumpRequest(c.Request, true)
		c.String(http.StatusOK, string(requestDump))
	})
	r.GET("/server/fileexist", gin.WrapF(osapi.FileexistHandler))
	r.GET("/server/gethostname", gin.WrapF(osapi.GethostnameHandler))
	r.GET("/server/catfile", func(c *gin.Context) {
		f := c.Query("filename")
		content, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusOK, string(content))
	})

	r.GET("/server/execprog", func(c *gin.Context) {
		prog := c.Query("exec")
		go runprog(prog)
		c.String(http.StatusOK, prog)
	})
	r.GET("/curl/upload", gin.WrapF(curl.UploadfileHandler))

	r.Run(port) // listen and serve on 0.0.0.0:8080
	//r.RunTLS(":10000", certFilename, keyFilename)
}
