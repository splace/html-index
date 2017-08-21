package dirlisting

import "bytes"

func escape(u string) string{
	for i,c:=range(u){
		switch c {
		case '"','\'','&','<','>':
			buf := bytes.NewBufferString(u[:i])
			for _,c:=range(u[i:]){
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



