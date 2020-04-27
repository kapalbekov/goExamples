package main
import (
	"regexp"
	"fmt"
)
func main(){
	r, _ := regexp.Compile("\\d{12}")
	fmt.Println(r.FindString("ф42545 авы авф 8608033025954 77"))
	return
}