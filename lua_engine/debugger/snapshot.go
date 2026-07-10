package debugger

func (d *Debugger) captureFrame(pos Position) *Frame {
	return &Frame{
		ID:       1,
		Name:     "lua",
		Position: pos,
	}
}
