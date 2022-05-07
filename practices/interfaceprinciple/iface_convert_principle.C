"".main STEXT size=272 args=0x0 locals=0x48 funcid=0x0 align=0x0
	0x0000 00000 (iface_convert_principle.go:9)	TEXT	"".main(SB), ABIInternal, $80-0
	0x0000 00000 (iface_convert_principle.go:9)	MOVD	16(g), R16
	0x0004 00004 (iface_convert_principle.go:9)	PCDATA	$0, $-2
	0x0004 00004 (iface_convert_principle.go:9)	MOVD	RSP, R17
	0x0008 00008 (iface_convert_principle.go:9)	CMP	R16, R17
	0x000c 00012 (iface_convert_principle.go:9)	BLS	260
	0x0010 00016 (iface_convert_principle.go:9)	PCDATA	$0, $-1
	0x0010 00016 (iface_convert_principle.go:9)	MOVD.W	R30, -80(RSP)
	0x0014 00020 (iface_convert_principle.go:9)	MOVD	R29, -8(RSP)
	0x0018 00024 (iface_convert_principle.go:9)	SUB	$8, RSP, R29
	0x001c 00028 (iface_convert_principle.go:9)	FUNCDATA	ZR, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001c 00028 (iface_convert_principle.go:9)	FUNCDATA	$1, gclocals·568470801006e5c0dc3947ea998fe279(SB)
	0x001c 00028 (iface_convert_principle.go:9)	FUNCDATA	$2, "".main.stkobj(SB)
	0x001c 00028 (iface_convert_principle.go:9)	PCDATA	$0, $-3
	0x001c 00028 (iface_convert_principle.go:10)	MOVD	"".EVALUE(SB), R0
	0x0028 00040 (iface_convert_principle.go:10)	PCDATA	$0, $-1
	0x0028 00040 (iface_convert_principle.go:10)	PCDATA	$1, ZR
	0x0028 00040 (iface_convert_principle.go:10)	CALL	runtime.convT64(SB)
	0x002c 00044 (iface_convert_principle.go:10)	MOVD	$type.int(SB), R1
	0x0034 00052 (iface_convert_principle.go:10)	PCDATA	$0, $-4
	0x0034 00052 (iface_convert_principle.go:10)	MOVD	R1, "".EBread(SB)
	0x0040 00064 (iface_convert_principle.go:10)	PCDATA	$0, $-1
	0x0040 00064 (iface_convert_principle.go:10)	PCDATA	ZR, $-2
	0x0040 00064 (iface_convert_principle.go:10)	MOVWU	runtime.writeBarrier(SB), R2
	0x004c 00076 (iface_convert_principle.go:10)	CBNZW	R2, 96
	0x0050 00080 (iface_convert_principle.go:10)	MOVD	R0, "".EBread+8(SB)
	0x005c 00092 (iface_convert_principle.go:10)	JMP	112
	0x0060 00096 (iface_convert_principle.go:10)	MOVD	$"".EBread+8(SB), R2
	0x0068 00104 (iface_convert_principle.go:10)	MOVD	R0, R3
	0x006c 00108 (iface_convert_principle.go:10)	CALL	runtime.gcWriteBarrier(SB)
	0x0070 00112 (iface_convert_principle.go:11)	PCDATA	ZR, $-1
	0x0070 00112 (iface_convert_principle.go:11)	PCDATA	$0, $-3
	0x0070 00112 (iface_convert_principle.go:11)	MOVD	"".EBread(SB), R0
	0x007c 00124 (iface_convert_principle.go:11)	PCDATA	$0, $-4
	0x007c 00124 (iface_convert_principle.go:11)	MOVD	"".EBread+8(SB), R3
	0x0088 00136 (iface_convert_principle.go:11)	PCDATA	$0, $-1
	0x0088 00136 (iface_convert_principle.go:11)	CMP	R1, R0
	0x008c 00140 (iface_convert_principle.go:11)	BNE	244
	0x0090 00144 (iface_convert_principle.go:11)	MOVD	(R3), R1
	0x0094 00148 (iface_convert_principle.go:11)	PCDATA	$0, $-3
	0x0094 00148 (iface_convert_principle.go:11)	MOVD	R1, "".a(SB)
	0x00a0 00160 (iface_convert_principle.go:11)	PCDATA	$0, $-1
	0x00a0 00160 (iface_convert_principle.go:12)	STP	(ZR, ZR), ""..autotmp_9-16(SP)
	0x00a4 00164 (iface_convert_principle.go:12)	PCDATA	$0, $-4
	0x00a4 00164 (iface_convert_principle.go:12)	MOVD	"".a(SB), R0
	0x00b0 00176 (iface_convert_principle.go:12)	PCDATA	$0, $-1
	0x00b0 00176 (iface_convert_principle.go:12)	PCDATA	$1, $1
	0x00b0 00176 (iface_convert_principle.go:12)	CALL	runtime.convT64(SB)
	0x00b4 00180 (iface_convert_principle.go:12)	MOVD	$type.int(SB), R1
	0x00bc 00188 (iface_convert_principle.go:12)	MOVD	R1, ""..autotmp_9-16(SP)
	0x00c0 00192 (iface_convert_principle.go:12)	MOVD	R0, ""..autotmp_9-8(SP)
	0x00c4 00196 (<unknown line number>)	NOP
	0x00c4 00196 (<unknown line number>)	PCDATA	$0, $-3
	0x00c4 00196 ($GOROOT/src/fmt/print.go:242)	MOVD	os.Stdout(SB), R1
	0x00d0 00208 ($GOROOT/src/fmt/print.go:242)	PCDATA	$0, $-1
	0x00d0 00208 ($GOROOT/src/fmt/print.go:242)	MOVD	$go.itab.*os.File,io.Writer(SB), R0
	0x00d8 00216 ($GOROOT/src/fmt/print.go:242)	MOVD	$""..autotmp_9-16(SP), R2
	0x00dc 00220 ($GOROOT/src/fmt/print.go:242)	MOVD	$1, R3
	0x00e0 00224 ($GOROOT/src/fmt/print.go:242)	MOVD	R3, R4
	0x00e4 00228 ($GOROOT/src/fmt/print.go:242)	PCDATA	$1, ZR
	0x00e4 00228 ($GOROOT/src/fmt/print.go:242)	CALL	fmt.Fprint(SB)
	0x00e8 00232 (iface_convert_principle.go:13)	MOVD	-8(RSP), R29
	0x00ec 00236 (iface_convert_principle.go:13)	MOVD.P	80(RSP), R30
	0x00f0 00240 (iface_convert_principle.go:13)	RET	(R30)
	0x00f4 00244 (iface_convert_principle.go:11)	MOVD	$type.interface {}(SB), R2
	0x00fc 00252 (iface_convert_principle.go:11)	CALL	runtime.panicdottypeE(SB)
	0x0100 00256 (iface_convert_principle.go:11)	HINT	ZR
	0x0104 00260 (iface_convert_principle.go:11)	NOP
	0x0104 00260 (iface_convert_principle.go:9)	PCDATA	$1, $-1
	0x0104 00260 (iface_convert_principle.go:9)	PCDATA	$0, $-2
	0x0104 00260 (iface_convert_principle.go:9)	MOVD	R30, R3
	0x0108 00264 (iface_convert_principle.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x010c 00268 (iface_convert_principle.go:9)	PCDATA	$0, $-1
	0x010c 00268 (iface_convert_principle.go:9)	JMP	0
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb c9 07 00 54  ..@.....?......T
	0x0010 fe 0f 1b f8 fd 83 1f f8 fd 23 00 d1 1b 00 00 90  .........#......
	0x0020 7b 03 00 91 60 03 40 f9 00 00 00 94 01 00 00 90  {...`.@.........
	0x0030 21 00 00 91 1b 00 00 90 7b 03 00 91 61 03 00 f9  !.......{...a...
	0x0040 1b 00 00 90 7b 03 00 91 62 03 40 b9 a2 00 00 35  ....{...b.@....5
	0x0050 1b 00 00 90 7b 03 00 91 60 03 00 f9 05 00 00 14  ....{...`.......
	0x0060 02 00 00 90 42 00 00 91 e3 03 00 aa 00 00 00 94  ....B...........
	0x0070 1b 00 00 90 7b 03 00 91 60 03 40 f9 1b 00 00 90  ....{...`.@.....
	0x0080 7b 03 00 91 63 03 40 f9 1f 00 01 eb 41 03 00 54  {...c.@.....A..T
	0x0090 61 00 40 f9 1b 00 00 90 7b 03 00 91 61 03 00 f9  a.@.....{...a...
	0x00a0 ff ff 03 a9 1b 00 00 90 7b 03 00 91 60 03 40 f9  ........{...`.@.
	0x00b0 00 00 00 94 01 00 00 90 21 00 00 91 e1 1f 00 f9  ........!.......
	0x00c0 e0 23 00 f9 1b 00 00 90 7b 03 00 91 61 03 40 f9  .#......{...a.@.
	0x00d0 00 00 00 90 00 00 00 91 e2 e3 00 91 e3 03 40 b2  ..............@.
	0x00e0 e4 03 03 aa 00 00 00 94 fd 83 5f f8 fe 07 45 f8  .........._...E.
	0x00f0 c0 03 5f d6 02 00 00 90 42 00 00 91 00 00 00 94  .._.....B.......
	0x0100 1f 20 03 d5 e3 03 1e aa 00 00 00 94 bd ff ff 17  . ..............
	rel 0+0 t=23 type.int+0
	rel 0+0 t=23 type.int+0
	rel 0+0 t=23 type.*os.File+0
	rel 28+8 t=3 "".EVALUE+0
	rel 40+4 t=9 runtime.convT64+0
	rel 44+8 t=3 type.int+0
	rel 52+8 t=3 "".EBread+0
	rel 64+8 t=3 runtime.writeBarrier+0
	rel 80+8 t=3 "".EBread+8
	rel 96+8 t=3 "".EBread+8
	rel 108+4 t=9 runtime.gcWriteBarrier+0
	rel 112+8 t=3 "".EBread+0
	rel 124+8 t=3 "".EBread+8
	rel 148+8 t=3 "".a+0
	rel 164+8 t=3 "".a+0
	rel 176+4 t=9 runtime.convT64+0
	rel 180+8 t=3 type.int+0
	rel 196+8 t=3 os.Stdout+0
	rel 208+8 t=3 go.itab.*os.File,io.Writer+0
	rel 228+4 t=9 fmt.Fprint+0
	rel 244+8 t=3 type.interface {}+0
	rel 252+4 t=9 runtime.panicdottypeE+0
	rel 264+4 t=9 runtime.morestack_noctxt+0
"".init STEXT size=16 args=0x0 locals=0x0 funcid=0x0 align=0x0 leaf
	0x0000 00000 (iface_convert_principle.go:5)	TEXT	"".init(SB), LEAF|NOFRAME|ABIInternal, $0-0
	0x0000 00000 (iface_convert_principle.go:5)	FUNCDATA	ZR, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (iface_convert_principle.go:5)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (iface_convert_principle.go:5)	RET	(R30)
	0x0000 c0 03 5f d6 00 00 00 00 00 00 00 00 00 00 00 00  .._.............
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info.fmt.Print$abstract SDWARFABSFCN dupok size=40
	0x0000 05 66 6d 74 2e 50 72 69 6e 74 00 01 01 13 61 00  .fmt.Print....a.
	0x0010 00 00 00 00 00 13 6e 00 01 00 00 00 00 13 65 72  ......n.......er
	0x0020 72 00 01 00 00 00 00 00                          r.......
	rel 0+0 t=22 type.[]interface {}+0
	rel 0+0 t=22 type.error+0
	rel 0+0 t=22 type.int+0
	rel 17+4 t=31 go.info.[]interface {}+0
	rel 25+4 t=31 go.info.int+0
	rel 35+4 t=31 go.info.error+0
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=-32767 os.(*File).Write+0
"".EBread SBSS size=16
"".a SNOPTRBSS size=8
"".EVALUE SNOPTRDATA size=8
	0x0000 9a 02 00 00 00 00 00 00                          ........
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·568470801006e5c0dc3947ea998fe279 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 02                    ..........
"".main.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
