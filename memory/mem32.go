package memory

type Mem32 struct {
	cells []uint32
	index int
	clamp bool
}

func NewMem32(size int, clamp bool) Memory {
	return &Mem32{
		cells: make([]uint32, size),
		index: 0,
		clamp: clamp,
	}
}

func (mem *Mem32) Add(n int) {
	if mem.clamp {
		mem.cells[mem.index] = uint32(clamp(1<<32-1, int(mem.cells[mem.index])+n))
	} else {
		mem.cells[mem.index] = uint32(int(mem.cells[mem.index]) + n)
	}
}

func (mem *Mem32) Sub(n int) {
	if mem.clamp {
		mem.cells[mem.index] = uint32(clamp(1<<32-1, int(mem.cells[mem.index])-n))
	} else {
		mem.cells[mem.index] = mem.cells[mem.index] - uint32(n)
	}
}

func (mem *Mem32) Set(n int) {
	mem.cells[mem.index] = uint32(n)
}

func (mem *Mem32) Get() int {
	return int(mem.cells[mem.index])
}

func (mem *Mem32) Left(n int) {
	mem.index = mem.index - n
	if mem.index < 0 {
		mem.index = len(mem.cells) - mem.index
	}
}

func (mem *Mem32) Right(n int) {
	mem.index = (mem.index + n) % len(mem.cells)
}
