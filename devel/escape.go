package dirlisting

import "bytes"

func escape(u string) string{
	for _,c:=range(u){
		switch c {
		case '"','\'','&','<','>':
			buf := new(bytes.Buffer)
			for _,c:=range(u){
				switch c {
				case '"':
					buf.WriteString("&quot;")
				case '\'':
					buf.WriteString("&apos;")
				case '&':
					buf.WriteString("&amp;")
				case '<':
					buf.WriteString("&lt;")
				case '>':
					buf.WriteString("&gt;")
				default:
					buf.WriteRune(c)
				}
			}
			return buf.String()
		}
	}
	return u   // no escapeing required
}



