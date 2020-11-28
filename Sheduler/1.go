package main

func hello(msg string) {
	println(msg)
}

func main() {
	go hello("hello world")
}

/**

hxzhouh@hxzhouh:~/workspace/study_mod/Sheduler$ go tool compile -S 1.go
"".hello STEXT size=91 args=0x10 locals=0x18
        0x0000 00000 (1.go:3)   TEXT    "".hello(SB), ABIInternal, $24-16
        0x0000 00000 (1.go:3)   MOVQ    (TLS), CX
        0x0009 00009 (1.go:3)   CMPQ    SP, 16(CX)
        0x000d 00013 (1.go:3)   PCDATA  $0, $-2
        0x000d 00013 (1.go:3)   JLS     84
        0x000f 00015 (1.go:3)   PCDATA  $0, $-1
        0x000f 00015 (1.go:3)   SUBQ    $24, SP
        0x0013 00019 (1.go:3)   MOVQ    BP, 16(SP)
        0x0018 00024 (1.go:3)   LEAQ    16(SP), BP
        0x001d 00029 (1.go:3)   FUNCDATA        $0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
        0x001d 00029 (1.go:3)   FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x001d 00029 (1.go:4)   PCDATA  $1, $0
        0x001d 00029 (1.go:4)   NOP
        0x0020 00032 (1.go:4)   CALL    runtime.printlock(SB)
        0x0025 00037 (1.go:4)   MOVQ    "".msg+32(SP), AX
        0x002a 00042 (1.go:4)   MOVQ    AX, (SP)
        0x002e 00046 (1.go:4)   MOVQ    "".msg+40(SP), AX
        0x0033 00051 (1.go:4)   MOVQ    AX, 8(SP)
        0x0038 00056 (1.go:4)   PCDATA  $1, $1
        0x0038 00056 (1.go:4)   CALL    runtime.printstring(SB)
        0x003d 00061 (1.go:4)   NOP
        0x0040 00064 (1.go:4)   CALL    runtime.printnl(SB)
        0x0045 00069 (1.go:4)   CALL    runtime.printunlock(SB)
        0x004a 00074 (1.go:5)   MOVQ    16(SP), BP
        0x004f 00079 (1.go:5)   ADDQ    $24, SP
        0x0053 00083 (1.go:5)   RET
        0x0054 00084 (1.go:5)   NOP
        0x0054 00084 (1.go:3)   PCDATA  $1, $-1
        0x0054 00084 (1.go:3)   PCDATA  $0, $-2
        0x0054 00084 (1.go:3)   CALL    runtime.morestack_noctxt(SB)
        0x0059 00089 (1.go:3)   PCDATA  $0, $-1
        0x0059 00089 (1.go:3)   JMP     0
        0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 45 48  dH..%....H;a.vEH
        0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 0f 1f 00  ...H.l$.H.l$....
        0x0020 e8 00 00 00 00 48 8b 44 24 20 48 89 04 24 48 8b  .....H.D$ H..$H.
        0x0030 44 24 28 48 89 44 24 08 e8 00 00 00 00 0f 1f 00  D$(H.D$.........
        0x0040 e8 00 00 00 00 e8 00 00 00 00 48 8b 6c 24 10 48  ..........H.l$.H
        0x0050 83 c4 18 c3 e8 00 00 00 00 eb a5                 ...........
        rel 5+4 t=17 TLS+0
        rel 33+4 t=8 runtime.printlock+0
        rel 57+4 t=8 runtime.printstring+0
        rel 65+4 t=8 runtime.printnl+0
        rel 70+4 t=8 runtime.printunlock+0
        rel 85+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=91 args=0x0 locals=0x28
        0x0000 00000 (1.go:7)   TEXT    "".main(SB), ABIInternal, $40-0
        0x0000 00000 (1.go:7)   MOVQ    (TLS), CX
        0x0009 00009 (1.go:7)   CMPQ    SP, 16(CX)
        0x000d 00013 (1.go:7)   PCDATA  $0, $-2
        0x000d 00013 (1.go:7)   JLS     84
        0x000f 00015 (1.go:7)   PCDATA  $0, $-1
        0x000f 00015 (1.go:7)   SUBQ    $40, SP
        0x0013 00019 (1.go:7)   MOVQ    BP, 32(SP)
        0x0018 00024 (1.go:7)   LEAQ    32(SP), BP
        0x001d 00029 (1.go:7)   FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (1.go:7)   FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (1.go:8)   MOVL    $16, (SP)  // 将最后一个参数的位置存到栈顶 0x00
        0x0024 00036 (1.go:8)   LEAQ    "".hello·f(SB), AX  // 将 go 语句调用的函数入口地址给 AX
        0x002b 00043 (1.go:8)   MOVQ    AX, 8(SP) // // 将 AX 存入 0x08
        0x0030 00048 (1.go:8)   LEAQ    go.string."hello world"(SB), AX // 将 "hello world" 的地址给 AX
        0x0037 00055 (1.go:8)   MOVQ    AX, 16(SP)    // 将 AX 的值放到SP
        0x003c 00060 (1.go:8)   MOVQ    $11, 24(SP)
        0x0045 00069 (1.go:8)   PCDATA  $1, $0
        0x0045 00069 (1.go:8)   CALL    runtime.newproc(SB) // 调用 newproc
        0x004a 00074 (1.go:9)   MOVQ    32(SP), BP
        0x004f 00079 (1.go:9)   ADDQ    $40, SP
        0x0053 00083 (1.go:9)   RET
        0x0054 00084 (1.go:9)   NOP
        0x0054 00084 (1.go:7)   PCDATA  $1, $-1
        0x0054 00084 (1.go:7)   PCDATA  $0, $-2
        0x0054 00084 (1.go:7)   CALL    runtime.morestack_noctxt(SB)
        0x0059 00089 (1.go:7)   PCDATA  $0, $-1
        0x0059 00089 (1.go:7)   JMP     0
        0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 45 48  dH..%....H;a.vEH
        0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 c7 04 24  ..(H.l$ H.l$ ..$
        0x0020 10 00 00 00 48 8d 05 00 00 00 00 48 89 44 24 08  ....H......H.D$.
        0x0030 48 8d 05 00 00 00 00 48 89 44 24 10 48 c7 44 24  H......H.D$.H.D$
        0x0040 18 0b 00 00 00 e8 00 00 00 00 48 8b 6c 24 20 48  ..........H.l$ H
        0x0050 83 c4 28 c3 e8 00 00 00 00 eb a5                 ..(........
        rel 5+4 t=17 TLS+0
        rel 39+4 t=16 "".hello·f+0
        rel 51+4 t=16 go.string."hello world"+0
        rel 70+4 t=8 runtime.newproc+0
        rel 85+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
go.string."hello world" SRODATA dupok size=11
        0x0000 68 65 6c 6c 6f 20 77 6f 72 6c 64                 hello world
""..inittask SNOPTRDATA size=24
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00                          ........
"".hello·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 "".hello+0
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
        0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
*/
