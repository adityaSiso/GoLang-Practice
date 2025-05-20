package functions

// Alias
type functionType func(int) int

// Passing functions as param.
func transformNumbers(numbers *[]int, transform functionType) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func function() {
	numbers := []int{1, 2, 3, 4, 5}
	requestFunc := getTransformer("double")
	numbers = transformNumbers(&numbers, requestFunc)

	// Anonymous Function
	anonymousFunction := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
}

// Returning a function.
func getTransformer(req string) functionType {
	transformerMap := map[string]functionType{
		"double": double,
		"triple": triple,
	}
	return transformerMap[req]
}

func double(i int) int {
	return i * 2
}

func triple(i int) int {
	return i * 3
}

// Closures
func createTranformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor //**************
	}
}

// Recursion
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return factorial(number-1) * number
}

// Variadic Functions (n numbers of parameters)
func sumUp(numbers ...int) int {
	return 1
}

// Splitting slice into parameters
// numbers := []int{1, 2, 3, 4, 5}
// function(1, numbers...)
