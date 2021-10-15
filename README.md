# satunna

`satunna` is a reproducible pseudo-random code execution hashing algo.   
It uses a virtual machine that executes pseudo-random instructions that produce a 256-bit Blake2b cryptographic hash.   

## Specifications

***__SPECS ARE WIP AND ARE SUBJECT TO CHANGE__***

### The Virtual Machine

#### Seed

The seed is used to determine which instructions in what order are used on the inputed data.   
The seed is a positive intenger derived from the first 16 bytes of the blake2b hash of the file(or string).

#### The Instruction Set

The satunnaVM is a CISC, turing complete virtual machine that is designed to compute 64-bit integers and nothing more.  


The current VM instruction set is being worked on as `satunna` is being developed.

|instruction|dst|src|`src == dst ?`|operation|
|-|-|-|-|-|
|IADD_M|R|R|`src = 0`|`dst = dst + [mem]`|
|ISUB_R|R|R|`src = imm32`|`dst = dst - src`|
|ISUB_M|R|R|`src = 0`|`dst = dst - [mem]`|
|IMUL_R|R|R|`src = imm32`|`dst = dst * src`|
|IMUL_M|R|R|`src = 0`|`dst = dst * [mem]`|
|IMULH_R|R|R|`src = dst`|`dst = (dst * src) >> 64`|
|IMULH_M|R|R|`src = 0`|`dst = (dst * [mem]) >> 64`|
|IMUL_RCP|R|-|-|<code>dst = 2<sup>x</sup> / imm32 * dst</code>|
|INEG_R|R|-|-|`dst = -dst`|
|IXOR_R|R|R|`src = imm32`|`dst = dst ^ src`|
|IXOR_M|R|R|`src = 0`|`dst = dst ^ [mem]`|
|IROR_R|R|R|`src = imm32`|`dst = dst >>> src`|
|IROL_R|R|R|`src = imm32`|`dst = dst <<< src`|
|ISWAP_R|R|R|`src = dst`|`temp = src; src = dst; dst = temp`|