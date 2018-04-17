package _go

/*
题目描述
在一个二维数组中，每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
 */

// 有序没有应用
func find(target int, array *[][]int) bool {
	for _, x := range *array {
		for _, y := range x {
			if y == target {
				return true
			}
		}
	}
	return false
}
