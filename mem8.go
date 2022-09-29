package gbf

type Mem8 struct {
	cells []uint8
}

func NewMem8(size int) Memory {
	return &Mem8{
		cells: make([]uint8, size),
	}
}

func (mem *Mem8) Add(n int) {
}

func (mem *Mem8) Sub(n int) {
}

func (mem *Mem8) Set(n int) {
}

func (mem *Mem8) Get() int {
	return 0
}

func (mem *Mem8) Left(n int) {
}

func (mem *Mem8) Right(n int) {
}
