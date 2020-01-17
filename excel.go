package goutils

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type Excel struct{
	File *excelize.File
}
var column map[int]string = map[int]string{
	0 : "A",
	1 : "B",
	2 : "C",
	3 : "D",
	4 : "E",
	5 : "F",
	6 : "G",
	7 : "H",
	8 : "I",
	9 : "J",
	10: "K",
	11: "L",
	12: "M",
	13: "N",
	14: "O",
	15: "P",
	16: "Q",
	17: "R",
	18: "S",
	19: "T",
	20: "U",
	21: "V",
	22: "W",
	23: "X",
	24: "Y",
	25: "Z",
}

func NewExcelInstance() *Excel{
	return &Excel {
		File : excelize.NewFile(),
	}
}

func (this *Excel)WriteExcelTitle(sheet string, title []string) (*Excel){
	for k,value:= range title{
		this.File.SetCellValue(sheet, column[k] + "1", value)
	}
	
	return this
}

func (this *Excel)WriteExcelData(sheet string, data [][]interface{}) (*Excel){
	index:=1
	for _,array:= range data{
		index++
		for i,value := range array{
			this.File.SetCellValue(sheet, column[i] + strconv.Itoa(index), value)
		}
	}
	
	return this
}

func CreateExcel(savePath string, title []string, data [][]interface{}) error{
	f := NewExcelInstance()   
	f.WriteExcelTitle("Sheet1", title)
	f.WriteExcelData("Sheet1", data)
	if err := f.File.SaveAs(savePath); err != nil {
        return err
	}
	return nil
}