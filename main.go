package main

import (
	"bufio"
	"fmt"
	"kript/iterator"
	"kript/rpm"
	"os"
	"strconv"
	"strings"
)

func StringPermutations(pool []string, r int) [][]string {
	n := len(pool)
	s := 1
	ss := 1 // s / ss == nPr
	// if r > n
	// err
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	cycles := make([]int, r)
	for i := 0; i < r; i++ {
		cycles[i] = n - i
		s *= (n - i)
		ss *= (i + 1)
	}
	rst := make([][]string, 0, s/ss)
	rst = append(rst, make([]string, r))
	for i, v := range indices[:r] {
		rst[0][i] = pool[v]
	}

	for n > 0 {
		f := true
		for i := r - 1; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				tmp := append(indices[i+1:], indices[i:i+1]...)
				for j := 0; j < len(tmp); j++ {
					indices[i+j] = tmp[j]
				}
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[len(indices)-j] = indices[len(indices)-j], indices[i]
				tmp := make([]string, 0, r)
				for _, v := range indices[:r] {
					tmp = append(tmp, pool[v])
				}
				rst = append(rst, tmp)
				f = false
				break
			}
		}
		if f {
			return rst
		}
	}
	return rst
}

func StringProducts(s []string) [][]string {
	n := len(s)
	rst := make([][]string, 0, n*n*n*n)
	for _, fourth := range s {
		for _, third := range s {
			for _, second := range s {
				for _, first := range s {
					rst = append(rst, []string{first, second, third, fourth})
				}
			}
		}
	}
	return rst
}

func Input(prompt string) (stringInput string) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rawInputed := scanner.Text()
	return strings.TrimSpace(rawInputed)
}

func main() {
	// set := set.NewSet()
	// fmt.Println(set)
	// StringPermutations([]string{"a", "b", "c"}, 2)
	cards := make([]string, 0, 30)
	for i := 1; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			cards = append(cards, fmt.Sprintf("%d", i))
		}
	}
	operators := []string{"+", "-", "*", "/"}
	cardComb := iterator.NewStringPermutationGenerator(cards, 6)
	for cardComb.Next() {
		c := cardComb.GetCurrent()
		operatorComb := iterator.NewStringProductGenerator(operators, 4)
		for operatorComb.Next() {
			o := operatorComb.GetCurrent()
			expr := fmt.Sprintf(
				"%s %s %s %s %s %s %s %s %s",
				c[0], c[1], c[2], c[3], c[4],
				o[0], o[1], o[2], o[3],
			)
			rst := rpm.Rpm(expr)
			expected, _ := strconv.Atoi(fmt.Sprintf("%s", c[5]))
			actual := int(rst)
            if rst - float32(actual) == 0 && expected == actual {
				fmt.Printf("%s == %d\n", expr, actual)
			}
		}
	}
}
