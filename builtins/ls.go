package builtins

import (
	"fmt"
	"os"
)

func ListFiles() error {
	fmt.Println("List Files")
	directory := "."		//current directory name

	dir, err := os.Open(directory)	//opens the directory
	if (err != nil) {
		fmt.Println(err)
		return err
	}

	files, err := dir.Readdir(-1)		//reads files and directories inside current WD, return a list of FileInfo objects
	if (err != nil) {
		fmt.Println(err)
		return err
	}

	//file is each FileInfo object returned by Readdir above
	//files is the list of FileInfo objects returned by Readdir
	//loop through each file in files
	for _, file := range files {
		fmt.Println(file.Name());
	}

	//close the directory
	dir.Close()
	return nil
}
