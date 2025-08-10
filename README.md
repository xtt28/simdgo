# Welcome to simdgo's source repository.

Many modern processors come with support for SIMD technology (single
instruction, multiple data), which allow for the same mathematical operation to
be performed on a set of numbers concurrently on a single core. simdgo library
uses x86 assembly SIMD instructions to allow you to leverage SIMD instructions
in your Go code.

## Use the library

### Installing

Execute in the shell of an active Go project:

    go get github.com/xtt28/simdgo

### Compatibility check

You can use `golang.org/x/sys/cpu` package to check if the CPU supports SSE and
AVX instructions.

| Instruction set extension | Functionality                        | Check if supported |
| ------------------------- | ------------------------------------ | ------------------ |
| SSE                       | 4 x `float32` vectors                | `cpu.X86.HasSSE`   |
| AVX                       | 8 x `float32`, 4 x `float64` vectors | `cpu.X86.HasAVX`   |

### Performing math



## Benchmarks

Benchmarks are coming soon.

## Features

- Arithmetic operations on floats (add, subtract, multiply, divide)
- 128- and 256-bit vectors (AVX needed for 256-bit vectors)

## License

This library is licensed under:

    SPDX-License-Identifier: MIT

being in concordance with the terms found in the LICENSE file in the root of
this source repository.