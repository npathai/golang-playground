package main

func main() {
	// Normal loop till condition
	var i int
	for i < 5 {
		println(i)
		i++
	}

	// Break out of the loop
	var j int
	for j < 5 {
		println(j)
		if j == 3 {
			break
		}
		j++
	}

	// Continuing in the loop
	var k int
	for k < 5 {
		k++
		if k == 3 {
			continue
		}
		println(k)
	}


	// Loop till condition with post clause
	// Initialization of variable in separate line
	var m int
	for ;m < 5; m++ {
		println(m)
	}

	// Initializing the variable inline in for loop
	for n:= 0; n < 5; n++ {
		println(n)
	}

	// Creating infinite loops
	var p int
	for {
		if p == 5 {
			break
		}
		p++
		println(p)
	}

	// Iterating a collection
	arr := []int{1, 2}

	println("Iterating a collection")
	for l := 0; l < len(arr); l++ {
		println(arr[l])
	}

	println("Better way of iterating a collection")
	for k, v := range arr {
		println(k, v)
	}

	println("Iterating over a map")
	someMap := map[string]int{"ABC": 123}
	for key, val := range someMap {
		println(key, val)
	}
}
