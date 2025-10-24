

# go-nes

![test](https://github.com/ghosind/go-nes/workflows/test/badge.svg)
[![go report](https://goreportcard.com/badge/github.com/ghosind/go-nes)](https://goreportcard.com/report/github.com/ghosind/go-nes)
[![codecov](https://codecov.io/gh/ghosind/go-nes/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/go-nes)
![license](https://img.shields.io/github/license/ghosind/go-nes)

`go-nes`是一个经过充分测试的 NES 核心模拟库，用于实现 NES（任天堂红白机）模拟器的核心逻辑。项目有意不包含图形界面或对真实硬件的直接 I/O（例如窗口显示、真实手柄、音频输出），目的是提供一个干净、可复用的模拟器核心，方便你在自己的前端、教学工具或测试环境中集成。

> [!IMPORTANT]
> 该项目仍在开发中，可能尚未覆盖所有特性或完全准确。

## 功能一览

- 精准的 6502 CPU 仿真，支持官方与常见的非官方指令及所有寻址模式。
- 已通过 NESTest 的 CPU 测试 ROM，可用于验证实现的准确性。
- 提供基础的 PPU（图像处理单元）实现，用于渲染背景与精灵（sprite）。
- 提供 APU（音频处理单元）的框架，完整音频功能尚未完成。
- 内存映射支持常见卡带映射器（Mapper），目前包含 NROM 等实现。

## 项目结构（主要模块）

- `nes`：顶层的 NES 结构体，负责将 CPU、PPU、APU 与内存组合在一起。
- `cpu`：6502 CPU 的实现，包含所有指令、寻址模式与状态寄存器。
- `ppu`：PPU（图像处理单元）实现，处理图像渲染逻辑。
- `apu`：APU（音频处理单元）骨架，用于后续音频功能扩展。
- `memory`：模拟 NES 的内存映射，包括 RAM、ROM 与 I/O 寄存器。
- `rom`：用于加载并解析 iNES 格式的 ROM 文件。
- `rom/ines`：解析 iNES 头信息的小模块。
- `rom/mapper`：实现卡带映射器以支持内存银行切换。

## 当前支持的映射器

- NROM（Mapper 0）

## 快速上手

下面给出一个最小化的使用示例，演示如何加载 ROM 并启动仿真主循环：

```go
import (
	"os"
	"github.com/ghosind/go-nes"
)

// 读取 ROM 文件
romData, err := os.ReadFile("path/to/your/game.nes")
if err != nil {
	// 处理错误
}

// 创建 NES 实例
myNes, err := nes.New(romData)
if err != nil {
	// 处理错误
}

// 主循环（仅示例）
for {
	myNes.Step() // 推进 CPU、PPU 和 APU 的周期
	// 在此处处理输入、渲染画面或保存状态等
}
```

## 运行测试

仓库包含大量单元测试，覆盖 CPU 指令、内存映射与 ROM 加载等。运行全部测试：

```bash
go test ./...
```

运行某个包的测试（例如 `cpu` 包）：

```bash
go test ./cpu -v
```

## 开发计划（Todo）

- [X] 6502 CPU 实现（包含非官方指令）
- [X] 通过 NESTest 验证 CPU 行为
- [ ] 内存映射与 RAM 的补全
- [X] 支持 iNES ROM 的加载与解析
- [ ] IRQ 与 NMI 中断处理
- [X] 完整实现 NROM（Mapper 0）
- [ ] 完整实现 PPU 以支持全部渲染特性
- [ ] 完整实现 APU 以输出音频
- [ ] 支持更多映射器（如 MMC1、MMC3 等）
- [ ] 控制器输入（手柄）支持
- [ ] 更完善的调试与日志工具

## 参考资料

- [NES Documentation](https://www.nesdev.org/NESDoc.pdf)
- [CPU Reference](https://www.nesdev.org/wiki/CPU_ALL)
- [Obelisk 6502 Guide](https://www.nesdev.org/obelisk-6502-guide/index.html)
- [Unofficial 6502 Opcodes](https://www.nesdev.org/wiki/CPU_unofficial_opcodes)
- [PPU Programmer Reference](https://www.nesdev.org/wiki/PPU_programmer_reference)
- [iNES File Format](https://www.nesdev.org/wiki/INES)
- [INES 1.0 Mapper Grid](https://www.nesdev.org/wiki/Mapper)

## 许可证

本项目采用 MIT 许可证，详见仓库根目录的 `LICENSE` 文件。
