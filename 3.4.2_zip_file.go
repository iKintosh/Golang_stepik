package main

import (
	"archive/zip"
	"fmt"
	"encoding/csv"

)

func main()  {
	zip_file, err := zip.OpenReader("task.zip")
	if err != nil {
        fmt.Println(err)
    }
    defer zip_file.Close()

	for _, f := range zip_file.File {
		if f.FileInfo().IsDir() == false {
			file, _ := f.Open()
			r, _ := csv.NewReader(file).ReadAll()
			if len(r) >= 4 {
				fmt.Print(r[4][2])
				fmt.Print(f.FileInfo().Name())
			}
		}
	}
}