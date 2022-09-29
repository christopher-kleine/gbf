package memory

type Mem16 struct {
	cells []uint16
	index int
	clamp bool
}

func NewMem16(size int, clamp bool) Memory {
	return &Mem16{
		cells: make([]uint16, size),
		index: 0,
		clamp: clamp,
	}
}

func (mem *Mem16) Add(n int) {
	if mem.clamp {
		mem.cells[mem.index] = uint16(clamp(1<<16-1, int(mem.cells[mem.index])+n))
	} else {
		mem.cells[mem.index] = uint16(int(mem.cells[mem.index]) + n)
	}
}

func (mem *Mem16) Sub(n int) {
	if mem.clamp {
		mem.cells[mem.index] = uint16(clamp(1<<16-1, int(mem.cells[mem.index])-n))
	} else {
		mem.cells[mem.index] = mem.cells[mem.index] - uint16(n)
	}
}

func (mem *Mem16) Set(n int) {
	mem.cells[mem.index] = uint16(n)
}

func (mem *Mem16) Get() int {
	return int(mem.cells[mem.index])
}

func (mem *Mem16) Left(n int) {
	mem.index = mem.index - n
	if mem.index < 0 {
		mem.index = len(mem.cells) - mem.index
	}
}

func (mem *Mem16) Right(n int) {
	mem.index = (mem.index + n) % len(mem.cells)
}
