/*
@Time : 2020/4/7 11:16
@Author : ZhouHui2
@File : mem_test
@Software: GoLand
*/
package test_example

import (
	"fmt"
	"strconv"
	"testing"
)

//func TestInsert(t *testing.T) {
//	type args struct {
//		userId int
//		corpId int
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{"1", args{
//			userId: 3214007,
//			corpId: 3214008,
//		}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			Insert(tt.args.userId, tt.args.corpId)
//		})
//	}
//}
//
//func BenchmarkInsert(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Insert(i, i<<1)
//	}
//}

func BenchmarkSprintf(b *testing.B){
	num:=10
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		fmt.Sprintf("%d",num)
	}
}

func BenchmarkFormat(b *testing.B){
	num:=int64(10)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		strconv.FormatInt(num,10)
	}
}

func BenchmarkItoa(b *testing.B){
	num:=10
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		strconv.Itoa(num)
	}
}
