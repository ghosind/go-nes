# go-nes

![Test](https://github.com/ghosind/go-nes/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghosind/go-nes)](https://goreportcard.com/report/github.com/ghosind/go-nes)
[![codecov](https://codecov.io/gh/ghosind/go-nes/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/go-nes)
![License Badge](https://img.shields.io/github/license/ghosind/go-nes)

go-nes is a focused, well-tested Go library that implements the core internals of an NES emulator. It intentionally excludes GUI and real hardware I/O (display/windowing, real controllers, audio output). The goal is to provide a clean, reusable emulator core you can embed into your own frontend, educational tooling, or testing harness.

## Features

- Accurate 6502 CPU emulation with support for all official opcodes and addressing modes.

## Components

- `cpu`: Implements the 6502 CPU with all instructions, addressing modes, and status flags.
- `memory`: Emulates the NES memory map, including RAM, ROM, and I/O registers.
- more components to be added...

## Reference

- [Obelisk 6502 Guide](https://www.nesdev.org/obelisk-6502-guide/index.html)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
