# Golang-refresher

## Purpose

To gain an understanding of Go.


## Basic Commands

- go version
- go run <file.go>


### Int types

#### Signed ints

Use case: Signed Integers: Use when you need to represent negative numbers (e.g., temperatures, financial calculations with losses, etc.).

type    size           range
int8	8 bits/1 byte	-128 to 127
int16	16 bits/2 byte	-32768 to 32767
int32	32 bits/4 byte	-2147483648 to 2147483647
int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807

int	Depends on platform:

32 bits in 32 bit systems and
64 bit in 64 bit systems

-2147483648 to 2147483647 in 32 bit systems and
-9223372036854775808 to 9223372036854775807 in 64 bit systems


#### Unsigned ints

Use case: Use when you are sure the values will always be non-negative and need a larger range (e.g., indexing arrays, bit manipulation, or representing sizes).

uint8: Range is 0 to 255.
uint16: Range is 0 to 65,535.
uint32: Range is 0 to 4,294,967,295.
uint64: Range is 0 to 18,446,744,073,709,551,615.