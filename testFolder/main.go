package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, a, b int
	fmt.Fscan(reader, &n, &m)
	graph := make([][]int, n+1)
	visited := make([]bool, n+1)

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	ans := make([][]int, 0)
	for {
		tmpPoint := isAllVisited(visited)
		if tmpPoint == -1 {
			break
		}
		tmpComp := dfs(graph, tmpPoint, visited, nil)
		ans = append(ans, tmpComp)
	}

	fmt.Println(len(ans))
	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans[i]); j++ {
			fmt.Print(ans[i][j], " ")
		}
	}
}

func dfs(graph [][]int, startPoint int, isVisitedList []bool, curComp []int) []int{
	isVisitedList[startPoint] = true
	if curComp == nil {
		curComp = make([]int, 0)
	}
	curComp = append(curComp, startPoint)
	for _, point := range graph[startPoint] {
		if !isVisitedList[point] {
			dfs(graph, point, isVisitedList, curComp)
		}
	}
	return curComp
}

func isAllVisited(in []bool) int {
	for i := 1; i < len(in); i++ {
		if !in[i] {
			return i
		}
	}
	return -1
}