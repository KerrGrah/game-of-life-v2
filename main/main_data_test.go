package main

var testBoards = []CellTable{
	{
		0:  []int{1, 2, 3, 4},
		1:  []int{1, 2, 3, 4},
		2:  []int{1, 2},
		10: []int{0, 4},
	},
	{
		0:  []int{1, 3},
		1:  []int{2, 4},
		2:  []int{1, 3},
		10: []int{2, 4},
	},
	{
		0: []int{1, 4, 7, 10},
		1: []int{3, 6, 9, 12},
		2: []int{5, 8, 11, 14},
		3: []int{7, 10, 13, 16},
	},
	{
		0:  []int{6, 7, 8, 9, 12, 13, 14, 16, 17, 19, 20, 21, 24, 25, 27, 31, 32, 35, 37, 39, 40, 43, 46, 47, 49, 51, 53, 55, 56, 58, 59, 64, 70, 74},
		1:  []int{2, 5, 8, 11, 12, 14, 15, 17, 18, 26, 28, 32, 33, 34, 42, 44, 46, 47, 51, 55, 57, 58, 60, 63, 64, 65, 75, 77, 78, 79},
		2:  []int{2, 4, 5, 6, 7, 8, 9, 10, 20, 21, 22, 23, 25, 26, 27, 28, 32, 33, 36, 39, 41, 42, 44, 47, 48, 50, 52, 55, 60, 63, 66, 68, 70, 74},
		3:  []int{0, 2, 7, 8, 10, 15, 18, 19, 22, 29, 32, 36, 38, 41, 42, 43, 45, 46, 48, 49, 53, 54, 55, 57, 61, 63, 64, 65, 66, 67, 68, 71, 73, 75},
		4:  []int{0, 6, 11, 14, 15, 29, 30, 32, 35, 37, 38, 39, 41, 44, 45, 47, 50, 51, 52, 54, 56, 57, 61, 63, 66, 67, 69, 73, 77},
		5:  []int{0, 1, 2, 6, 8, 10, 13, 16, 18, 20, 21, 22, 25, 28, 38, 45, 46, 48, 49, 51, 57, 60, 62, 65, 66, 74, 75, 76},
		6:  []int{0, 2, 3, 6, 7, 8, 10, 11, 12, 18, 19, 20, 24, 27, 28, 29, 30, 31, 32, 34, 35, 39, 42, 44, 46, 47, 50, 51, 52, 54, 55, 58, 59, 60, 62, 63, 65, 66, 67, 68, 69, 70, 73, 75, 76, 79},
		7:  []int{0, 1, 4, 5, 7, 8, 9, 10, 13, 17, 19, 20, 22, 23, 26, 30, 32, 33, 35, 40, 41, 45, 46, 48, 50, 52, 57, 60, 64, 65, 78},
		8:  []int{4, 5, 6, 7, 8, 10, 11, 13, 14, 17, 19, 20, 22, 23, 25, 29, 30, 33, 35, 38, 39, 44, 47, 48, 49, 51, 55, 57, 59, 60, 69, 70, 75, 76, 77},
		9:  []int{1, 2, 3, 7, 8, 17, 23, 25, 27, 32, 37, 40, 43, 46, 48, 49, 50, 52, 53, 54, 58, 59, 61, 67, 68, 71, 73, 75, 78, 79},
		10: []int{3, 4, 11, 13, 15, 16, 17, 18, 24, 26, 31, 34, 35, 39, 41, 46, 47, 50, 51, 55, 61, 62, 65, 66, 72, 74, 79},
		11: []int{0, 1, 3, 10, 12, 13, 25, 26, 32, 33, 34, 35, 38, 39, 41, 43, 44, 45, 46, 49, 55, 57, 60, 63, 69, 70, 72, 73, 75, 77, 78, 79},
		12: []int{2, 3, 4, 5, 10, 16, 17, 18, 19, 22, 23, 26, 28, 30, 31, 33, 34, 36, 42, 44, 47, 50, 51, 52, 54, 55, 56, 61, 64, 65, 69, 70, 73, 75, 78},
		13: []int{2, 6, 10, 12, 13, 15, 17, 21, 22, 23, 24, 27, 28, 30, 34, 38, 47, 48, 49, 50, 57, 62, 70, 71, 73, 74, 75, 76, 79},
		14: []int{0, 5, 7, 10, 11, 15, 20, 21, 22, 23, 25, 29, 33, 34, 39, 41, 42, 44, 46, 47, 48, 51, 53, 60, 62, 63, 66, 67, 69, 70, 75, 76, 77, 78, 79},
		15: []int{1, 2, 3, 7, 13, 14, 19, 21, 23, 32, 33, 34, 37, 39, 40, 46, 51, 52, 56, 57, 62, 64, 66, 67, 68, 74, 77},
		16: []int{4, 5, 13, 14, 22, 25, 26, 30, 32, 36, 38, 41, 44, 48, 49, 52, 54, 55, 57, 61, 62, 64, 65, 66, 68, 69, 70, 71, 73, 77, 78},
		17: []int{3, 6, 7, 9, 10, 15, 18, 22, 23, 26, 27, 28, 29, 30, 34, 35, 39, 42, 46, 47, 48, 49, 52, 56, 57, 61, 63, 69, 73, 74},
		18: []int{3, 11, 13, 14, 21, 22, 23, 24, 32, 34, 35, 36, 37, 44, 45, 47, 49, 53, 56, 57, 62, 65, 67, 72, 76, 77, 79},
		19: []int{2, 4, 7, 12, 17, 18, 19, 21, 25, 28, 29, 30, 31, 35, 36, 37, 38, 39, 40, 41, 44, 45, 48, 49, 50, 54, 56, 57, 58, 60, 61, 62, 63, 64, 67, 68, 69, 70, 71, 72, 73, 76, 77},
		20: []int{1, 4, 5, 8, 9, 13, 14, 15, 16, 17, 18, 19, 24, 25, 27, 29, 30, 31, 32, 36, 40, 42, 43, 44, 45, 51, 52, 55, 57, 60, 61, 62, 66, 67, 68, 69, 72, 75, 76, 78},
		21: []int{6, 8, 10, 11, 12, 14, 16, 18, 20, 23, 24, 27, 30, 31, 34, 35, 39, 41, 42, 43, 45, 46, 47, 52, 55, 56, 60, 61, 68, 69, 70, 71, 72, 74, 75, 77, 79},
		22: []int{1, 4, 6, 7, 9, 12, 17, 18, 19, 20, 22, 28, 31, 38, 42, 45, 48, 49, 59, 62, 68, 69, 76, 77},
		23: []int{0, 1, 13, 16, 17, 19, 23, 26, 27, 28, 30, 31, 32, 33, 35, 44, 45, 46, 48, 51, 54, 55, 56, 58, 61, 64, 67, 68, 69, 71, 72, 74, 77, 78},
		24: []int{0, 1, 2, 3, 4, 6, 7, 8, 12, 15, 16, 17, 18, 22, 24, 25, 27, 31, 40, 41, 42, 47, 49, 51, 55, 56, 57, 58, 60, 61, 64, 65, 66, 68, 69, 73, 75, 76, 77},
		25: []int{4, 6, 7, 9, 10, 21, 23, 26, 27, 30, 31, 33, 34, 35, 37, 38, 40, 41, 42, 43, 48, 54, 55, 56, 58, 59, 63, 64, 69, 71, 74, 75, 76, 78, 79},
		26: []int{5, 7, 8, 10, 15, 18, 25, 31, 34, 37, 38, 39, 42, 43, 45, 47, 49, 51, 54, 56, 57, 59, 61, 65, 67, 71, 73, 75, 77, 78},
		27: []int{3, 7, 9, 11, 13, 14, 16, 17, 18, 22, 24, 25, 31, 32, 33, 38, 39, 42, 45, 46, 50, 53, 54, 57, 59, 60, 64, 65, 66, 67, 70, 72, 74, 75},
		28: []int{1, 4, 6, 9, 17, 18, 22, 29, 30, 33, 35, 39, 41, 46, 47, 48, 50, 51, 53, 55, 56, 59, 62, 64, 65, 67, 69, 70, 71, 72, 74, 78},
		29: []int{1, 3, 7, 9, 11, 12, 13, 15, 18, 21, 22, 28, 30, 32, 34, 37, 43, 45, 47, 48, 52, 54, 59, 61, 62, 64, 65, 66, 68, 72, 73, 77, 78},
		30: []int{0, 1, 6, 7, 8, 16, 18, 19, 21, 23, 25, 26, 30, 33, 34, 36, 37, 40, 42, 44, 47, 48, 49, 50, 52, 57, 58, 59, 64, 65, 66, 67, 68, 69, 72, 75, 79},
		31: []int{1, 2, 5, 11, 12, 15, 17, 25, 27, 28, 35, 36, 38, 41, 43, 47, 48, 50, 52, 54, 55, 56, 57, 58, 61, 62, 63, 64, 67, 69, 72, 75, 76, 79},
		32: []int{0, 3, 5, 6, 8, 10, 11, 19, 21, 22, 23, 31, 33, 34, 35, 40, 41, 44, 45, 48, 51, 56, 57, 60, 63, 72, 75, 77},
		33: []int{1, 2, 5, 7, 8, 11, 12, 16, 21, 22, 23, 24, 26, 28, 29, 34, 37, 41, 43, 47, 48, 53, 55, 60, 63, 64, 65, 66, 73, 74, 76},
		34: []int{0, 4, 5, 6, 8, 10, 12, 15, 17, 18, 23, 25, 29, 31, 33, 36, 37, 38, 39, 40, 42, 43, 47, 49, 51, 53, 54, 59, 60, 61, 66, 70, 76, 77},
		35: []int{0, 2, 4, 5, 6, 8, 9, 10, 11, 12, 17, 23, 25, 29, 31, 32, 35, 37, 38, 39, 43, 49, 52, 53, 54, 55, 58, 59, 60, 62, 63, 72, 75, 77, 78},
		36: []int{0, 1, 3, 4, 6, 7, 8, 13, 14, 15, 19, 26, 32, 33, 38, 39, 40, 41, 44, 45, 50, 53, 63, 64, 68, 69, 70, 72, 73, 75, 77},
		37: []int{2, 3, 7, 9, 13, 15, 16, 17, 19, 23, 24, 26, 27, 30, 34, 36, 37, 38, 39, 40, 42, 43, 44, 46, 47, 52, 55, 56, 59, 61, 62, 63, 64, 66, 70, 73, 76, 78, 79},
		38: []int{0, 1, 2, 4, 5, 8, 10, 13, 15, 20, 22, 26, 30, 31, 37, 41, 44, 46, 48, 53, 55, 57, 59, 60, 62, 63, 71, 73, 76, 77, 78},
		39: []int{0, 3, 4, 10, 15, 16, 29, 32, 35, 36, 42, 44, 45, 47, 50, 53, 55, 56, 58, 59, 62, 63, 68, 70, 75, 76, 79},
		40: []int{4, 9, 10, 11, 12, 18, 19, 22, 23, 25, 27, 30, 31, 35, 37, 39, 40, 45, 46, 49, 54, 56, 58, 59, 62, 66, 68, 74, 75, 78},
		41: []int{4, 5, 7, 13, 14, 21, 22, 25, 31, 32, 34, 45, 46, 49, 50, 51, 54, 56, 57, 61, 62, 63, 64, 65, 66, 68, 73, 75, 76, 77, 78},
		42: []int{4, 8, 17, 18, 19, 21, 22, 23, 25, 27, 28, 31, 32, 34, 37, 38, 39, 41, 42, 43, 44, 46, 47, 50, 52, 55, 56, 60, 62, 63, 66, 67, 73, 75, 77},
		43: []int{6, 10, 11, 13, 14, 15, 18, 22, 30, 31, 32, 33, 34, 39, 40, 43, 45, 48, 49, 51, 52, 55, 56, 60, 63, 65, 68, 74, 76, 77, 79},
		44: []int{1, 2, 3, 4, 7, 8, 11, 12, 13, 16, 17, 18, 22, 23, 24, 26, 28, 29, 31, 37, 40, 42, 45, 46, 47, 48, 49, 54, 56, 60, 62, 64, 65, 69, 70, 76, 77, 78},
		45: []int{0, 5, 10, 12, 13, 17, 19, 20, 28, 29, 30, 31, 34, 35, 37, 39, 40, 47, 48, 51, 55, 56, 57, 58, 59, 60, 62, 64, 65, 66, 70, 72, 76, 77, 79},
		46: []int{2, 3, 6, 8, 14, 15, 16, 25, 30, 31, 33, 34, 37, 39, 41, 43, 45, 51, 52, 53, 54, 56, 66, 68, 70, 74, 75, 76, 79},
		47: []int{2, 6, 7, 10, 14, 17, 18, 19, 23, 24, 25, 28, 30, 35, 36, 37, 39, 40, 44, 48, 49, 52, 54, 55, 59, 60, 63, 64, 67, 68, 73, 76},
		48: []int{10, 12, 13, 16, 19, 27, 29, 31, 36, 37, 45, 47, 51, 54, 57, 58, 63, 68, 70, 72, 75, 78, 79},
		49: []int{3, 4, 5, 7, 8, 14, 19, 20, 24, 28, 31, 32, 33, 36, 40, 41, 42, 46, 53, 58, 59, 64, 66, 67, 74, 76, 77, 78, 79},
		50: []int{1, 7, 10, 12, 14, 18, 19, 20, 23, 25, 29, 30, 35, 41, 43, 44, 48, 49, 51, 52, 56, 57, 58, 60, 65, 67, 69, 73, 76, 77},
	},
	{
		0: []int{1, 4, 7, 10},
		1: []int{3, 6, 9, 12},
		3: []int{7, 10, 13, 16},
	},
}
