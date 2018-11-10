package main

import (
	"fmt"
)

/*
	パイプラインはデータを受け取って、何らかに処理のを行い、どこかに渡すという一連の作業。＝ ステージ
	* ステージは受け取るものと返すものが同じ型
	* ステージは引き回せるように具体化されていなければならない。
*/

func main() {

	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}
}
