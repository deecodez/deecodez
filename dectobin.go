package main
 
import "fmt"
 
func main() {
        fmt.Print("Please Enter Number to convert: ")   
        var num int    
		fmt.Scanln(&num)                
		
		bin := fmt.Sprintf("%b", num)
    
        fmt.Println("Binary: ", bin)              
}