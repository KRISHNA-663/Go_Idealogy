package main

import "fmt"

func main(){
	var n int
	fmt.Printf("Enter a number : ")
	fmt.Scanf("%d",&n)
	if n%2==0{
		fmt.Printf("%d is even",n)
	}else{
		fmt.Printf("%d is odd",n)
	}
}