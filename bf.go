package gbf

import (
	"io"

	"gbf/memory"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"
var Cls = "\033[2J"

func Exec(mem memory.Memory, src string, in io.Reader, out io.Writer) {
	pos := 0
	for pos < len(src) {
		// l := src[:pos]
		// r := src[pos+1:]
		// fmt.Printf("%s%s%s%s%v%s%s (%3d/%3d)", Cls, White, l, Red, string(src[pos]), White, r, mem.Get(), pos)
		// time.Sleep(5 * time.Second)
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
					// log.Printf("[ %d / %d", pos, brackets)
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
