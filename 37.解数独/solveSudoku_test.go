package solveSudoku

import (
	"fmt"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	cases := []struct{
		board 	[][]byte
	} {
		{
			[][]byte{
				[]byte{'5','3','.','.','7','.','.','.','.'},
				[]byte{'6','.','.','1','9','5','.','.','.'},
				[]byte{'.','9','8','.','.','.','.','6','.'},
				[]byte{'8','.','.','.','6','.','.','.','3'},
				[]byte{'4','.','.','8','.','3','.','.','1'},
				[]byte{'7','.','.','.','2','.','.','.','6'},
				[]byte{'.','6','.','.','.','.','2','8','.'},
				[]byte{'.','.','.','4','1','9','.','.','5'},
				[]byte{'.','.','.','.','8','.','.','7','9'},
			},
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			solveSudoku(v.board)
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					fmt.Printf("%c ", v.board[i][j])
				}
				fmt.Println()
			}
		})
	}
}