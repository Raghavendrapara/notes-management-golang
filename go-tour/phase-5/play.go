package main

import ("fmt"
 "runtime")

func main() {
	users := map[string]string{"admin": "root"}
	tryToReplace(users)

	fmt.Printf("Address of map m: %p\n", &users)
	fmt.Println("main ", *&users)
	addressPrint(&users)
	fmt.Println("main ", users)
	fmt.Println(runtime.GOOS)

	switchcase("b")
}

func tryToReplace(m map[string]string) {
	m = make(map[string]string)
	m["hack"] = "attack"
}

func addressPrint(ptr *map[string]string) {
	fmt.Println(&ptr)
	fmt.Println(*ptr)

	// *ptr = map[string]string{"n":"n"}
	// ptr = &map[string]string{"m": "n"}
	// fmt.Println(*ptr)

	*ptr = nil
}

/* 
Notice there is no need for break like java
*/
func switchcase(a string){
	switch a{
	case "a" : fmt.Println("A")
	case "b" : fmt.Println("B")
	}
}
