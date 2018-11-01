package main
import (
		"fmt"
		"os"
		"bufio"
		"strings"
		"io"
)

func main() {
	type Names struct {
		fname string
		lname string
	}

	var filename string
	fmt.Println("Please provide the file name: ")
	fmt.Scan(&filename)
	file,_:=os.Open(filename)
	reader:=bufio.NewReader(file)
	sli:=[]Names {}
	/*l, _, err := reader.ReadLine()
	line:=string(l)
	nameSlice:=strings.Split(line," ")
	sli=append(sli,Names{fname: nameSlice[0],lname: nameSlice[1]}) */
	for  {
		l, _, err := reader.ReadLine()
		if err==io.EOF {
			break
		}
		
		line:=string(l)
		nameSlice:=strings.Split(line," ")
		sli=append(sli,Names{fname: nameSlice[0],lname: nameSlice[1]})
	}
	fmt.Println(sli) //checking contents of the final slice
	for _,ele:= range sli {
		fmt.Println(ele.fname,ele.lname); //iterating and displaying final results
	}

}