package main

import (
"sort"
"fmt"
)

type people []string

func (p people) Len() int { return len(p) }
func (p people) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j]}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	fmt.Println("---------------------")

	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

	fmt.Println("---------------------")

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)
}
