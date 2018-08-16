package main
import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//L左转 R右转，F前一步 B后一步
func main() {
	robot := &Robot{}
	//s := "FRB"
	//s := "R2(B2(LF)BF2(BF))FBF"
	//s := "FRRB"
	//s := "FRRBLFRB"
	//s := "FRRRBLLFRBLF"
	s := "R2(B2(LF))F"
	//RBLFLFBLFLFF
	//RBLFLFBFBFBFBLFLFBFBFBFFBF
	robot.Reset()
	fmt.Println(robot)
	err := robot.RunCmd(s)
	if err != nil {
		log.Fatal(err)
	}
	x, y := robot.Report()
	fmt.Println("location: ", x, y)
}

const (
	turnLeft  complex128 = 1i
	turnRight complex128 = -1i
)
type Robot struct {
	loc complex128
	dir complex128
}
func (r *Robot) L()     { r.dir *= turnLeft }
func (r *Robot) R()     { r.dir *= turnRight }
func (r *Robot) F()     { r.loc += r.dir }
func (r *Robot) B()     { r.loc -= r.dir }
func (r *Robot) Reset() { r.loc, r.dir = 0, 0+1i }
func (r *Robot) Report() (int, int) {
	return int(real(r.loc)), int(imag(r.loc))
}
type initState struct{}
func (s *initState) Parse(cmd string) (string, string) {
	if len(cmd) == 0 {
		return "", ""
	}
	switch cmd[0] {
	case 'L', 'R', 'F', 'B':
		output, left := s.Parse(cmd[1:])
		return cmd[:1] + output, left
	}
	if cmd[0] >= '1' && cmd[0] < '9' {
		output, left := (&digitState{cmd[:1]}).Parse(cmd[1:])
		if len(output) == 0 {
			return "", cmd
		}
		a, b := s.Parse(left)
		return output + a, b
	}
	return "", cmd
}
type digitState struct {
	n string
}
func (s *digitState) Parse(cmd string) (string, string) {
	if len(cmd) == 0 {
		return "", ""
	}
	if c := cmd[0]; c >= '0' && c < '9' {
		s.n += cmd[:1]
		return s.Parse(cmd[1:])
	} else if c == '(' {
		n, err := strconv.Atoi(s.n)
		if err != nil {
			return "", s.n + cmd
		}
		output, left := (&initState{}).Parse(cmd[1:])
		if len(left) == 0 || left[0] != ')' {
			return "", s.n + cmd
		}
		return strings.Repeat(output, n), left[1:]
	}
	return "", s.n + cmd
}
func ParseCmd(cmd string) (string, error) {
	a, b := (&initState{}).Parse(cmd)
	if len(b) > 0 {
		return "", fmt.Errorf("invalid input: %s", b)
	}
	log.Println(a)
	return a, nil
}
func (r *Robot) RunCmd(cmd string) error {
	instructions, err := ParseCmd(cmd)
	if err != nil {
		return err
	}
	for i := range instructions {
		switch instructions[i] {
		case 'L':
			r.L()
		case 'R':
			r.R()
		case 'F':
			r.F()
		case 'B':
			r.B()
		}
	}
	return nil
}
