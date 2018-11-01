
package main
import "fmt"
import (
	"bufio"
	"os"
  )
func main() {
	var x string
	fmt.Println("Enter the string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() //Accept spaces as input too
	x=scanner.Text()
	fmt.Println(x); //Checking if input came in correctly
	switch {
	case len(x)<3: fmt.Println("Not found!") //Either a or n can't possibly be present
	case (x[0]=='i'||x[0]=='I')&&(x[len(x)-1]=='n'||x[len(x)-1]=='N'): 
		var flag bool=false
		var pos int=1
		for pos<len(x)-1 {
			if x[pos]=='a'||x[pos]=='A' {
					fmt.Println("Found!")
					flag=true
					break
			}
			pos=pos+1
		}
		if flag==false {
			fmt.Println("Not found!")
		}
	default: 
			fmt.Println("Not found here!")
	}
}