package bits

type Bits uint32

//SetBit sets a bit in the bitset
func SetBit(bitset, flag Bits) Bits { return bitset | flag }

//ClearBit clears a flag in the bitset
func ClearBit(bitset, flag Bits) Bits { return bitset &^ flag }

//ToggleBit toggles a flag is the bitset
func ToggleBit(bitset, flag Bits) Bits { return bitset ^ flag }

//HasBit checks if the bitset has a flag set
func HasBit(bitset, flag Bits) bool { return bitset&flag != 0 }
