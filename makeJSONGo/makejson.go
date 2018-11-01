package main
import (
	"fmt";
	"encoding/json"
	"bufio"
	"os"
)
func main() {
	details:=make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter name:")
	scanner.Scan()
	name:=scanner.Text()
	fmt.Println("Enter address:")
	scanner.Scan()
	address:=scanner.Text()
	details["name"]=name
	details["address"]=address
	ans,_:=json.Marshal(details)
	fmt.Println(ans)
	fmt.Println(string(ans))
}