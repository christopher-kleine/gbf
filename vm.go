package gbf

import (
	"io"

	"gbf/memory"
)

type Instruction struct {
	Token byte
	Value int
}

func VM(mem memory.Memory, src []byte, in io.Reader, out io.Writer) {
	var (
		allIns  []Instruction
		currIns = Instruction{
			Token: src[0],
		}
	)

	pos := 0
	for pos < len(src) {
		token := src[pos]

		switch token {
		case '+':
			if currIns.Token != token {
				allIns = append(allIns, currIns)
				currIns = Instruction{
					Token: '+',
					Value: 0,
				}
			}
			currIns.Value++

		case '-':
			if currIns.Token != token {
				allIns = append(allIns, currIns)
				currIns = Instruction{
					Token: '-',
					Value: 0,
				}
			}
			currIns.Value++

		case '>':
			if currIns.Token != token {
				allIns = append(allIns, currIns)
				currIns = Instruction{
					Token: '>',
					Value: 0,
				}
			}
			currIns.Value++

		case '<':
			if currIns.Token != token {
				allIns = append(allIns, currIns)
				currIns = Instruction{
					Token: '<',
					Value: 0,
				}
			}
			currIns.Value++

		case '.':
			if currIns.Token != token {
				allIns = append(allIns, currIns)
				currIns = Instruction{
					Token: '.',
					Value: 0,
				}
			}
			currIns.Value++

		case ',':
			if currIns.Token != token {
				allIns = append(allIns, currIns)
				currIns = Instruction{
					Token: ',',
					Value: 0,
				}
			}
			currIns.Value++

		case '[':
			allIns = append(allIns, currIns)
			currIns = Instruction{
				Token: '[',
				Value: 0,
			}
			if d := string(src[pos : pos+3]); d == "[-]" || d == "[+]" {
				currIns = Instruction{
					Token: '#',
					Value: 0,
				}
				pos += 2
			} else if d := string(src[pos : pos+3]); d == "[>]" {
				currIns = Instruction{
					Token: '\\',
					Value: 0,
				}
				pos += 2
			} else if d := string(src[pos : pos+3]); d == "[<]" {
				currIns = Instruction{
					Token: '/',
					Value: 0,
				}
				pos += 2
			}

		case ']':
			allIns = append(allIns, currIns)
			currIns = Instruction{
				Token: ']',
				Value: 0,
			}
			brackets := 1
			pos := len(allIns)
			for brackets > 0 {
				pos--
				currIns.Value++
				if allIns[pos].Token == ']' {
					brackets++
				}
				if allIns[pos].Token == '[' {
					brackets--
				}
			}
			allIns[pos].Value = currIns.Value
		}

		pos++
	}
	allIns = append(allIns, currIns)

	pos = 0
	for pos < len(allIns) {
		switch allIns[pos].Token {
		case '+':
			mem.Add(allIns[pos].Value)

		case '-':
			mem.Sub(allIns[pos].Value)

		case '>':
			mem.Right(allIns[pos].Value)

		case '<':
			mem.Left(allIns[pos].Value)

		case '#':
			mem.Set(0)

		case '/':
			for mem.Get() != 0 {
				mem.Left(1)
			}

		case '\\':
			for mem.Get() != 0 {
				mem.Right(1)
			}

		case '[':
			if mem.Get() == 0 {
				pos += allIns[pos].Value
			}

		case ']':
			if mem.Get() != 0 {
				pos -= allIns[pos].Value
			}

		case '.':
			for index := 0; index < allIns[pos].Value; index++ {
				out.Write([]byte{byte(mem.Get())})
			}
		}
		pos++
	}
}
