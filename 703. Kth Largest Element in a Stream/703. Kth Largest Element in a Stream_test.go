package _703

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		k    int
		nums []int
	}
	tests := struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		name: "test",
		args: args{k: 2, nums: []int{0}},
	}
	t.Run(tests.name, func(t *testing.T) {
		obj := Constructor(tests.args.k, tests.args.nums)
		fmt.Println(obj.Add(-1))
		fmt.Println(obj.Add(1))
		fmt.Println(obj.Add(-2))
		fmt.Println(obj.Add(-4))
		fmt.Println(obj.Add(3))
	})
}
