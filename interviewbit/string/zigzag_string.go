package string

import "bytes"

func ConvertZigZag(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows), 0, numRows-2

	for i := 0; i != len(s); {
		if down < numRows {
			matrix[down] = append(matrix[down], s[i])
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], s[i])
			up--
			i++
		} else {
			down = 0
			up = numRows - 2
			matrix[down] = append(matrix[down], s[i])
			down++
			i++
		}
	}

	var ans bytes.Buffer
	for _, row := range matrix {
		ans.WriteString(string(row))
	}

	return ans.String()
}
