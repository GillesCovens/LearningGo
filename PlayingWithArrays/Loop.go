package  main

import "fmt"

func main(){
	//this is a Array in long, this can be written shorter :)
	var a [10]int
	a[5] = 222288288282
	fmt.Println(a[5])
	fmt.Println("################################")
	//here is how to declare an array shorter
	numbers := [7]int{1,2,3,4}
	fmt.Println(numbers)

	fmt.Println("################################")
	//here is how to do an good forloop with an print
	for i := 0 ; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	//removing values equal to 0 in my array
	fmt.Println("################################")
	for i := 0; i < len(numbers); {
		
		if numbers[i] == 0 {
			i++
		}else {
			fmt.Println(numbers[i])
			i++
		}
	}

	
}

