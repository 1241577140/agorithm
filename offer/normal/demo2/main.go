package main

import "fmt"

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。



*/
func main() {
	//fmt.Println(findNumberIn2DArray([][]int{
	//	{1, 4, 7, 11, 15},
	//	{2, 5, 8, 12, 19},
	//	{3, 6, 9, 16, 22},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30},
	//}, 4))
	//fmt.Println(findNumberIn2DArray([][]int{{5}, {6}}, 6))
	fmt.Println(findNumberIn2DArray([][]int{{1, 4}, {2, 4}}, 2))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	col := len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {
				col = j
				break
			}
		}
	}
	return false
}
