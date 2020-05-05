package stringtools

import "golang.org/x/text/encoding/charmap"

func ConvertWindows1251ToUTF8(str string) string  {
	byteArr := []byte(str)
	dec := charmap.Windows1251.NewDecoder()
	out, _ := dec.Bytes(byteArr)
	return string(out)
}

