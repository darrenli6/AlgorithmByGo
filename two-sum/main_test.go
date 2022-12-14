package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 13
	expected := []int{0, 2}
	res := twoSum(nums, target)
	t.Logf("res : %#v", res)

	if !reflect.DeepEqual(expected, res) {
		t.Fail()
	}
}
