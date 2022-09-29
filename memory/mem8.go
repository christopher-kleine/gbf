package memory

type Mem8 struct {
	cells []uint8
	index int
	clamp bool
}

func NewMem8(size int, clamp bool) Memory {
	return &Mem8{
		cells: make([]uint8, size),
		index: 0,
		clamp: clamp,
	}
}

func (mem *Mem8) Add(n int) {
	if mem.clamp {
		mem.cells[mem.index] = uint8(clamp(1<<8-1, int(mem.cells[mem.index])+n))
	} else {
		mem.cells[mem.index] = uint8(int(mem.cells[mem.index]) + n)
	}
}

func (mem *Mem8) Sub(n int) {
	if mem.clamp {
		mem.cells[mem.index] = uint8(clamp(1<<8-1, int(mem.cells[mem.index])-n))
	} else {
		mem.cells[mem.index] = mem.cells[mem.index] - uint8(n)
	}
}

func (mem *Mem8) Set(n int) {
	mem.cells[mem.index] = uint8(n)
}

func (mem *Mem8) Get() int {
	return int(mem.cells[mem.index])
}

func (mem *Mem8) Left(n int) {
	mem.index = wrap(len(mem.cells)-1, mem.index-n)
}

func (mem *Mem8) Right(n int) {
	mem.index = wrap(len(mem.cells)-1, mem.index+n)
}
