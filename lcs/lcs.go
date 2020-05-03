package main

import (
	"fmt"
	"math"
)

func lcsRec(a, b string) float64 {
	aByte := []byte(a)
	bByte := []byte(b)

	if len(a) == 0 || len(b) == 0 {
		return 0
	} else if aByte[len(a)-1] == bByte[len(b)-1] {
		return 1 + lcsRec(string(aByte[0:len(a)-1]), string(bByte[0:len(b)-1]))
	} else {
		return math.Max(float64(lcsRec(string(aByte[0:len(a)-1]), string(bByte))), float64(lcsRec(string(aByte), string(bByte[0:len(b)-1]))))
	}
}

func lcsDP(a, b string) int {
	aByte := []byte(a)
	bByte := []byte(b)

	var arr [100][100]int
	arrSlice := arr[0 : len(a)+1][0 : len(b)+1]

	for i, aChar := range aByte {
		for j, bChar := range bByte {
			if bChar == aChar {
				arrSlice[i+1][j+1] = arrSlice[i][j] + 1
			} else if bChar != aChar {
				arrSlice[i+1][j+1] = int(math.Max(float64(arr[i][j+1]), float64(arr[i+1][j])))
			}
		}
	}

	return arrSlice[len(a)][len(b)]
}

func main() {
	a := "AGGTAB"
	b := "GXTXAYB"
	fmt.Println(lcsRec(a, b)) // Prints 4 (GTAB)
	fmt.Println(lcsDP(a, b))  // Prints 4
}
