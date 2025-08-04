package utils

import(
	"fmt"
)

func Reader(){
	fmt.Println("Test message for checking imports for cmd/app/main.go")
}

func VariableScanner(){
	var  var1 int
	fmt.Scan(&var1)
	fmt.Println("Scanned variable: ", var1)
}

func TestFunc(){
	var string1 string
	fmt.Scan(&string1)
	fmt.Println("You enter the line: ", string1)
}

func NumberFoo(){
	var num int
	fmt.Println("You enter number: ", num)
}

func TestString(){
	var str_one string
	fmt.Println("Main line is here. Test message")
}
