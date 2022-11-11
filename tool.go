package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func RenameDirFile(p string)  {
	err := filepath.Walk(p, func(path string, info fs.FileInfo, err error) error {
		if path != p {
			newName := path + ".c"
			if err := os.Rename(path, newName); err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("success:%s\n", path)
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func FileServer(port string, path string)  {
	router := gin.Default()
	router.StaticFS("/", http.Dir(path))
	router.Run(":"+port)
}

func GetRandomString2(n int) string {
	rand.Seed(time.Now().UnixNano())
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
