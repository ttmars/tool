package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main()  {

}

// 重命名目录的所有文件
func renameDirFile()  {
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
