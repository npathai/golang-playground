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
}
