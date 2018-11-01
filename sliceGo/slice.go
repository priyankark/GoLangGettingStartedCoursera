package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli:=[]int{}
	var inp string 
	for true {
		fmt.Println("Please enter a number:")
		fmt.Scan(&inp);
		if inp=="x" {
		 break
		}
		x,_:=strconv.Atoi(inp)
		sli=append(sli,x)
		sort.Ints(sli)
		fmt.Println(sli)	
	}
}