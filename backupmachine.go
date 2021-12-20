package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: backupmachine <path file sumber> <path file tujuan>")
		} else if errors.Is(err, os.ErrNotExist) {
			copy(os.Args[1], os.Args[2])
			fmt.Println("Copying...")
			time.Sleep(3 * time.Second)
			err = os.Chmod(os.Args[2], 0000)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Success!")
			}
		} else {
			fmt.Println("Wut??")
		}

	}
}
