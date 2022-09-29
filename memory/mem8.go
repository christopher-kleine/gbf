package memory

type Mem8 struct {
	cells []uint8
	index int
}

func NewMem8(size int) Memory {
	return &Mem8{
		cells: make([]uint8, size),
		index: 0,
	}
}

func (mem *Mem8) Add(n int) {
	mem.cells[mem.index] = uint8(int(mem.cells[mem.index]) + n)
}

func (mem *Mem8) Sub(n int) {
	mem.cells[mem.index] = mem.cells[mem.index] - uint8(n)
}

func (mem *Mem8) Set(n int) {
	mem.cells[mem.index] = uint8(n)
}

func (mem *Mem8) Get() int {
	return int(mem.cells[mem.index])
}

func (mem *Mem8) Left(n int) {
	mem.index = mem.index - n
	if mem.index < 0 {
		mem.index = len(mem.cells) - mem.index
	}
}

func (mem *Mem8) Right(n int) {
	mem.index = (mem.index + n) % len(mem.cells)
}
