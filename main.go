package main

import "fmt"

func count(n int) (int){
	counter := 0
	for i := 0; i < n; i+=3 {

			counter += i
		}
	for i:=0; i < n; i+=5{

		counter += i
	}


	return counter
}

 func main() {
	fmt.Println(count(1000))
}
/*

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
