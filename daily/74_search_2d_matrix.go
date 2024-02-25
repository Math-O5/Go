
/*
   You are given an m x n integer matrix matrix with the following two properties:

   Each row is sorted in non-decreasing order.
   The first integer of each row is greater than the last integer of the previous row.

   Given an integer target, return true if target is in matrix or false otherwise.

   You must write a solution in O(log(m * n)) time complexity.

   Testing solution:

   target: 14
   n = 4 i = mid (0,4) = 2
   m = 2 j = mid (0,2) = 1

    1  2  3  4  5
    6  7  8  9  10
    11 12 13 14 15

    n = 4 n0 = 3 i = 3
    m = 2 m0 = 2 j = 2

    n = 4 n0 = 3 i = 3
    m = 2 m0 = 1 j = 1

*/

import (
	"io"
	"fmt"
	"strconv"
)

func binary_search(matrix [][]int, target int) bool {
	mid := 0
	init := 0
	n, m := len(matrix[0])-1, len(matrix)-1

	for init <= m {
		mid = (init + m) / 2

		if target >= matrix[mid][0] && target <= matrix[mid][n] {
			// do binary search
			init = 0
			m = mid

			for init <= n {
				mid = (init + n) / 2

				if matrix[m][mid] == target {
					return true
				}

				if matrix[m][mid] < target {
					init = mid + 1
				} else {
					n = mid - 1
				}
			}

			break
		}

		if matrix[mid][0] < target {
			init = mid + 1
		} else {
			m = mid - 1
		}
	}

	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	return binary_search(matrix, target)
}

