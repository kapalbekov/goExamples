package main
import
(
	"time"
	"fmt"
)

func main(){
	t := time.Unix(1592762400, 0)
	fmt.Println("time = ", t)
	/*t = time.Unix(1589889600, 0)
	fmt.Println("time = ", t)
	t = time.Unix(1592740800, 0)
	fmt.Println("time = ", t)
	t = time.Unix(1595160000, 0)
	fmt.Println("time = ", t)
	t = time.Unix(1597838400, 0)
	fmt.Println("time = ", t)*/
	//fmt.Println("time unix = ", time.Now().Add(19 * time.Hour).Unix())
	//t := time.Unix(time.Now().Add(10 * time.Hour).Unix(), 0)
	//fmt.Println("time = ", t)
	//fmt.Println(reflect.TypeOf(t.Format("2006-01-02")))
	//fmt.Println(t.Format(time.RFC3339))
	
	//fmt.Println(time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC).Unix())
}