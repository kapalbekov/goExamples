package main 
import (
	"fmt"
	"strings"
	"strconv"
)	

func main(){
	fmt.Println("res ==========", isIINValid("260803302502"))
}

func isIINValid(iin string) bool{	
	if len(iin) != 12 {
		fmt.Println("length of iin must be equal 12!")
		return false
	}
	ves1 := []int{1,2,3,4,5,6,7,8,9,10,11}
	ves2 := []int{3,4,5,6,7,8,9,10,11,1,2}
	var razrIIN [12]int
	iinArr := strings.Split(iin, "")
	sum := 0
	
	//fmt.Println("type of iinArr vals is ", reflect.TypeOf(iinArr[1]))
	
	for i, k := range iinArr {
		val, err := strconv.Atoi(k)
		if err != nil {
			fmt.Println("error :", err.Error())
			return false
		}
		razrIIN[i] = val
	}
	
	for i, k := range ves1 {
		sum = sum + k*razrIIN[i]
	}	
	ctrlSum := sum%11
	
	if ctrlSum == 10 {
		sum = 0
		for i, k := range ves2 {
			//fmt.Println(i, k, "*", razrIIN[i])
			sum = sum + k*razrIIN[i]
		}
		ctrlSum = sum%11
		fmt.Println("Sum2 = ", sum)
		fmt.Println("ctrlSum2 = ", ctrlSum)
	}	
	
	return razrIIN[11]==ctrlSum
	
}