package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name   string
		num1   string
		num2   string
		result string
	}{
		{
			name:   "example 1",
			num1:   "[1,2]",
			num2:   "[[3,4],5]",
			result: "[[1,2],[[3,4],5]]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			n := SnailfishNumber{value: tc.num1}
			n.add(tc.num2)
			require.Equal(t, tc.result, n.String())
		})
	}
}

func TestAddAndReduce(t *testing.T) {
	tests := []struct {
		name   string
		num1   string
		num2   string
		result string
	}{
		{
			name:   "example 1",
			num1:   "[[[[4,3],4],4],[7,[[8,4],9]]]",
			num2:   "[1,1]",
			result: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			name:   "example 2",
			num1:   "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			num2:   "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			result: "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
		{
			name:   "example 3",
			num1:   "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
			num2:   "[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
			result: "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		},
		{
			name:   "example 4",
			num1:   "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
			num2:   "[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
			result: "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
		},
		{
			name:   "example 5",
			num1:   "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
			num2:   "[7,[5,[[3,8],[1,4]]]]",
			result: "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
		},
		{
			name:   "example 6",
			num1:   "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
			num2:   "[[2,[2,2]],[8,[8,1]]]",
			result: "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
		},
		{
			name:   "example 7",
			num1:   "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
			num2:   "[2,9]",
			result: "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
		},
		{
			name:   "example 8",
			num1:   "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
			num2:   "[1,[[[9,3],9],[[9,0],[0,7]]]]",
			result: "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
		},
		{
			name:   "example 9",
			num1:   "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
			num2:   "[[[5,[7,4]],7],1]",
			result: "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
		},
		{
			name:   "example 10",
			num1:   "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
			num2:   "[[[[4,2],2],6],[8,7]]",
			result: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			n := SnailfishNumber{value: tc.num1}
			n.add(tc.num2)
			n.reduce()
			require.Equal(t, tc.result, n.String())
		})
	}
}

func TestExplode(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result string
	}{
		{
			name:   "example 1",
			input:  "[[[[[9,8],1],2],3],4]",
			result: "[[[[0,9],2],3],4]",
		},
		{
			name:   "example 2",
			input:  "[7,[6,[5,[4,[3,2]]]]]",
			result: "[7,[6,[5,[7,0]]]]",
		},
		{
			name:   "example 3",
			input:  "[[6,[5,[4,[3,2]]]],1]",
			result: "[[6,[5,[7,0]]],3]",
		},
		{
			name:   "example 4",
			input:  "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			result: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			name:   "example 5",
			input:  "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			result: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			name:   "example 6",
			input:  "[[[[4,0],[5,4]],[[7,7],[0,[6,7]]]],[10,[[11,9],[11,0]]]]",
			result: "[[[[4,0],[5,4]],[[7,7],[6,0]]],[17,[[11,9],[11,0]]]]",
		},
		{
			name:   "example 7",
			input:  "[[[[12,12],[6,14]],[[15,0],[17,[8,1]]]],[2,9]]",
			result: "[[[[12,12],[6,14]],[[15,0],[25,0]]],[3,9]]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			n := SnailfishNumber{value: tc.input}
			n.explode()
			require.Equal(t, tc.result, n.String())
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result string
	}{
		{
			name:   "example 1",
			input:  "[[[[0,7],4],[15,[0,13]]],[1,1]]",
			result: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		},
		{
			name:   "example 2",
			input:  "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			result: "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			n := SnailfishNumber{value: tc.input}
			n.split()
			require.Equal(t, tc.result, n.String())
		})
	}
}

func TestGetMagnitude(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result int
	}{
		{
			name:   "example -1",
			input:  "[9,1]",
			result: 29,
		},
		{
			name:   "example 0",
			input:  "[[9,1],[1,9]]",
			result: 129,
		},
		{
			name:   "example 1",
			input:  "[[1,2],[[3,4],5]]",
			result: 143,
		},
		{
			name:   "example 2",
			input:  "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			result: 1384,
		},
		{
			name:   "example 3",
			input:  "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			result: 445,
		},
		{
			name:   "example 4",
			input:  "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			result: 791,
		},
		{
			name:   "example 5",
			input:  "[[[[5,0],[7,4]],[5,5]],[6,6]]",
			result: 1137,
		},
		{
			name:   "example 6",
			input:  "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			result: 3488,
		},
		{
			name:   "example 7",
			input:  "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
			result: 4140,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			n := SnailfishNumber{value: tc.input}
			result := n.getMagnitude()
			require.Equal(t, tc.result, result)
		})
	}
}
