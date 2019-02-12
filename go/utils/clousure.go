package p1

import(
	"fmt"
)

func main(){
	// usage
	p:=Add2();
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
}



func Add2() func() int {
	a := 1
	return func() int {
		a = a+1;
		return a
	}
}
