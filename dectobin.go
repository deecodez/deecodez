package main
 
import "fmt"
 
func main() {
        fmt.Print("Please Enter Number to convert: ")   //Print function is used to display output in same line
        var num int    
		fmt.Scanln(&num)                  // Take input from user
		
		bin := fmt.Sprintf("%b", num)
    
        fmt.Println("Binary: ", bin)           // Addition of two string   
}