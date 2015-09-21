package mdb

import (
	"fmt"
)


// PrintMap
func PrintMap(m map[string]string){
	for i, _:=range m {
		fmt.Println(m[i])
	}
}

//PrintCloumns
func PrintCloumns(cols []string){
	ncols:=len(cols) - 1
	for i:=range cols{
		if i == ncols{
			fmt.Printf("%s", cols[i])
		}else{
			fmt.Printf("%s,", cols[i])
		}
	}
	fmt.Println()
}

// PrintRowsArray
func PrintRowsArray(data [][]string){
	rows := len(data)
	fmt.Println()
	for row := 0; row < rows; row++ {
		cols:=len(data[row])
		for col := 0; col < cols; col++ {
			if col==cols-1{
				fmt.Printf("%s",data[row][col])
			}else{
				fmt.Printf("%s,",data[row][col])
			}
		}
		fmt.Println()
	}
}

