package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Add(1, 2) returned %d, expected 3.", result)
	}
}

type divideTestCase struct {
	a, b           int
	expectedResult int
	expectErr      bool
}

func TestSubtract(t *testing.T) {
	result := Subtract(2, 1)
	expected := 1
	if result != expected {
		t.Errorf("Subtract(2, 1) returned %d, expected 1.", result)
	}
}

func TestDivide(t *testing.T) {

	result, err := Divide(6, 3)

	assert.NoError(t, err, "an unexpected error occurred: %v", err)
	assert.Equal(t, 2, result, "they should be equal")

	// if err != nil {
	// 	t.Errorf("Divide(6, 3) failed. Expected no error, got %v", err.Error())
	// }
	// expected := 2
	// if result != expected {
	// 	t.Errorf("Divide(6, 3) returned %d, expected 2.", result)
	// }

}

// テーブルドリブンテスト 標準パッケージver
func TestDivide2(t *testing.T) {

	testCases := []divideTestCase{
		{a: 6, b: 3, expectedResult: 2, expectErr: false},
		{a: 6, b: 0, expectedResult: 0, expectErr: true},
	}

	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)
		if tc.expectErr {
			if err == nil {
				t.Errorf("Divide(%d, %d) failed. Expected an error, got nil.", tc.a, tc.b)
			}
		} else {
			if err != nil {
				t.Errorf("Divide(%d, %d) failed. Expected no error, got %v", tc.a, tc.b, err.Error())
			}
		}
		if result != tc.expectedResult {
			t.Errorf("Divide(%d, %d) returned %d, expected %d.", tc.a, tc.b, result, tc.expectedResult)
		}
	}

}

// / テーブルドリブンテスト assert関数ver
func TestDivide3(t *testing.T) {

	// testCases := []struct {
	// 	a, b           int
	// 	expectedResult int
	// 	expectErr      bool
	// }{
	// 	{a: 6, b: 3, expectedResult: 2, expectErr: false},
	// 	{a: 6, b: 0, expectedResult: 0, expectErr: true},
	// }

	testCases := []divideTestCase{
		{a: 6, b: 3, expectedResult: 2, expectErr: false},
		{a: 6, b: 0, expectedResult: 0, expectErr: true},
	}

	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)
		if tc.expectErr {
			assert.Errorf(t, err, "Divide(%d, %d) failed. Expected an error, got nil.", tc.a, tc.b)
		} else {
			if err != nil {
				assert.Errorf(t, err, "Divide(%d, %d) failed. Expected no error, got %v", tc.a, tc.b, err.Error())
			}
		}
		if result != tc.expectedResult {
			assert.Errorf(t, err, "Divide(%d, %d) returned %d, expected %d.", tc.a, tc.b, result, tc.expectedResult, "they should be equal")
		}
	}

}
