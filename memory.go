package gbf

type Memory interface {
	Add(int)
	Sub(int)
	Set(int)
	Get() int
	Left(int)
	Right(int)
}
