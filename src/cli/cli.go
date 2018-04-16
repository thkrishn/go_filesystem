/*
* Cli for the go file system 
*/

package main

import "fmt"
import "os"

func main() {

	for {
		fmt.Print("##> ")
		var input string
		
		fmt.Scanln(&input)

		switch input {
			case "h":
				helpList()
			case "quit":
				os.Exit(0)
			case "q":
				os.Exit(0)
			case "ls":
				//listFiles()
			default:
				fmt.Print(input + " ")
				fmt.Println("unrecognized conmmand line option")
				//panic("unrecognized input value")
		}
	}
	//fmt.Println(sum)
}

func helpList() {
	fmt.Println(" h  -- help ")
	fmt.Println(" ls -- list files ")
}
