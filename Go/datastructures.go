package main
// Task: Program must demonstrate the major data structures of the language

import (
	"fmt"
)

/*
	Arrays are of a fixed size which can't be changed during runtime because of this
	only the first 256 values will be used in the array and the rest of the array
	will be filled with zeros.
*/
func testArray(values []int) {
	// Initializing and copying array from values
	array := [256]int{}
	copy(array[:], values)	
	
	fmt.Println(" -- Array Operations -- ")
	
	// Accessing array elements
	fmt.Println("Accessing 0th, 1st, and 2nd array elements")
	fmt.Printf("\t array[0] = %d\n", array[0])
	fmt.Printf("\t array[1] = %d\n", array[1])
	fmt.Printf("\t array[2] = %d\n", array[2])
	
	// Printing array length (fixed at 256)
	fmt.Printf("\nArray length: %d\n", len(array))
	
	// Iterating over array using range
	fmt.Println("\nIterating over array using range")
	for _, val := range array {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	
	// Printing array
	fmt.Println("\nPrinting array:", array, "\n")
	
}

/*
	Slices are more flexible than arrays but like arrays can only hold one data type.
	All slices reference an underlying array.
*/
func testSlice(slice []int) {
	fmt.Println(" -- Slice Operations -- ")
	
	if len(slice) > 0 {
		// Accessing slice elements
		fmt.Println("Accessing some slice elements")
		fmt.Printf("\tslice[0] = %d\n", slice[0])
		if len(slice) > 1 {
			fmt.Printf("\tslice[1] = %d\n", slice[1])
		}
		if len(slice) > 1 {
			fmt.Printf("\tslice[2] = %d\n", slice[2])
		}
		
		// Printing slice length 
		fmt.Printf("\nSlice length: %d\n", len(slice))
		
		// Printing different slices
		fmt.Printf("\nPrinting slices of slice\n")
		if len(slice) > 3 {
			fmt.Printf("\tslice[1:3] = %v\n", slice[1:3])
		}
		if len(slice) > 7 {
			fmt.Printf("\tslice[3:7] = %v\n", slice[3:7])
		}
		if len(slice) > 20 {
			fmt.Printf("\tslice[:20] = %v\n", slice[:20])
		}
		
	
	} else {
		fmt.Println("Slice is empty")
	}
	// Printing slice
	fmt.Println("\nPrinting slice:", slice, "\n")
}
/*
	Maps are golangs version of dictionary or hashmaps and are used for storing the
	relations between two or more types.
*/
func testMap(values []int) {
	// Creating a map from values to the index of the last occurance
	mapping := map[int]int{}
	for i, val := range values {
		mapping[val] = i
	}
	
	fmt.Println(" -- Map Operations -- ")
	if val, yes := mapping[1]; yes {
		fmt.Printf("\tmapping[1] = %d\n", val)
	} else {
		fmt.Printf("\tmapping does not contain 1\n")
	}
	if val, yes := mapping[5]; yes {
		fmt.Printf("\tmapping[5] = %d\n", val)
	} else {
		fmt.Printf("\tmapping does not contain 5\n")
	}
	if val, yes := mapping[16]; yes {
		fmt.Printf("\tmapping[16] = %d\n", val)
	} else {
		fmt.Printf("\tmapping does not contain 16\n")
	}
	
	fmt.Printf("\nPrinting map: ")
	fmt.Println(mapping)
	
	fmt.Println()
}

/*
	Structs can hold multiple data types, the test struct below will hold an id and value
	both of which are integers. The test struct will use the values index as the id and
	value as value.
*/
func testStruct(values []int) {
	fmt.Println(" -- Struct Operations -- ")
	type TestStruct struct {
		id int
		value int
	}
	
	structs := make([]TestStruct, len(values))
	for i, val := range values {
		structs[i] = TestStruct{i, val}
	}
	// Print struct ids
	fmt.Printf("Printing struct ids: ")
	for _, s := range structs {
		fmt.Printf("%d ", s.id)
	}
	
	// Print struct values
	fmt.Printf("\n\nPrinting struct values: ")
	for _, s := range structs {
		fmt.Printf("%d ", s.value)
	}
	
	// Print list of structs
	fmt.Printf("\n\nPrinting structs: ")
	fmt.Println(structs)
	
	fmt.Println()
}

func testAll(values []int) {
	testArray(values)
	testSlice(values)
	testMap(values)
	testStruct(values)
}

func main() {
	testAll([]int{2,4,16,256,65536})
	testAll([]int{1,2,3})
	testAll([]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1, 1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1})
	testAll([]int{})// Empty data structure case

}
