// generated by Textmapper; DO NOT EDIT

package tm

const tmNumClasses = 36

var tmRuneClass = []uint8{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 1, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 21, 22, 23, 24, 25, 26, 27, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 29, 30, 31, 1, 28, 1, 28, 28, 28, 28, 28, 28, 28, 28, 28,
	28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 32, 33,
	34, 35,
}

const tmRuneClassLen = 127
const tmFirstRule = -4

var tmStateMap = []int{
	0, 0, 51, 61,
}

var tmLexerAction = []int8{
	-4, -5, 50, 50, 50, 48, 45, 44, 43, 40, 38, 35, 32, 31, 30, 28, 27, 25, 24,
	20, 19, 17, 16, 15, 13, 12, 11, 10, 8, 7, -5, 6, 5, 3, 2, 1, -36, -36, -36,
	-36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36,
	-36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36,
	-36, -36, -36, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, 4, -13, -13, -14,
	-14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -14, -81, -81, -81, -81, -81, -81, -81, -81, -81, -81,
	-81, -81, -81, -81, -81, -81, -81, -81, -81, -81, -81, -81, -81, -81, -81,
	-81, -81, -81, -81, -81, -81, -81, -81, -81, -81, -81, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -42, -42, -42, -42, -42, -42, -42,
	-42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -1, -42, -42, 8, -42, -42,
	-42, -42, -42, -42, -42, 8, -42, -42, -42, -42, -42, -42, -42, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 9, -5, -5, 8, -5, -5,
	-5, -5, -5, -5, -5, 8, -5, -5, -5, -5, -5, -5, -5, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, 14, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-18, -18, -18, -18, -18, -18, -18, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, 18, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, 19, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84, -84,
	-84, -2, -84, -84, -84, -84, 44, -84, -84, -84, -84, -84, -84, -84, -84, -84,
	-84, -84, -84, -84, -84, -84, -84, -5, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 22, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, -5, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 22, 21, 21, 21, 21, 23, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -19, -19, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, 19, -5, -5, -5, -5, 26, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, 29, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -3, -24, -24, -24, -24, -24, -24, -24, -24, -24, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 34,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25,
	-5, 35, 35, -5, 35, 35, 35, 35, 35, 35, 35, 37, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 36, 35, 35, 35, 35, 35, -5, 35,
	35, -5, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, 35, -42, -42, -42,
	-42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42,
	-42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42,
	-42, -42, -42, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, 39, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -38, -38, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, 41, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -4, 41, 41, 42, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -9, 44, 44, -9, -9, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, -5, 45, 45, -5, 45, 45, 47, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	46, 45, 45, 45, 45, 45, -5, 45, 45, -5, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35,
	-35, -35, -35, -35, -35, -35, -35, -35, -35, -35, 49, -35, -35, -35, -35,
	-35, -35, -35, -35, -35, -35, -35, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -8, -8, 50,
	50, 50, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -5, -5, 50, 50, 50,
	48, 45, 44, 43, 40, 38, 35, 32, 31, 30, 28, 27, 25, 24, 52, 19, 17, 16, 15,
	13, 12, 11, 10, 8, 7, -5, 6, 5, 3, 2, 1, -5, 56, 56, -5, -5, 56, 56, 56, 56,
	56, 56, 56, 56, 56, 21, 56, 56, 56, 56, 44, 56, 56, 56, 56, 56, 56, 56, 56,
	56, 54, 53, 56, 56, 56, 56, 56, -5, 56, 56, -5, 56, 56, 56, 56, 56, 56, 56,
	56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56,
	56, 56, 56, 56, 56, 56, -5, 54, 54, -5, -5, 54, 54, 54, 54, 54, 54, 54, 54,
	54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 55, 56,
	54, 54, 54, 54, -5, 54, 54, -5, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54,
	54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54, 54,
	54, 54, -5, 56, 56, -5, -5, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56,
	56, 56, 60, 56, 56, 56, 56, 56, 56, 56, 56, 56, 58, 57, 56, 56, 56, 56, 56,
	-5, 56, 56, -5, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56,
	56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, -5, 58,
	58, -5, -5, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58,
	58, 58, 58, 58, 58, 58, 58, 58, 58, 59, 56, 58, 58, 58, 58, -5, 58, 58, -5,
	58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58,
	58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, -83, -83, -83, -83, -83,
	-83, -83, -83, -83, -83, -83, -83, -83, -83, -83, -83, -83, -83, -83, -83,
	-83, -83, -83, -83, -83, -83, -83, -83, -83, -83, -83, -83, -83, -83, -83,
	-83, -5, -5, 50, 50, 50, 48, 45, 44, 43, 40, 38, 35, 32, 31, 30, 28, 27, 25,
	24, 20, 19, 17, 16, 15, 13, 12, 11, 10, 8, 7, -5, 6, 62, 3, 2, 1, -82, -82,
	-82, -82, -82, -82, -82, -82, -82, -82, -82, -82, -82, -82, -82, -82, -82,
	-82, -82, -82, -82, -82, -82, -82, -82, -82, -82, -82, -82, -82, -82, -82,
	-82, -82, -82, -82,
}

var tmBacktracking = []int{
	38, 9, // in ID
	80, 21, // in '/'
	20, 33, // in '('
}
