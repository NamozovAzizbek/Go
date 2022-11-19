package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanf("%d", &a)
	if a == 0{
		fmt.Print(-1)
	}else{
	fmt.Print( top(a))
	}
}
func top(a int )(res int){
	if a < 0 {
		a = -a
	}
	for i := 1; i*i <= a; i ++{
			if a % i == 0{
				res ++
				if i*i != a{ 
					res ++
				}
			}
	}
	return
}