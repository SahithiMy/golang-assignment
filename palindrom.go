// You can edit this code! // Click here and start typing.
package main

import "fmt"

func main() {
	var a string
	fmt.Print("enter string")
	fmt.Scanln(&a)
	i := 0
	for ; i < len(a); i++ {
		if a[i] == a[len(a)-1-i] {
			continue
		} else {
			break
		}
	}
	if i == len(a) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Not palindrom")
	}
}
