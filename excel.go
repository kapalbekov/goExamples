package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
)

func main() {
    excelFileName := "foo.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        fmt.Println("err = ", err.Error())
    }
    for _, sheet := range xlFile.Sheets {
        for i, row := range sheet.Rows {
			fmt.Println("i = ", i, "_________________________________________")
            for _, cell := range row.Cells {
                text := cell.String()
                fmt.Printf("%s\n", text)
            }
        }
    }
}