// generated by Textmapper; DO NOT EDIT

package tm

const tmNumClasses = 36

var tmRuneClass = []uint8{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 35, 4, 1, 1, 12, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	35, 17, 9, 13, 31, 11, 30, 2, 22, 23, 8, 28, 21, 10, 20, 5,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 14, 19, 27, 15, 18, 29,
	32, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 6, 3, 7, 1, 33,
	1, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 24, 16, 26, 25,
}

const tmRuneClassLen = 127

var tmStateMap = []int{
	0, 0, 60,
}

var tmToken = []Token{
	39, 1, 2, 3, 0, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78,
}

var tmLexerAction = []int8{
	-2, -1, 57, -1, 56, 44, 43, 42, 41, 38, 37, 34, 56, 32, 29, 26,
	24, 22, 21, 20, 19, 18, 17, 16, 14, 13, 12, 11, 9, 8, 6, 5,
	4, 2, 1, 56, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, 1, -6, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, 3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, 2, 2, -3, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, 3, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 2, 2, -1,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -38, -38, -38, -38, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, 7, -38, -38, -38, -38, -38, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35,
	-35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35,
	-35, -35, -35, -35, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, 10, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -79, -79, -79, -79, -79, -79, -79, -79,
	-79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79,
	-79, 15, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -22, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -36, -36, -36, -36, -36, -36, -36, -36,
	-36, -36, -36, -36, -36, -36, -36, 23, -36, -36, -36, -36, -36, -36, -36, -36,
	-36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -18, -18, -18, -18,
	-18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
	25, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, 28, -16, -16, 27, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -19, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, 30, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, 31, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-9, 32, 32, 32, 33, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 32, 32, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, 35, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -7, 35, 35, 35,
	36, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, 1, -1, -1, 38, 38, 40, -1, 38, 38, 38,
	38, 39, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-1, 38, 38, 38, -1, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-1, 51, 51, 50, -1, -1, 48, 51, 45, 51, 51, 51, -1, 51, 51, 51,
	51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51,
	51, 51, 51, 51, -1, 45, 45, 45, 45, 45, 45, 45, 46, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, -1, 45, 45, 45, 45, 47, 45, 45,
	46, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-1, 48, 48, 49, -1, 48, 48, 51, 48, 48, 48, 48, -1, 48, 48, 48,
	48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
	48, 48, 48, 48, -1, 48, 48, 48, -1, 48, 48, 48, 48, 48, 48, 48,
	48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
	48, 48, 48, 48, 48, 48, 48, 48, -1, 51, 51, 51, -1, 51, 51, 51,
	51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51,
	51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, -1, 51, 51, 55,
	-1, 54, 52, 51, 51, 51, 51, 51, -1, 51, 51, 51, 51, 51, 51, 51,
	51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51,
	-1, 52, 52, 53, -1, 52, 52, 51, 52, 52, 52, 52, -1, 52, 52, 52,
	52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52,
	52, 52, 52, 52, -1, 52, 52, 52, -1, 52, 52, 52, 52, 52, 52, 52,
	52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52,
	52, 52, 52, 52, 52, 52, 52, 52, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -1, 51, 51, 51,
	-1, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51,
	51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51,
	-8, -8, -8, -8, 56, -8, -8, -8, -8, -8, -8, -8, 56, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, 56, -1, 57, 59, 58, -1, 57, 57, 57, 57, 57, 57, 57,
	57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57,
	57, 57, 57, 57, 57, 57, 57, 57, -1, 57, 57, 57, -1, 57, 57, 57,
	57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57,
	57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-1, -1, 57, -1, 56, 44, 43, 42, 41, 38, 37, 34, 56, 32, 29, 26,
	24, 22, 21, 20, 19, 18, 17, 16, 61, 13, 12, 11, 9, 8, 6, 5,
	4, 2, 1, 56, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80,
	-80, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80, 15, -80, -80,
	-80, -80, -80, -80, -80, -80, -80, -80,
}
