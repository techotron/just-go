// Following this guide: https://www.youtube.com/watch?v=C8LgvuEBraI

package main

import (
	"errors"
	"fmt"
	"math"
)

// Struct =================================================================================
// A struct is a collection of fields. It's simply a data structure so it's comparable to json or yaml. It should be defined outside of a function(not sure why?)
type person struct {
	name string
	age int
}

func main() {
	// Variables ================================================================================
	// Declaring a var as an integer like this:
	var x int = 5

	// Is the same as this. The type 'int' is inferred.
	y := 7
	var sum int = x + y

	if sum < 100 {
		fmt.Println(sum)
	} else {
		fmt.Println("Woah - large number!")
	}

	// Arrays ==========================================================================================
	// This will create a fixed length array with 5 empty elements. In GoLong, empty = 0, so it'll output "[0 0 0 0 0]"
	var array [5]int
	fmt.Println(array)

	// Set object in position '2' to equal 7. Will output "[0 0 7 0 0]"
	array[2] = 7
	fmt.Println(array)

	// Create a fixed length array with specified values.
	array2 := [5]int{5, 4, 3, 2, 1}
	fmt.Println(array2)

	// You can't add to a fixed length array because the length is part of the type. To make the array more versatile, you can create a 'slice' like this
	array3 := []int{1, 2, 3, 4, 5}
	// Use the built in function 'append' to add an element to the slice:
	array3 = append(array3, 6)
	fmt.Println(array3)

	// Maps (key/value pairs, like dictionaries in Python)
	// Syntax for creating the map is: map[key data type]value data type. You need to use the built-in 'make' function and give it the 'map' type eg:
	map1 := make(map[string]int)

	// Set values for the key/values like this
	map1["two"] = 2
	map1["five"] = 5

	// Output the whole map. This will look like: "map[two:2 five:5]"
	fmt.Println(map1)

	// Out the value for a particular key:
	fmt.Println(map1["two"])

	// Delete something from the map:
	delete(map1, "five")
	fmt.Println(map1)

	// Loops ==========================================================================================
	// The only type of loop in Go is the 'for' loop. The syntax is 'for counter; stopping condition; counter increment'
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// To turn this into a while loop, you can rewrite the for loop like this:
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Iterate over each element in an array or slice. This will print the index and the value of each element in the array, 'array4':
	array4 := []string{"a", "b", "c"}
	for index, value := range array4 {
		fmt.Println("index:", index, "value:", value)
	}

	// Iterate over key/values in a map:
	map2 := make(map[string]string)
	map2["a"] = "alpha"
	map2["b"] = "bravo"
	map2["c"] = "charlie"

	for mykey, myvalue := range map2 {
		fmt.Println("key:", mykey, "value:", myvalue)
	}

	// Use a function from outside the 'main' function ====================================================================
	result := sum_calc(5, 10)
	fmt.Println(result)

	// Return the square root using function outside of the main function. This is saying that we're expecting 2 returned values, result and error.
	squareRootResult, err := square_root(16)

	// If the returned error is not 'nil', then something went wrong:
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(squareRootResult)
	}

	// Create the struct from the struct 'type' defined above:
	p := person{name: "Eddy", age: 33}
	fmt.Println(p)

	// Print only the name element from the person struct:
	fmt.Println(p.name)

	// Pointers =========================================================================================================
	// Print out the memory address of where the variable "pointer_int" is referenced
	pointer_int := 7
	fmt.Println(&pointer_int)

	// Attempt to add pointer_int by one using the "increment_by_one" function:
	increment_by_one(pointer_int)
	fmt.Println(pointer_int)

	// The above doesn't work as we might expect because the value of "pointer_int" is just copied within the "increment_by_one" function. The function doesn't change the value of "pointer_int". We can modify "pointer_int" by using the pointer reference instead. This requires the function which modifies the variable to expect a pointer ref by prefixing the type with an astrix:
	increment_pointer_by_one(&pointer_int)
	fmt.Println(pointer_int)
}

// Arguements in the function need to have the type added in the brackets. The returned value needs to have the type specified outside of the brackets:
func sum_calc(x int, y int) int {
	return x + y
}

// Create function which returns a float64 type and and error type:
func square_root(x float64) (float64, error) {
	// Return an error if x is negative
	if x < 0 {
		// We're returning an error, so we need to add the "errors" package to the list of imports at the top
		return 0, errors.New("Undefined for negative numbers")
	}

	// If x is positive, then calculate the square root using the "math" package (needs to be added to the imports at the top).
	// nil here is is to satisfy the requirement that 2 returned objects are specified in the function. We don't want to return an error so we use "nil".
	return math.Sqrt(x), nil
}

// Example function for demonstrating a pointer
func increment_by_one(x int) {
	x++
}

// Set the type of x to be a pointer using astrix. The references the pointer instead of the value
func increment_pointer_by_one(x *int) {
	// Deference the memory address by using another astrix (otherwise we'll be trying to add 1 to the memory address)
	*x++
}

