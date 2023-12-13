package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "test_1",
			input: []string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			expected: 142,
		},
		{
			name: "test_2",
			input: []string{
				"gtlbhbjgkrb5sixfivefivetwosix",
				"ninesixrgxccvrqscbskgzxh6cpvpxsqnb6",
				"dxxzrlzkksfsffp4",
				"sbzvmddhnjtwollnjv33d2lbcscstqt",
				"88xpnfpb",
				"ninevct4cpdvqfxmspbz9xrvxfvbpzthreesfnncrqn",
				"mqsxzsglbtmsbltrbzkbrt23",
				"seven16ninefc",
				"8jdddctvxs3",
				"fivennhhdfpmrnpjhdm2sixkrsgdt",
				"2bgkccmp2khbhmptwogsz",
				"rrddms473xcjrtsdnhp4",
				"two7eighteight29",
				"onek3three8lpfcfivenxpkhbvdm",
				"spcdsstlkr7fone1",
				"9ninejcngjshghz",
				"six78fivethreesix7fivetwo",
				"nine17bsrfbzpcr",
				"1qcdtzzfdpbcshjsjdmgsix",
				"fourtwosix3one",
			},
			expected: 969,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle1(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "test_1",
			input: []string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			expected: 281,
		},
		{
			name: "test_2",
			input: []string{
				"szsvltgsc1onecccbfour3oneightfh",
			},
			expected: 18,
		},
		{
			name: "test_3",
			input: []string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			expected: 281,
		},
		{
			name: "test_4",
			input: []string{
				"fourfive5jhgpcxmqpr41two",
				"fgjfpxkbszsdkbrzmpronecxzhhcsixonegmkppllnd2",
				"five2853",
				"zszgxvdtq3three",
				"5tkvc5twosixspzhdfourninestpsj",
				"1sevenoneeightlbxskgvpft",
				"ptcfnjgchx1",
				"8mfeightqrqrhvfsb18nnlcc9",
				"nineltnphnvhpvrxbfc1",
				"tcvmnphpmthree56rx",
				"99rsvrmxbgxtjtclnbbmb8sixone3",
				"sevenvfjzctwoqjqmvjrdxpxzsfour5",
				"3d2nhtrhbtfourgsml",
				"9672seven",
				"zttst57zkbjlrmr",
				"455six8fivehqthdfmjfgc",
				"5sevenkgsmclonedvgzqfkgjtwo4",
				"sixsixjkv6nqdhrlkr",
				"qxjbsevenlp56",
				"one13tpbhsixfive8jbt",
				"6five8nnqlcbvmb7",
				"2three1jlghgbzxz751",
				"ninevktrhnf71one4l",
				"13foursixqcrzsrrffknfive",
				"2crb2",
				"1mjmxsvnb3eight",
				"eight517feightmxl8",
				"one1threeeight76rcmkskpvmrz",
				"onethreecqnczs8tdfiveeightthree",
				"fourtwo15nine1",
				"1rxfour4xjzdfgqsixmjvvzfnh6m",
				"zvcfive2eight5hghsix19",
				"3nxlmh448two899",
				"three98oneightzn",
				"3fourtwozszhmcp",
				"twopvhd73",
				"oneqdlsb7sixllszjbceight",
				"xmqxqsixpgclxldnvlzvjm7nine4",
				"fourfive4tttldbmmkxvhqrmvmrkpxfzbd7",
				"44two1",
				"eightrtsjszc2",
			},
			expected: 1913,
		},
		{
			name: "test_5",
			input: []string{
				"twoneight",
			},
			expected: 28,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
