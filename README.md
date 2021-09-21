# satunna

`satunna` is a reproducible pseudo-random code execution hashing algo.   
It uses a virtual machine that executes pseudo-random instructions that produce a 256-bit Blake2b cryptographic hash.   

## The Virtual Machine

### Seed

The seed is used to determine which instructions in what order are used on the user data.   
The first 7 charachters of the Blake2b hash of the input, plus a parity bit determine the seed.

## Instruction set

|instruction|dst|src|`src == dst ?`|operation|
|-|-|-|-|-|
|IADD_RS|R|R|`src = dst`|`dst = dst + (src << mod.shift) (+ imm32)`|
|IADD_M|R|R|`src = 0`|`dst = dst + [mem]`|
|ISUB_R|R|R|`src = imm32`|`dst = dst - src`|
|ISUB_M|R|R|`src = 0`|`dst = dst - [mem]`|
|IMUL_R|R|R|`src = imm32`|`dst = dst * src`|
|IMUL_M|R|R|`src = 0`|`dst = dst * [mem]`|
|IMULH_R|R|R|`src = dst`|`dst = (dst * src) >> 64`|
|IMULH_M|R|R|`src = 0`|`dst = (dst * [mem]) >> 64`|
|ISMULH_R|R|R|`src = dst`|`dst = (dst * src) >> 64` (signed)|
|ISMULH_M|R|R|`src = 0`|`dst = (dst * [mem]) >> 64` (signed)|
|IMUL_RCP|R|-|-|<code>dst = 2<sup>x</sup> / imm32 * dst</code>|
|INEG_R|R|-|-|`dst = -dst`|
|IXOR_R|R|R|`src = imm32`|`dst = dst ^ src`|
|IXOR_M|R|R|`src = 0`|`dst = dst ^ [mem]`|
|IROR_R|R|R|`src = imm32`|`dst = dst >>> src`|
|IROL_R|R|R|`src = imm32`|`dst = dst <<< src`|
|ISWAP_R|R|R|`src = dst`|`temp = src; src = dst; dst = temp`|