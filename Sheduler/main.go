package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg)
	}
	wg.Wait()
	// Wait to see the global run queue deplete.
	time.Sleep(3 * time.Second)
}
func work(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	var counter int
	for i := 0; i < 1e10; i++ {
		counter++
	}
	wg.Done()
}

/**
hxzhouh@hxzhouh:~/workspace/study_mod/Sheduler$ go tool compile -S main.go
"".main STEXT size=185 args=0x0 locals=0x30
        0x0000 00000 (main.go:8)        TEXT    "".main(SB), ABIInternal, $48-0
        0x0000 00000 (main.go:8)        MOVQ    (TLS), CX
        0x0009 00009 (main.go:8)        CMPQ    SP, 16(CX)
        0x000d 00013 (main.go:8)        PCDATA  $0, $-2
        0x000d 00013 (main.go:8)        JLS     175
        0x0013 00019 (main.go:8)        PCDATA  $0, $-1
        0x0013 00019 (main.go:8)        SUBQ    $48, SP
        0x0017 00023 (main.go:8)        MOVQ    BP, 40(SP)
        0x001c 00028 (main.go:8)        LEAQ    40(SP), BP
        0x0021 00033 (main.go:8)        FUNCDATA        $0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x0021 00033 (main.go:8)        FUNCDATA        $1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x0021 00033 (main.go:9)        LEAQ    type.sync.WaitGroup(SB), AX
        0x0028 00040 (main.go:9)        MOVQ    AX, (SP)
        0x002c 00044 (main.go:9)        PCDATA  $1, $0
        0x002c 00044 (main.go:9)        CALL    runtime.newobject(SB)
        0x0031 00049 (main.go:9)        MOVQ    8(SP), AX
        0x0036 00054 (main.go:9)        MOVQ    AX, "".&wg+32(SP)
        0x003b 00059 (main.go:10)       MOVQ    AX, (SP)
        0x003f 00063 (main.go:10)       MOVQ    $10, 8(SP)
        0x0048 00072 (main.go:10)       PCDATA  $1, $1
        0x0048 00072 (main.go:10)       CALL    sync.(*WaitGroup).Add(SB)
        0x004d 00077 (main.go:10)       XORL    AX, AX
        0x004f 00079 (main.go:11)       JMP     128
        0x0051 00081 (main.go:11)       MOVQ    AX, "".i+24(SP)
        0x0056 00086 (main.go:12)       MOVL    $8, (SP)
        0x005d 00093 (main.go:12)       LEAQ    "".work·f(SB), AX
        0x0064 00100 (main.go:12)       MOVQ    AX, 8(SP)
        0x0069 00105 (main.go:12)       MOVQ    "".&wg+32(SP), CX
        0x006e 00110 (main.go:12)       MOVQ    CX, 16(SP)
        0x0073 00115 (main.go:12)       CALL    runtime.newproc(SB)
        0x0078 00120 (main.go:11)       MOVQ    "".i+24(SP), AX
        0x007d 00125 (main.go:11)       INCQ    AX
        0x0080 00128 (main.go:11)       CMPQ    AX, $10
        0x0084 00132 (main.go:11)       JLT     81
        0x0086 00134 (main.go:14)       MOVQ    "".&wg+32(SP), AX
        0x008b 00139 (main.go:14)       MOVQ    AX, (SP)
        0x008f 00143 (main.go:14)       PCDATA  $1, $0
        0x008f 00143 (main.go:14)       CALL    sync.(*WaitGroup).Wait(SB)
        0x0094 00148 (main.go:16)       MOVL    $3000000000, AX
        0x0099 00153 (main.go:16)       MOVQ    AX, (SP)
        0x009d 00157 (main.go:16)       NOP
        0x00a0 00160 (main.go:16)       CALL    time.Sleep(SB)
        0x00a5 00165 (main.go:17)       MOVQ    40(SP), BP
        0x00aa 00170 (main.go:17)       ADDQ    $48, SP
        0x00ae 00174 (main.go:17)       RET
        0x00af 00175 (main.go:17)       NOP
        0x00af 00175 (main.go:8)        PCDATA  $1, $-1
        0x00af 00175 (main.go:8)        PCDATA  $0, $-2
        0x00af 00175 (main.go:8)        CALL    runtime.morestack_noctxt(SB)
        0x00b4 00180 (main.go:8)        PCDATA  $0, $-1
        0x00b4 00180 (main.go:8)        JMP     0
        0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 9c  dH..%....H;a....
        0x0010 00 00 00 48 83 ec 30 48 89 6c 24 28 48 8d 6c 24  ...H..0H.l$(H.l$
        0x0020 28 48 8d 05 00 00 00 00 48 89 04 24 e8 00 00 00  (H......H..$....
        0x0030 00 48 8b 44 24 08 48 89 44 24 20 48 89 04 24 48  .H.D$.H.D$ H..$H
        0x0040 c7 44 24 08 0a 00 00 00 e8 00 00 00 00 31 c0 eb  .D$..........1..
        0x0050 2f 48 89 44 24 18 c7 04 24 08 00 00 00 48 8d 05  /H.D$...$....H..
        0x0060 00 00 00 00 48 89 44 24 08 48 8b 4c 24 20 48 89  ....H.D$.H.L$ H.
        0x0070 4c 24 10 e8 00 00 00 00 48 8b 44 24 18 48 ff c0  L$......H.D$.H..
        0x0080 48 83 f8 0a 7c cb 48 8b 44 24 20 48 89 04 24 e8  H...|.H.D$ H..$.
        0x0090 00 00 00 00 b8 00 5e d0 b2 48 89 04 24 0f 1f 00  ......^..H..$...
        0x00a0 e8 00 00 00 00 48 8b 6c 24 28 48 83 c4 30 c3 e8  .....H.l$(H..0..
        0x00b0 00 00 00 00 e9 47 ff ff ff                       .....G...
        rel 5+4 t=17 TLS+0
        rel 36+4 t=16 type.sync.WaitGroup+0
        rel 45+4 t=8 runtime.newobject+0
        rel 73+4 t=8 sync.(*WaitGroup).Add+0
        rel 96+4 t=16 "".work·f+0
        rel 116+4 t=8 runtime.newproc+0
        rel 144+4 t=8 sync.(*WaitGroup).Wait+0
        rel 161+4 t=8 time.Sleep+0
        rel 176+4 t=8 runtime.morestack_noctxt+0
"".work STEXT size=109 args=0x8 locals=0x18
        0x0000 00000 (main.go:18)       TEXT    "".work(SB), ABIInternal, $24-8
        0x0000 00000 (main.go:18)       MOVQ    (TLS), CX
        0x0009 00009 (main.go:18)       CMPQ    SP, 16(CX)
        0x000d 00013 (main.go:18)       PCDATA  $0, $-2
        0x000d 00013 (main.go:18)       JLS     102
        0x000f 00015 (main.go:18)       PCDATA  $0, $-1
        0x000f 00015 (main.go:18)       SUBQ    $24, SP
        0x0013 00019 (main.go:18)       MOVQ    BP, 16(SP)
        0x0018 00024 (main.go:18)       LEAQ    16(SP), BP
        0x001d 00029 (main.go:18)       FUNCDATA        $0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
        0x001d 00029 (main.go:18)       FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x001d 00029 (main.go:19)       MOVQ    $1000000000, (SP)
        0x0025 00037 (main.go:19)       PCDATA  $1, $0
        0x0025 00037 (main.go:19)       CALL    time.Sleep(SB)
        0x002a 00042 (main.go:19)       XORL    AX, AX
        0x002c 00044 (main.go:21)       JMP     49
        0x002e 00046 (main.go:21)       INCQ    AX
        0x0031 00049 (main.go:21)       MOVQ    $10000000000, CX
        0x003b 00059 (main.go:21)       NOP
        0x0040 00064 (main.go:21)       CMPQ    AX, CX
        0x0043 00067 (main.go:21)       JLT     46
        0x0045 00069 (main.go:24)       PCDATA  $1, $-1
        0x0045 00069 (<unknown line number>)    NOP
        0x0045 00069 ($GOROOT/src/sync/waitgroup.go:99) MOVQ    "".wg+32(SP), AX
        0x004a 00074 ($GOROOT/src/sync/waitgroup.go:99) MOVQ    AX, (SP)
        0x004e 00078 ($GOROOT/src/sync/waitgroup.go:99) MOVQ    $-1, 8(SP)
        0x0057 00087 ($GOROOT/src/sync/waitgroup.go:99) PCDATA  $1, $1
        0x0057 00087 ($GOROOT/src/sync/waitgroup.go:99) CALL    sync.(*WaitGroup).Add(SB)
        0x005c 00092 (main.go:24)       MOVQ    16(SP), BP
        0x0061 00097 (main.go:24)       ADDQ    $24, SP
        0x0065 00101 (main.go:24)       RET
        0x0066 00102 (main.go:24)       NOP
        0x0066 00102 (main.go:18)       PCDATA  $1, $-1
        0x0066 00102 (main.go:18)       PCDATA  $0, $-2
        0x0066 00102 (main.go:18)       CALL    runtime.morestack_noctxt(SB)
        0x006b 00107 (main.go:18)       PCDATA  $0, $-1
        0x006b 00107 (main.go:18)       JMP     0
        0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 57 48  dH..%....H;a.vWH
        0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 c7 04  ...H.l$.H.l$.H..
        0x0020 24 00 ca 9a 3b e8 00 00 00 00 31 c0 eb 03 48 ff  $...;.....1...H.
        0x0030 c0 48 b9 00 e4 0b 54 02 00 00 00 0f 1f 44 00 00  .H....T......D..
        0x0040 48 39 c8 7c e9 48 8b 44 24 20 48 89 04 24 48 c7  H9.|.H.D$ H..$H.
        0x0050 44 24 08 ff ff ff ff e8 00 00 00 00 48 8b 6c 24  D$..........H.l$
        0x0060 10 48 83 c4 18 c3 e8 00 00 00 00 eb 93           .H...........
        rel 5+4 t=17 TLS+0
        rel 38+4 t=8 time.Sleep+0
        rel 88+4 t=8 sync.(*WaitGroup).Add+0
        rel 103+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
go.info.sync.(*WaitGroup).Done$abstract SDWARFINFO dupok size=36
        0x0000 04 73 79 6e 63 2e 28 2a 57 61 69 74 47 72 6f 75  .sync.(*WaitGrou
        0x0010 70 29 2e 44 6f 6e 65 00 01 01 11 77 67 00 00 00  p).Done....wg...
        0x0020 00 00 00 00                                      ....
        rel 0+0 t=24 type.*sync.WaitGroup+0
        rel 31+4 t=29 go.info.*sync.WaitGroup+0
""..inittask SNOPTRDATA size=40
        0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0020 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 sync..inittask+0
        rel 32+8 t=1 time..inittask+0
"".work·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 "".work+0
type..importpath.sync. SRODATA dupok size=7
        0x0000 00 00 04 73 79 6e 63                             ...sync
type..importpath.time. SRODATA dupok size=7
        0x0000 00 00 04 74 69 6d 65                             ...time
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
        0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10

hxzhouh@hxzhouh:~/workspace/study_mod/Sheduler$
        0x0000 02 00 00 00 01 00 00 00 01 00                    ..........

*/
