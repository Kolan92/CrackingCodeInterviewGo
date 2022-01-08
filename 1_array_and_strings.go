package main

import (
	"constraints"
	"errors"
	"fmt"
	"strings"
)

func hasOnlyUniqueChars(input string) bool {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}
	return true
}

func reverse(input string) string {
	inputLength := len(input)
	if inputLength == 0 {
		return input
	}

	reversedArray := make([]rune, inputLength)
	for index, char := range input {
		reversedArray[inputLength-index-1] = char

	}

	return string(reversedArray)
}

func isPermutation(input, other string) bool {
	return isPermutationOf[rune]([]rune(input), []rune(other))
}

func isPermutationOf[K constraints.Ordered](input, other []K) bool {
	existingChars := map[K]bool{}
	for _, char := range input {
		existingChars[char] = true
	}

	for _, char := range other {
		if _, hasKey := existingChars[char]; !hasKey {
			return false
		}
	}

	return true
}

func encodeSpaces(input string, length int) string {
	if length < 0 {
		length = 0
	}

	runes := []rune(input)[:length]
	output := make([]string, length)

	for index, char := range runes {
		if char == ' ' {
			output[index] = "%20"
		} else {
			output[index] = string(char)
		}
	}

	return strings.Join(output, "")
}

type compressEntry struct {
	count int
	char  rune
}

func (entry compressEntry) toString() string {
	return fmt.Sprintf("%s%d", string(entry.char), entry.count)
}

func compres(input string) string {
	inputLength := len(input)
	if inputLength == 0 {
		return input
	}

	entries := []compressEntry{}
	runes := []rune(input)
	currentEntry := compressEntry{
		char:  runes[0],
		count: 0,
	}

	for _, char := range runes {
		if currentEntry.char == char {
			currentEntry.count++
		} else {
			entries = append(entries, currentEntry)
			currentEntry = compressEntry{
				char:  char,
				count: 1,
			}
		}
	}
	entries = append(entries, currentEntry)

	convertedEntries, _ := Map(entries, func(entry compressEntry) string {
		return entry.toString()
	})

	compressedString := strings.Join(convertedEntries, "")
	if inputLength > len(compressedString) {
		return compressedString
	} else {
		return input
	}
}

func Map[KInput any, KOutput any](input []KInput, transform func(KInput) KOutput) ([]KOutput, error) {
	if input == nil {
		return []KOutput{}, errors.New("Nil input")
	}

	if transform == nil {
		return []KOutput{}, errors.New("Nil transform")
	}
	output := make([]KOutput, len(input))

	for index, value := range input {
		output[index] = transform(value)
	}

	return output, nil
}

type pixel struct {
	red, green, blue, alpha byte
}

func rotateMatrix90Degrees(matrix *[][]pixel) {
	size := len(*matrix)
	halfSize := size / 2
	for layer := 0; layer < halfSize; layer++ {
		first := layer
		last := size - 1 - layer
		for index := first; index < last; index++ {
			elementIndex := last - index - first
			topElement := (*matrix)[first][index]

			//left -> top
			(*matrix)[first][index] = (*matrix)[elementIndex][first]
			//bottom -> left
			(*matrix)[elementIndex][first] = (*matrix)[last][elementIndex]
			//right -> bottom
			(*matrix)[last][elementIndex] = (*matrix)[index][last]
			//top -> right
			(*matrix)[index][last] = topElement
		}
	}
}

func zeroColumnsAndRows(matrix *[][]int) {
	zeroRows := make(map[int]bool)
	zeroColumns := make(map[int]bool)

	size := len(*matrix)

	for row := 0; row < size; row++ {
		for column := 0; column < size; column++ {
			if (*matrix)[row][column] == 0 {
				zeroRows[row] = true
				zeroColumns[column] = true
			}
		}
	}

	for row := 0; row < size; row++ {
		_, isZeroRow := zeroRows[row]

		for column := 0; column < size; column++ {
			_, isZeroColumn := zeroColumns[column]
			if isZeroRow || isZeroColumn {
				(*matrix)[row][column] = 0
			}
		}
	}
}

func isRotated(source, candidate string) bool {
	if len(source) != len(candidate) {
		return false
	}
	return strings.Contains(source+source, candidate)
}
