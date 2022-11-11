package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// 重命名目录的所有文件
func RenameDirFile()  {
	p := "C:\\Users\\Ares\\Desktop\\C\\project\\figlinks"
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

// fileServer("8888", "./")
func fileServer(port string, path string)  {
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
