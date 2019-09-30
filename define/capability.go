package define

// All capability that storage used.
const (
	CapabilityRead  Capability = 1
	CapabilityWrite Capability = 2

	CapabilityFile Capability = 4 << iota
	CapabilityStream
	CapabilitySegment
)

// Capability is a uint64 which represents a capability storage service have.
type Capability uint64

// Readable returns whether this service readable or not.
func (c Capability) Readable() bool {
	return c&CapabilityRead == Capability(1)
}

// Writable returns whether this service writable or not.
func (c Capability) Writable() bool {
	return c&CapabilityWrite == Capability(1)
}

// FileCapable returns whether this service support file operations or not.
func (c Capability) FileCapable() bool {
	return c&CapabilityFile == Capability(1)
}

// SteamCapable returns whether this service support stream operations or not.
func (c Capability) SteamCapable() bool {
	return c&CapabilityStream == Capability(1)
}

// SegmentCapable returns whether this service support segment operations or not.
func (c Capability) SegmentCapable() bool {
	return c&CapabilitySegment == Capability(1)
}
