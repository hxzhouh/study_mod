1.
```go
package main

import "fmt"

func main() {
		var a = map[int]int{}
    	a[1] = 1
    	fmt.Println(a)
}
```

```bash
hxzhouh@hxzhouh:~/workspace/study_mod/assemble$ go tool compile -S map.go  |grep map.go:6
        0x0021 00033 (map.go:6) PCDATA  $1, $0
        0x0021 00033 (map.go:6) CALL    runtime.makemap_small(SB)
        0x0026 00038 (map.go:6) MOVQ    (SP), AX
        0x002a 00042 (map.go:6) MOVQ    AX, ""..autotmp_22+64(SP)
```
可以发现最终map函数调用的runtime函数是,runtime.makemap_small 根据这去找源码就行了.

再来看一个例子

```go
func map2() {
	var a = map[int]int{}
	a[1] = 1
	//fmt.Println(a)
}
```

```bash
hxzhouh@hxzhouh:~/workspace/study_mod/assemble$ go tool compile -S map.go  |grep map.go:13
        0x002f 00047 (map.go:13)        XORPS   X0, X0
        0x0032 00050 (map.go:13)        MOVUPS  X0, ""..autotmp_1+176(SP)
        0x003a 00058 (map.go:13)        MOVUPS  X0, ""..autotmp_1+192(SP)
        0x0042 00066 (map.go:13)        MOVUPS  X0, ""..autotmp_1+208(SP)
        0x004a 00074 (map.go:13)        LEAQ    ""..autotmp_2+32(SP), DI
        0x004f 00079 (map.go:13)        LEAQ    -48(DI), DI
        0x0053 00083 (map.go:13)        NOP
        0x0060 00096 (map.go:13)        DUFFZERO        $258
        0x0073 00115 (map.go:13)        LEAQ    ""..autotmp_2+32(SP), AX
        0x0078 00120 (map.go:13)        MOVQ    AX, ""..autotmp_1+192(SP)
        0x0080 00128 (map.go:13)        PCDATA  $1, $1
        0x0080 00128 (map.go:13)        CALL    runtime.fastrand(SB)
        0x0085 00133 (map.go:13)        MOVL    (SP), AX
        0x0088 00136 (map.go:13)        MOVL    AX, ""..autotmp_1+188(SP)
hxzhouh@hxzhouh:~/workspace/study_mod/assemble$ 
```
可以看到跟第一个例子截然不同的答案，因为这种写法的map是被直接分配到栈上的。如果不通过汇编，一定看不出结果。