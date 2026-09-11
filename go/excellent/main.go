package main

// 奇数か偶数か(test)
func EvenOrOdd(number int) string {
	if number % 2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}