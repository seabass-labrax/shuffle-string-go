// SPDX-FileCopyrightText: 2021 Sebastian Crane <seabass-labrax@gmx.com>
// SPDX-License-Identifier: AGPL-3.0-or-later

package split

import (
	"reflect"
	"testing"
)

type TestSplitStringValues struct {
	inputString string
	expected    []StringFragment
}

type TestIsDelimiterValues struct {
	inputRune rune
	expected  bool
}

func TestIsDelimiter(t *testing.T) {
	var testCases [3]TestIsDelimiterValues

	testCases[0] = TestIsDelimiterValues{inputRune: '\\', expected: true}
	testCases[1] = TestIsDelimiterValues{inputRune: '(', expected: true}
	testCases[2] = TestIsDelimiterValues{inputRune: 'a', expected: false}

	for _, testCase := range testCases {
		actual := IsDelimiter(testCase.inputRune)
		if actual != testCase.expected {
			t.Errorf("Test failed; %+q is a delimeter: expected %t but got %t", testCase.inputRune, testCase.expected, actual)
		}
	}

}
func TestSplitString(t *testing.T) {
	var testCases [2]TestSplitStringValues

	testCases[0] = TestSplitStringValues{
		inputString: "hello, world",
		expected: []StringFragment{
			StringFragment{Value: "hello", Shuffle: true},
			StringFragment{Value: ", ", Shuffle: false},
			StringFragment{Value: "world", Shuffle: true}}}

	testCases[1] = TestSplitStringValues{
		inputString: "What ho!",
		expected: []StringFragment{
			StringFragment{Value: "What", Shuffle: true},
			StringFragment{Value: " ", Shuffle: false},
			StringFragment{Value: "ho", Shuffle: true},
			StringFragment{Value: "!", Shuffle: false}}}

	for _, testCase := range testCases {

		actual := SplitString(testCase.inputString)

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test failed; expected %+v but got %+v", testCase.expected, actual)
		}
	}
}
