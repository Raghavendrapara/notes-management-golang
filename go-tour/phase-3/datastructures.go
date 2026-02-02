package main

import (
	"fmt"
)

/*
The "Memory Leak" Risk: This is actually a common production issue. If you load a 100MB file into a slice fileData, and then do:

Go
header := fileData[:100] // Take first 100 bytes
return header
If you return header and keep using it, the entire 100MB array stays in RAM, because header is still pointing to the start of that huge block.

The Fix (if you actually wanted to free memory): You must copy the data to a new, independent slice.

Go
// Create a brand new independent slice
newS := make([]int, len(s))
copy(newS, s)
// Now the old backing array can be garbage collecte
*/
func main() {
	var a [2]string
	a[0] = "abv"
	a[1] = "abcc"
	fmt.Println(len(a))

	p := [2]string{"", ""}
	fmt.Println(p)

	//Note slices are references to array not new arrays per se
	names := [5]string{"ina", "mina", "dika", "amar", "akbar"}

	fmt.Println(names)
	c := names[0:2]
	d := names[2:5]
	fmt.Println(c, d)
	c[1] = "jane"
	fmt.Println(c)
	fmt.Println(names)

	//struct arrays

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(s)
}
