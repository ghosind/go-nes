# go-nes

![Test](https://github.com/ghosind/go-nes/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghosind/go-nes)](https://goreportcard.com/report/github.com/ghosind/go-nes)
[![codecov](https://codecov.io/gh/ghosind/go-nes/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/go-nes)
![License Badge](https://img.shields.io/github/license/ghosind/go-nes)

go-nes is a focused, well-tested Go library that implements the core internals of an NES emulator. It intentionally excludes GUI and real hardware I/O (display/windowing, real controllers, audio output). The goal is to provide a clean, reusable emulator core you can embed into your own frontend, educational tooling, or testing harness.

## Features

- Accurate 6502 CPU emulation with support for all official opcodes and addressing modes.
<!-- - Basic PPU implementation for rendering graphics (background and sprites). -->
<!-- - APU stubs for audio processing (full audio implementation is a work in progress). -->
<!-- - Memory mapping that supports various NES mappers (NROM, MMC1, etc.). -->
- Loading and parsing of iNES format ROM files.

## Components

- `nes`: High-level NES struct that ties together CPU, PPU, APU, and memory.
- `cpu`: Implements the 6502 CPU with all instructions, addressing modes, and status flags.
- `ppu`: Handles the NES Picture Processing Unit (PPU) for graphics rendering.
- `apu`: Manages the NES Audio Processing Unit (APU) for sound generation.
- `memory`: Emulates the NES memory map, including RAM, ROM, and I/O registers.
- `rom`: Loads and parses NES ROM files in iNES format.
- `rom/ines`: Loads and parses iNES header.
- `rom/mapper`: Supports various NES cartridge mappers for memory banking.
- more components to be added...

## Reference

- [NES Documentation](https://www.nesdev.org/NESDoc.pdf)
- [Obelisk 6502 Guide](https://www.nesdev.org/obelisk-6502-guide/index.html)
- [iNES File Format](https://www.nesdev.org/wiki/INES)
- [INES 1.0 Mapper Grid](https://www.nesdev.org/wiki/Mapper)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
