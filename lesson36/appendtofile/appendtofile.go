package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("./writebyline/lines", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newline := "File handling is easy."
	_, err = fmt.Fprintln(f, newline)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}
