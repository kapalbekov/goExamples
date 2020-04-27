package main
import (
	"fmt"
	"strconv"
)
func main(){
	p, _ := strconv.ParseInt("-1", 10, 64)
	fmt.Println("convert to Int = ", p)
	var mzp, okl, res int64
	//mrp = 2525
	mzp = 42500
	
	
	okl = 372543 //612035 //735495 612035
	res = okl - (okl/10) - (okl-mzp-okl/10)/10
	fmt.Println("result = ", res)
	
	
	res = 1000000
	okl = (100*res-10*mzp)/81
	fmt.Println("okl = ", okl)
}