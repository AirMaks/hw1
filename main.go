package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	//path := os.Args[1]
	//printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	// err := tree(out, path, printFiles)
	// if err != nil {
	// 	panic(err.Error())
	// }

	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := tree(arg, "")
		if err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
	}

}

func tree(root, indent string) error {

	fi, err := os.Stat(root)

	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}

	if fi.Name() == "hw1.md" {
		return nil
	}

	fmt.Println(fi.Name())
	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not read dir %s: %v", root, err)
	}

	for i, fi := range fis {
		add := "│   "

		if i == len(fis)-1 {
			fmt.Printf(indent + "└───")
			add = "    "

		} else if fi.Name() == "hw1.md" {
			fmt.Printf(indent + "")
		} else {
			fmt.Printf(indent + "├───")
		}

		if err := tree(filepath.Join(root, fi.Name()), indent+add); err != nil {
			return err
		}
	}

	return nil

}
