package gbf

import (
	"io"

	"gbf/memory"
)

func Eval(mem memory.Memory, src []byte, in io.Reader, out io.Writer) {
	pos := 0
	for pos < len(src) {
		switch src[pos] {
		case '+':
			mem.Add(1)
			pos++

		case '-':
			mem.Sub(1)
			pos++

		case '>':
			mem.Right(1)
			pos++

		case '<':
			mem.Left(1)
			pos++

		case '[':
			if mem.Get() == 0 {
				brackets := 1
				pos++
				for pos < len(src) && brackets > 0 {
					if src[pos] == '[' {
						brackets++
					}

					if src[pos] == ']' {
						brackets--
					}

					pos++
				}
				pos--
			}
			pos++

		case ']':
			if mem.Get() != 0 {
				brackets := 1
				pos--
				for pos < len(src) && brackets > 0 {
					if src[pos] == ']' {
						brackets++
					}

					if src[pos] == '[' {
						brackets--
					}

					pos--
				}
			}
			pos++

		case '.':
			out.Write([]byte{byte(mem.Get())})
			pos++

		case ',':
			b := make([]byte, 1)
			in.Read(b)
			mem.Set(int(b[0]))
			pos++

		default:
			pos++
		}
	}
}
