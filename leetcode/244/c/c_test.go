// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`"111000"`, 
			`2`,
		},
		{
			`"010"`, 
			`0`,
		},
		{
			`"1110"`, 
			`1`,
		},
		// TODO 测试入参最小的情况
		{
			`"01001001101"`,
			`2`,
		},
		{
			`"10001100101000000"`,
			`5`,
		}, // 00110010100000010
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minFlips, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-244/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/
