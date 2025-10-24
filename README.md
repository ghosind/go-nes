# go-nes

![Test](https://github.com/ghosind/go-nes/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghosind/go-nes)](https://goreportcard.com/report/github.com/ghosind/go-nes)
[![codecov](https://codecov.io/gh/ghosind/go-nes/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/go-nes)
![License Badge](https://img.shields.io/github/license/ghosind/go-nes)

go-nes is a focused, well-tested Go library that implements the core internals of an NES emulator. It intentionally excludes GUI and real hardware I/O (display/windowing, real controllers, audio output). The goal is to provide a clean, reusable emulator core you can embed into your own frontend, educational tooling, or testing harness.

> [!IMPORTANT]
> This project is a work in progress and may not yet be feature-complete or fully accurate.

## Features

- Accurate 6502 CPU emulation with support for all official/unofficial opcodes and addressing modes.
- Passed the NESTest CPU test ROM for accuracy verification of 6502 instructions and unofficial opcodes.
- Basic PPU implementation for rendering graphics (background and sprites).
- APU stubs for audio processing (full audio implementation is a work in progress).
- Memory mapping that supports various NES mappers (NROM, MMC1, etc.).

## Components

- `nes`: High-level NES struct that ties together CPU, PPU, APU, and memory.
- `cpu`: Implements the 6502 CPU with all instructions, addressing modes, and status flags.
- `ppu`: Handles the NES Picture Processing Unit (PPU) for graphics rendering.
- `apu`: Manages the NES Audio Processing Unit (APU) for sound generation.
- `memory`: Emulates the NES memory map, including RAM, ROM, and I/O registers.
- `rom`: Loads and parses NES ROM files in iNES format.
- `rom/ines`: Loads and parses iNES header.
- `rom/mapper`: Supports various NES cartridge mappers for memory banking.
- `controller`: Emulates NES controller input handling.

## Supported Mappers

- NROM (Mapper 0)

## Getting Started

There are a minimal example of how to use the library to load a ROM and run the emulation loop:

```go
// Import the library
import "github.com/ghosind/go-nes"

// Read a ROM file
romData, err := os.ReadFile("path/to/your/game.nes")

// Create a new NES instance
myNes, err := nes.New(romData)

// Run the emulation loop
for {
    myNes.Step() // Step the CPU, PPU, and APU
    // Handle input, rendering, etc.
}
```

## Testing

The project includes extensive unit tests covering CPU instructions, memory mapping, and ROM loading. To run all tests, use:

```bash
go test ./...
```

To run tests for a specific package, e.g., CPU:

```bash
go test ./cpu -v
```

## Todo List

- [X] 6502 CPU implementation with unofficial opcodes
- [X] Complete NESTest compliance
- [ ] Memory mapping and RAM emulation
- [X] iNES ROM loading and parsing
- [ ] IRQ and NMI handling
- [X] Complete NROM mapper support (mapper 0)
- [ ] PPU implementation for full graphics rendering
- [ ] APU implementation for audio output
- [ ] Support for additional mappers (MMC1, MMC3, etc.)
- [ ] Controller input handling
- [ ] Debugging and logging utilities

## Reference

- [NES Documentation](https://www.nesdev.org/NESDoc.pdf)
- [CPU Reference](https://www.nesdev.org/wiki/CPU_ALL)
- [Obelisk 6502 Guide](https://www.nesdev.org/obelisk-6502-guide/index.html)
- [Unofficial 6502 Opcodes](https://www.nesdev.org/wiki/CPU_unofficial_opcodes)
- [PPU Programmer Reference](https://www.nesdev.org/wiki/PPU_programmer_reference)
- [iNES File Format](https://www.nesdev.org/wiki/INES)
- [INES 1.0 Mapper Grid](https://www.nesdev.org/wiki/Mapper)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
