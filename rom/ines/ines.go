package ines

type NameTableMirror int

const (
	HorizontalMirror NameTableMirror = iota
	VerticalMirror
)

type ConsoleType int

const (
	ConsoleNES ConsoleType = iota
	ConsoleVsUnisystem
	ConsolePlayChoice10
	ConsoleNes2
)

const (
	inesFlag6MirroringMask   = 0x01
	inesFlag6BatteryMask     = 0x02
	inesFlag6TrainerMask     = 0x04
	inesFlag6FourScreenMask  = 0x08
	inesFlag6MapperLowerMask = 0xF0
	inesFlag7VsUnisystemMask = 0x01
	inesFlag7PlayChoiceMask  = 0x02
	inesFlag7Nes2Mask        = 0x0C
	inesFlag7MapperUpperMask = 0xF0
)

type INESHeader struct {
	PrgRomBanks int
	ChrRomBanks int
	PrgRamUnits int
	MapperID    int
	Mirroring   NameTableMirror
	HasBattery  bool
	Trainer     bool
	FourScreen  bool
	ConsoleType ConsoleType
}

func (header *INESHeader) Parse(data []byte) {
	// Get PRG and CHR sizes
	header.PrgRomBanks = int(data[4])
	header.ChrRomBanks = int(data[5])
	header.PrgRamUnits = int(data[8])
	if data[8] == 0 {
		header.PrgRamUnits = 1 // Default to 8KB if zero
	}

	// Get mapper ID
	mapperLower := (data[6] & inesFlag6MapperLowerMask) >> 4
	mapperUpper := (data[7] & inesFlag7MapperUpperMask)
	header.MapperID = int(mapperUpper | mapperLower)

	// Parse flags
	header.parseFlags(data)
}

func (header *INESHeader) parseFlags(data []byte) {
	header.HasBattery = (data[6] & inesFlag6BatteryMask) != 0
	header.Trainer = (data[6] & inesFlag6TrainerMask) != 0
	header.FourScreen = (data[6] & inesFlag6FourScreenMask) != 0
	if (data[6] & inesFlag6MirroringMask) != 0 {
		header.Mirroring = VerticalMirror
	} else {
		header.Mirroring = HorizontalMirror
	}

	header.ConsoleType = header.parseConsoleType(data)
}

func (header *INESHeader) parseConsoleType(data []byte) ConsoleType {
	switch data[7] & (inesFlag7VsUnisystemMask | inesFlag7PlayChoiceMask | inesFlag7Nes2Mask) {
	case 0:
		return ConsoleNES
	case inesFlag7VsUnisystemMask:
		return ConsoleVsUnisystem
	case inesFlag7PlayChoiceMask:
		return ConsolePlayChoice10
	case inesFlag7Nes2Mask:
		return ConsoleNes2
	default:
		return ConsoleNES
	}
}
