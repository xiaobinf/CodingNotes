package atoi_8

import (
	"fmt"
	"testing"
	"unicode"
)

const INT_MAX = 1<<31 - 1
const INT_MIN = -1 << 31

type Automation struct {
	state string
	sign  int
	ans   int
	table map[string][]string
}

func NewAutomation() *Automation {
	return &Automation{
		state: "start",
		sign:  1,
		ans:   0,
		table: map[string][]string{
			"start":     {"start", "signed", "in_number", "end"},
			"signed":    {"end", "end", "in_number", "end"},
			"in_number": {"end", "end", "in_number", "end"},
			"end":       {"end", "end", "end", "end"},
		},
	}
}

func (a *Automation) getCol(c byte) int {
	if unicode.IsSpace(rune(c)) {
		return 0
	}
	if c == '+' || c == '-' {
		return 1
	}
	if unicode.IsDigit(rune(c)) {
		return 2
	}
	return 3
}

func (a *Automation) updateAutomation(c byte) {
	a.state = a.table[a.state][a.getCol(c)]
	// 判断溢出
	if a.state == "in_number" {
		a.ans = a.ans*10 + int(c-'0')
		if a.sign == 1 {
			a.ans = min(a.ans, INT_MAX)
		} else {
			a.ans = min(a.ans, -INT_MIN)
		}
	}
	if a.state == "signed" {
		if c == '+' {
			a.sign = 1
		} else {
			a.sign = -1
		}
	}
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func myAtoi(s string) int {
	a := NewAutomation()
	for i := 0; i < len(s); i++ {
		a.updateAutomation(s[i])
	}
	return a.sign * a.ans
}

func TestName(t *testing.T) {
	fmt.Println(myAtoi("2147483649"))
}
