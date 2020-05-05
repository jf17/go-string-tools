package stringtools

import "fmt"

func ConvertWindows1251ToUTF8(ba []uint8) string  {
	dec := charmap.Windows1251.NewDecoder()
	out, _ := dec.Bytes(ba)
	return string(out)
}

