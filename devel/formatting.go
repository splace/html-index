package dirlisting

import "os"
import "io"
import "fmt"
import "time"

const (
	NameOnly = iota
	NameSizeModTime
	NameSizeModTimeMode
	NameTypeSizeModTimeMode
)

var FileFormatting = "\t<txt %s/>\n"
var DirFormatting = "\t<dir %s/>\n"
var DefaultFormatting = "\t<raw %s/>\n"

type AttribNameInfo struct {
	os.FileInfo
}

func (fi AttribNameInfo) String() string {
	return fmt.Sprintf("name=\"%s\"", escape(fi.Name()))
}

type AttribFileInfo struct {
	os.FileInfo
}

func (fi AttribFileInfo) String() string {
	return fmt.Sprintf("name=\"%s\" size=\"%d\" modified=\"%s\"", escape(fi.Name()), fi.Size(), fi.ModTime().Format(time.RFC3339))
}

type AttribDirInfo struct {
	os.FileInfo
}

func (fi AttribDirInfo) String() string {
	return fmt.Sprintf("name=\"%s\" modified=\"%s\"", escape(fi.Name()), fi.ModTime().Format(time.RFC3339))
}

type AttribFileExtendedInfo struct {
	os.FileInfo
}

func (fi AttribFileExtendedInfo) String() string {
	return fmt.Sprintf("name=\"%s\" size=\"%d\" modified=\"%s\" mode=\"%s\"", escape(fi.Name()), fi.Size(), fi.ModTime().Format(time.RFC3339), fi.Mode())
}

type AttribTypedFileExtendedInfo struct {
	mime string
	os.FileInfo
}

func (fi AttribTypedFileExtendedInfo) String() string {
	return fmt.Sprintf("name=\"%s\" type=\"%s\" size=\"%d\" modified=\"%s\" mode=\"%s\"", escape(fi.Name()),fi.mime, fi.Size(), fi.ModTime().Format(time.RFC3339), fi.Mode())
}

type AttribDirExtendedInfo struct {
	os.FileInfo
}

func (fi AttribDirExtendedInfo) String() string {
	return fmt.Sprintf("name=\"%s\" modified=\"%s\" mode=\"%s\"", escape(fi.Name()), fi.ModTime().Format(time.RFC3339), fi.Mode())
}

func XMLEncode(fis []os.FileInfo, details uint) (r io.Reader) {
	r, w := io.Pipe()
	go func() {
		switch details {
		case NameOnly:
			for _, fi := range fis {
				if fi.IsDir() {fmt.Fprintf(w,DirFormatting, AttribNameInfo{fi})}
			}
			for _, fi := range fis {
				if !fi.IsDir() {fmt.Fprintf(w,FileFormatting, AttribNameInfo{fi})}
			}

		case NameSizeModTime:
			for _, fi := range fis {
				if fi.IsDir() {fmt.Fprintf(w,DirFormatting, AttribDirInfo{fi})}
			}
			for _, fi := range fis {
				if !fi.IsDir() {fmt.Fprintf(w,FileFormatting, AttribFileInfo{fi})}
			}
		case NameSizeModTimeMode:
			for _, fi := range fis {
				if fi.IsDir() {fmt.Fprintf(w,DirFormatting, AttribDirExtendedInfo{fi})}
			}
			for _, fi := range fis {
				if !fi.IsDir() {fmt.Fprintf(w,FileFormatting, AttribFileExtendedInfo{fi})}
			}
		case NameTypeSizeModTimeMode:
			for _, fi := range fis {
				if fi.IsDir() {fmt.Fprintf(w,DirFormatting, AttribDirExtendedInfo{fi})}
			}
			for _, fi := range fis {
				// TODO retrieve type from xattribs?
				if !fi.IsDir() {fmt.Fprintf(w,FileFormatting, AttribTypedFileExtendedInfo{"application/octet-stream",fi})}
			}
		}
		w.Close()
	}()
	return
}


