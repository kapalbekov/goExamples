package main
import "fmt"

func setVal(b *int, c int) {
    *b = 2
    c = 4
    fmt.Printf("B from setVal(). Poiner to: %p, val: %v\n", b, *b)
    fmt.Printf("C from setVal(). Addr: %p, val: %v\n", &c, c)
}
func main() {
    a := 1
    b := &a
    c := 3
    fmt.Println("Init values")
    fmt.Printf("A from main(). Addr: %p, val: %v\n", &a, a)
    fmt.Printf("B from main(). Poiner to: %p, val: %v\n", b, *b)
    fmt.Printf("C from main(). Addr: %p, val: %v\n", &c, c)
    fmt.Println("Changed values")
    setVal(b, c)
    
    fmt.Printf("A from main(). Addr: %p, val: %v\n", &a, a)
    fmt.Printf("B from main(). Poiner to: %p, val: %v\n", b, *b)
    fmt.Printf("C from main(). Addr: %p, val: %v\n", &c, c)
	
	fmt.Printf("ABC = %p, %p, %p\n", &a, &b, &c)
}

/*
	A from main(). Addr: ADDRESSAAAA, val: 1
	B from main(). Poiner to: ADDRESSAAAA, val: 1
	C from main(). Addr: ADDRESSCCCCCCC, val: 3
	
	Changed values
	B from setVal(). Poiner to: ADDRESSAAAA, val: 2
	C from setVal(). Addr: ADDRESSC222222, val: 4
	
	A from main(). Addr: ADDRESSAAAA, val: 2
	B from main(). Poiner to: ADDRESSAAAA, val: 2
	C from main(). Addr: ADDRESSCCCCCCC, val: 3
	
	ABC = ADDRESSAAAA, ADDRESSBBBB, ADDRESSCCCCCCC
	
*/