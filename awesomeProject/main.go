package main

import "fmt"

func isPalindrome(n int) bool {
	original := n
	reversed := 0

	for n > 0 {

		digit := n % 10
		reversed = reversed*10 + digit
		n = n / 10
	}

	return original == reversed

}

func reverseInteger(n int) int {

	znak := 1
	if n < 0 {
		znak = -1
		n = -n
	} else if n == 0 {
		return 0
	}

	reversedint := 0
	for n > 0 {

		digit := n % 10
		reversedint = reversedint*10 + digit
		n = n / 10
	}

	return reversedint * znak

}

func canHold(width1 int, length1 int, width2 int, length2 int) bool {
	rectangle1 := width1 * length1
	rectangle2 := width2 * length2
	if rectangle1 > rectangle2 {
		return true
	}

	return false

}

func guessNumber(n int) string {
	switch n {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	default:
		return "eto drugoe"

	}
}

func pow(n int, pow int) int {

	if pow == 0 {
		return 1
	}
	if pow < 0 {
		return 0
	}

	res := 1
	for i := 0; i < pow; i++ {
		res = res * n

	}

	return res

}

func fibonacci(n int) int {

	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	f1 := 0
	f2 := 1
	for i := 2; i < n; i++ {

		f1, f2 = f2, f1+f2
	}

	return f2

}

func fibonaccirec(n int) int {

	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	return fibonaccirec(n-1) + fibonaccirec(n-2)

}
func array2d(arr [2][3]int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}

func main() {

	fmt.Println("Palindrome:", isPalindrome(454))
	fmt.Println("Reversint:", reverseInteger(1534236469))
	fmt.Println("Canhold:", canHold(10, 20, 10, 20))
	fmt.Println("guessNumber:", guessNumber(9))
	fmt.Println("stepen:", pow(2, 4))
	fmt.Println("fibonacci:", fibonacci(10))
	fmt.Println("fibonaccirec:", fibonaccirec(10))

	arr := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	array2d(arr)
}

//fdsdf
