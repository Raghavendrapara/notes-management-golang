package main

import (
	"fmt"
	"strings"
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

	/* We can append to slices
	However it creates new allocation on overflow
	so need to check if we can experience performance issues
	*/

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

	rangeOnArrays()

	// pic.Show(Pic)
	maps()

	WordCount("abc def ghi")

}

func rangeOnArrays() {
	var pow = []int{1, 2, 4, 8, 16}
	for i := range pow {
		fmt.Println(i)
	}
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			pic[i][j] = uint8(i*i + j*j)
		}
	}
	return pic
}

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func maps() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{1, 5}
	fmt.Println(m)

	a := make(map[string]int)
	a["a"] = 1
	v, ok := a["a"]
	fmt.Println(a, v, ok)
}
func WordCount(s string) map[string]int {

	a := make(map[string]int)
	if len(s) == 0 {
		return a
	}
	words := strings.Fields(s)
	for f := range words {
		fmt.Println(words[f])
		// v,ok :=a[f]
		// if ok {
		// 	a[f]+=1
		// }else{
		// 	a[f]=1
		// }
	}

	return a
}
