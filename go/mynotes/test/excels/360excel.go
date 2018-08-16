package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"sync"
)

var wg sync.WaitGroup
func main() {
	wg.Add(1)
	a := 0
	go func()  {
		a = createFile(a)
	}()
	wg.Wait()
	fmt.Println(a)
	//readFile()
	//tableFile()
}

//创建文件
func createFile(a int) int {
	a++
	fmt.Println("---",a)
	xlsx := excelize.NewFile()
	// Create a new sheet.
	//index := xlsx.NewSheet("Sheet2")
	// Set value of a cell.
	xlsx.SetCellValue("Sheet1", "A1", "Hello world.")
	xlsx.SetCellValue("Sheet1", "B1", 100)
	// Set active sheet of the workbook.
	//xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
	return a
}

//读取文件
func readFile()  {
	xlsx, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell := xlsx.GetCellValue("Sheet1", "B1")
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	xlsx.SetCellValue("Sheet1", "A2", "Hello ")
	xlsx.SetCellValue("Sheet1", "B2", "Hello ")
	xlsx.SetCellValue("Sheet1", "A3", "456789")
	//xlsx.DeleteSheet("Sheet2")
	fmt.Println("err:",xlsx.Save())
}

//图表文件
func tableFile()  {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	xlsx := excelize.NewFile()
	for k, v := range categories {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	xlsx.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

//插图文件
func pictureFile()  {
	xlsx, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Insert a picture.
	err = xlsx.AddPicture("Sheet1", "A2", "./image1.png", "")
	if err != nil {
		fmt.Println(err)
	}
	// Insert a picture to worksheet with scaling.
	err = xlsx.AddPicture("Sheet1", "D2", "./image2.jpg", `{"x_scale": 0.5, "y_scale": 0.5}`)
	if err != nil {
		fmt.Println(err)
	}
	// Insert a picture offset in the cell with printing support.
	err = xlsx.AddPicture("Sheet1", "H2", "./image3.gif", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`)
	if err != nil {
		fmt.Println(err)
	}
	// Save the xlsx file with the origin path.
	err = xlsx.Save()
	if err != nil {
		fmt.Println(err)
	}
}