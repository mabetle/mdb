package mdb

import (
	"fmt"
)

// PrintMap prints map
// map has no sequence
func PrintMap(m map[string]string) {
	for i, _ := range m {
		fmt.Println(m[i])
	}
}

func mapToArray(ma []map[string]interface{}, columns []string) (data [][]string) {
	// has columns
	if len(columns) > 0 {
		data = append(data, columns)
	}
	// loop datas
	for _, m := range ma {
		rowData := make([]string, len(columns))
		for index, col := range columns {
			if v, b := m[col]; b {
				rowData[index] = GetString(v)
			}
		}
		data = append(data, rowData)
	}
	return
}

// PrintMapWithColumns
func PrintMapWithColumns(m []map[string]interface{}, columns []string) {
	data := mapToArray(m, columns)
	PrintArrayFriendly(data)
}

//PrintCloumns
func PrintCloumns(cols []string) {
	ncols := len(cols) - 1
	for i := range cols {
		if i == ncols {
			fmt.Printf("%s", cols[i])
		} else {
			fmt.Printf("%s,", cols[i])
		}
	}
	fmt.Println()
}

// PrintRowsArray
func PrintRowsArray(data [][]string) {
	rows := len(data)
	fmt.Println()
	for row := 0; row < rows; row++ {
		cols := len(data[row])
		for col := 0; col < cols; col++ {
			if col == cols-1 {
				fmt.Printf("%s", data[row][col])
			} else {
				fmt.Printf("%s,", data[row][col])
			}
		}
		fmt.Println()
	}
}
