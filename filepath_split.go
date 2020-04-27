package main
	import (
		"path/filepath"
		"fmt"
	)
	func main (){
		dir, file := filepath.Split("C:/Users/00036639.UNIVERSAL.000/Desktop/go examples")
		fmt.Println("dir = ", dir)
		fmt.Println("file = ", file)
	}