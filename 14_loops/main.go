package main

import "fmt"

func main() {
	fmt.Println("Lets Go Loops")

	// for d:=0;d<10;d++{
	// 	fmt.Println(d)
	// }

	// days:=[]string{"Sunday","Saturday","Friday"}

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	// for i,day:= range days{
	// 	fmt.Println(i,day)
	// }

	rougueValue :=1
	// for rougueValue<10{
	// 	if rougueValue==5{
		
	// 		break
	// 	}
	// 	fmt.Println(rougueValue)
	// 	rougueValue++
	// }

	// for rougueValue<10{
	// 	if rougueValue==5{
	// 		rougueValue++
	// 		continue
	// 	}
	// 	fmt.Println(rougueValue)
	// 	rougueValue++
	// }

	for rougueValue<10{
		if rougueValue==5{
			goto myLabel
		}
		fmt.Println(rougueValue)
		rougueValue++
	}
	myLabel:
		fmt.Println("This is myLabel")
}