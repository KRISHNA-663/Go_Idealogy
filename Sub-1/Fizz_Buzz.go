package main

import(
	 "fmt" 
	 "strconv"
)

func main(){
	var n int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%d",&n)
	fmt.Printf("%s",FizzBuzz(n))
}

func FizzBuzz(n int) string{
	if n%3==0 && n%5==0{
		return "FizzBuzz"
	}else if n%3==0{
		return "Fizz"
	}else if n%5==0{
		return "Buzz"
	}
	return strconv.Itoa(n)
}