package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHasOnlyUniqueChars(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		hasOnlyUniqueChars := hasOnlyUniqueChars("")
		assert.True(t, hasOnlyUniqueChars)
	})

	t.Run("One letter string", func(t *testing.T) {
		hasOnlyUniqueChars := hasOnlyUniqueChars("a")
		assert.True(t, hasOnlyUniqueChars)
	})

	t.Run("Two same letters string", func(t *testing.T) {
		hasOnlyUniqueChars := hasOnlyUniqueChars("dd")
		assert.False(t, hasOnlyUniqueChars)
	})

	t.Run("Upper and lower case letters", func(t *testing.T) {
		hasOnlyUniqueChars := hasOnlyUniqueChars("Dd")
		assert.True(t, hasOnlyUniqueChars)
	})

	t.Run("Longger unique string", func(t *testing.T) {
		hasOnlyUniqueChars := hasOnlyUniqueChars("DawIoqd")
		assert.True(t, hasOnlyUniqueChars)
	})

	t.Run("Longger not unique string", func(t *testing.T) {
		hasOnlyUniqueChars := hasOnlyUniqueChars("DawIoqdD")
		assert.False(t, hasOnlyUniqueChars)
	})
}

func TestReverse(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		reversed := reverse("")
		assert.Equal(t, "", reversed)
	})

	t.Run("Longer string", func(t *testing.T) {
		reversed := reverse("tyDesR")
		assert.Equal(t, "RseDyt", reversed)
	})
}

func TestIsPermutation(t *testing.T) {
	t.Run("Empty strings", func(t *testing.T) {
		isPermutation := isPermutation("", "")
		assert.True(t, isPermutation)
	})

	t.Run("Not permutated int inputs", func(t *testing.T) {
		isPermutation := isPermutationOf([]int{1, 2, 3}, []int{1, 2, 5})
		assert.False(t, isPermutation)
	})

	t.Run("Not permutated double inputs", func(t *testing.T) {
		isPermutation := isPermutationOf([]float32{1.5, 2.5, 3.5}, []float32{1.5, 2.5, 5.5})
		assert.False(t, isPermutation)
	})

	t.Run("Not permutated when some letters are not the same case", func(t *testing.T) {
		isPermutation := isPermutation("Axl Rose", "Oral Sex")
		assert.False(t, isPermutation)
	})

	t.Run("Permutated string inputs", func(t *testing.T) {
		isPermutation := isPermutation(strings.ToLower("Axl Rose"), strings.ToLower("Oral Sex"))
		assert.True(t, isPermutation)
	})

	t.Run("Permutated double inputs", func(t *testing.T) {
		isPermutation := isPermutationOf([]float32{1.5, 2.5, 3.5}, []float32{1.5, 2.5, 3.5})
		assert.True(t, isPermutation)
	})
}

func TestEncodeString(t *testing.T) {
	t.Run("Empty strings", func(t *testing.T) {
		encoded := encodeSpaces("", 0)
		assert.Equal(t, "", encoded)
	})

	t.Run("Negative length", func(t *testing.T) {
		encoded := encodeSpaces("", 0)
		assert.Equal(t, "", encoded)
	})

	t.Run("Only spaces", func(t *testing.T) {
		encoded := encodeSpaces("   ", 0)
		assert.Equal(t, "", encoded)
	})

	t.Run("Longer input", func(t *testing.T) {
		encoded := encodeSpaces("Axl Rose   ", 8)
		assert.Equal(t, "Axl%20Rose", encoded)
	})
}

func TestCompresString(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		compressed := compres("")
		assert.Equal(t, "", compressed)
	})

	t.Run("Should compres string", func(t *testing.T) {
		compressed := compres("aabcccccaaa")
		assert.Equal(t, "a2b1c5a3", compressed)
	})

	t.Run("Should return original string when compressed is longer", func(t *testing.T) {
		compressed := compres("abc")
		assert.Equal(t, "abc", compressed)
	})
}

func TestRotateMatrix90Degrees(t *testing.T) {
	t.Run("Empty matrix", func(t *testing.T) {
		matrix := [][]pixel{}
		rotateMatrix90Degrees(&matrix)
		assert.Equal(t, [][]pixel{}, matrix)
	})

	t.Run("One element matrix", func(t *testing.T) {
		matrix := [][]pixel{
			[]pixel{
				pixel{
					red:   1,
					green: 1,
					blue:  1,
					alpha: 1,
				},
			},
		}
		expectedMatrix := [][]pixel{
			[]pixel{
				pixel{
					red:   1,
					green: 1,
					blue:  1,
					alpha: 1,
				},
			},
		}
		rotateMatrix90Degrees(&matrix)
		assert.Equal(t, expectedMatrix, matrix)
	})

	t.Run("Four elements matrix", func(t *testing.T) {
		matrix := [][]pixel{
			[]pixel{
				pixel{
					red:   1,
					green: 1,
					blue:  1,
					alpha: 1,
				},
				pixel{
					red:   2,
					green: 2,
					blue:  2,
					alpha: 2,
				},
			},
			[]pixel{
				pixel{
					red:   3,
					green: 3,
					blue:  3,
					alpha: 3,
				},
				pixel{
					red:   4,
					green: 4,
					blue:  4,
					alpha: 4,
				},
			},
		}
		rotateMatrix90Degrees(&matrix)

		expectedMatrix := [][]pixel{
			[]pixel{
				pixel{
					red:   3,
					green: 3,
					blue:  3,
					alpha: 3,
				},
				pixel{
					red:   1,
					green: 1,
					blue:  1,
					alpha: 1,
				},
			},
			[]pixel{
				pixel{
					red:   4,
					green: 4,
					blue:  4,
					alpha: 4,
				},
				pixel{
					red:   2,
					green: 2,
					blue:  2,
					alpha: 2,
				},
			},
		}
		assert.Equal(t, expectedMatrix, matrix)
	})

	t.Run("Nine elements matrix", func(t *testing.T) {
		matrix := [][]pixel{
			[]pixel{
				pixel{
					red:   1,
					green: 1,
					blue:  1,
					alpha: 1,
				},
				pixel{
					red:   2,
					green: 2,
					blue:  2,
					alpha: 2,
				},
				pixel{
					red:   3,
					green: 3,
					blue:  3,
					alpha: 3,
				},
			},
			[]pixel{
				pixel{
					red:   4,
					green: 4,
					blue:  4,
					alpha: 4,
				},
				pixel{
					red:   5,
					green: 5,
					blue:  5,
					alpha: 5,
				},
				pixel{
					red:   6,
					green: 6,
					blue:  6,
					alpha: 6,
				},
			},
			[]pixel{
				pixel{
					red:   7,
					green: 7,
					blue:  7,
					alpha: 7,
				},
				pixel{
					red:   8,
					green: 8,
					blue:  8,
					alpha: 8,
				},
				pixel{
					red:   9,
					green: 9,
					blue:  9,
					alpha: 9,
				},
			},
		}
		rotateMatrix90Degrees(&matrix)

		expectedMatrix := [][]pixel{
			[]pixel{
				pixel{
					red:   7,
					green: 7,
					blue:  7,
					alpha: 7,
				},
				pixel{
					red:   4,
					green: 4,
					blue:  4,
					alpha: 4,
				},
				pixel{
					red:   1,
					green: 1,
					blue:  1,
					alpha: 1,
				},
			},
			[]pixel{
				pixel{
					red:   8,
					green: 8,
					blue:  8,
					alpha: 8,
				},
				pixel{
					red:   5,
					green: 5,
					blue:  5,
					alpha: 5,
				},
				pixel{
					red:   2,
					green: 2,
					blue:  2,
					alpha: 2,
				},
			},
			[]pixel{
				pixel{
					red:   9,
					green: 9,
					blue:  9,
					alpha: 9,
				},
				pixel{
					red:   6,
					green: 6,
					blue:  6,
					alpha: 6,
				},
				pixel{
					red:   3,
					green: 3,
					blue:  3,
					alpha: 3,
				},
			},
		}

		assert.Equal(t, expectedMatrix, matrix)
	})
}

func TestZeroColumnsAndRows(t *testing.T) {
	t.Run("Empty matrix", func(t *testing.T) {
		matrix := [][]int{}
		zeroColumnsAndRows(&matrix)
		assert.Equal(t, [][]int{}, matrix)
	})

	t.Run("Matrix without zeros", func(t *testing.T) {
		matrix := [][]int{
			[]int{2, 4, 6},
			[]int{1, 2, 3},
			[]int{7, 9, 8},
		}
		expectedMatrix := [][]int{
			[]int{2, 4, 6},
			[]int{1, 2, 3},
			[]int{7, 9, 8},
		}

		assert.Equal(t, expectedMatrix, matrix)
	})

	t.Run("Should insert zeroes in small matrix", func(t *testing.T) {
		matrix := [][]int{
			[]int{0, 4},
			[]int{1, 2},
		}
		expectedMatrix := [][]int{
			[]int{0, 0},
			[]int{0, 2},
		}
		zeroColumnsAndRows(&matrix)

		assert.Equal(t, expectedMatrix, matrix)
	})

	t.Run("Should insert zeroes in big matrix", func(t *testing.T) {
		matrix := [][]int{
			[]int{0, 4, 6, 6, 7},
			[]int{1, 2, 3, 4, 6},
			[]int{7, 0, 4, 9, 8},
			[]int{1, 2, 5, 4, 6},
			[]int{7, 2, 3, 4, 6},
		}
		expectedMatrix := [][]int{
			[]int{0, 0, 0, 0, 0},
			[]int{0, 0, 3, 4, 6},
			[]int{0, 0, 0, 0, 0},
			[]int{0, 0, 5, 4, 6},
			[]int{0, 0, 3, 4, 6},
		}
		zeroColumnsAndRows(&matrix)

		assert.Equal(t, expectedMatrix, matrix)
	})
}

func TestIsRotated(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		isRotated := isRotated("", "")
		assert.True(t, isRotated)
	})

	t.Run("Not rotated  string", func(t *testing.T) {
		isRotated := isRotated("David", "Bowie")
		assert.False(t, isRotated)
	})

	t.Run("Rotated string", func(t *testing.T) {
		isRotated := isRotated("AxlRose", "lRoseAx")
		assert.True(t, isRotated)
	})

	t.Run("Parial string", func(t *testing.T) {
		isRotated := isRotated("David", "avi")
		assert.False(t, isRotated)
	})

}
