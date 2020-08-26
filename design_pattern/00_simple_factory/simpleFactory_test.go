/*
@Time : 2020/3/31 15:47
@Author : ZhouHui2
@File : simpleFactory_test.go
@Software: GoLand
*/
package _0_simple_factory

import (
	"strings"
	"testing"
)

func TestNewDuck(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"blackDuck", args{2}, "BlackDuck say hello"},
		{"yellowDuck", args{1}, "yellowDuck say hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDuck(tt.args.t)
			result := got.Say("hello")
			if ! strings.EqualFold(result, tt.want) {
				t.Errorf("result() = %v, want %v", result, tt.want)
			}
		})
	}
}
