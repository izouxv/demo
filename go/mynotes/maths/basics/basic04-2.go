package main
import (
	"fmt"
	"log"
	"strconv"
)
func main() {
	s := "FRB"
	//s := "R2(B2(LF)BF2(BF))FBF"
	//s := "FRRB"
	//s := "FRRBLFRB"
	//s := "FRRRBLLFRBLF"
	//s := "R2(B2(LF))F"
	cmds := getCmd(s)
	fmt.Println(cmds)
	x, y := getXY(cmds)
	fmt.Println("x:", x, "y:", y)
}
func getXY(s string) (int, int) {
	//aa[]  {y,-x,-y,x}
	var aa [4]int
	var k, tem int
	for _, v := range s {
		a := string(v)
		switch a {
		case "F":
			tem = (4 - (k % 4)) % 4
			aa[tem]++
		case "B":
			tem = (4 - (k % 4)) % 4
			aa[tem]--
		case "R":
			k++
		case "L":
			k--
		}
	}
	return aa[3] - aa[1], aa[0] - aa[2]
}
func getCmd(s string) string {
	var stack []string
	var tem string
	for _, v := range s {
		a := string(v)
		if a >= "0" && a < "9" {
			stack = append(stack, tem)
			stack = append(stack, a)
			tem = ""
			continue
		}
		if a == "(" {
			if tem == "" {
				continue
			}
			stack = append(stack, tem)
			tem = ""
			continue
		}
		tem += a
		if a == ")" {
			tem = tem[:len(tem)-1]
			if stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
				tem = getCopyStr(stack[len(stack)-1], tem)
				tem = stack[len(stack)-2] + tem
			} else {
				stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
			}
			stack = stack[:len(stack)-2]
		}
	}
	return tem
}
func getCopyStr(n, str string) string {
	nn, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < nn; i++ {
		str += str
	}
	return str
}
