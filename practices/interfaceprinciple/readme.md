Interface Know types always is faster than unknown types, almost 4-5 times faster than unknown types.

Interface known types are always doing convert during compile time.
Interface unknown types are always doing convert during invoke time, not doing convert during compile time.
Doing convert during compile time is faster than doing convert doing invoke time.

#反汇编
go tool compile -S xxx.go >> xxx.S 