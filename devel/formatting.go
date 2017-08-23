package dirindex

import "os"
import "io"
import "fmt"
import "time"
import "path"

const (
	NameOnly = iota
	NameSizeModTime
	NameSizeModTimeMode
	NameTypeSizeModTimeMode
)

// make a hierarchy of formatting, by making separate stringers, all embedding the source, for the different compositions of information.
// this hierarchy fits XML structure, and means element id's are all just strings.   

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


func XMLEncode(w io.WriteCloser, dir string,details uint) (err error){
	folder,err:=os.Open(dir) 
	if err!=nil{return}
	fileInfos,err:=folder.Readdir(0)
	if err!=nil{return}
	fmt.Fprint(w,"<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n")
	fmt.Fprint(w,"<?xml-stylesheet type=\"text/xsl\" href=\"index.xsl\"?>\n")
	fmt.Fprint(w,"<index host=\""+ignoreError(os.Hostname)+"\" name=\""+path.Join(ignoreError(os.Getwd),folder.Name())+"\" >\n")
	TagEncode(w,fileInfos,details)
	fmt.Fprint(w,"</index>\n")
	w.Close()
	return
}

func ignoreError(fn func ()(string,error)) string{
	r,_:=fn()
	return r
}

func TagEncode(w io.WriteCloser, fis []os.FileInfo,details uint) {
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
			// TODO retrieve type maybe from xattribs?
			if !fi.IsDir() {fmt.Fprintf(w,FileFormatting, AttribTypedFileExtendedInfo{"application/octet-stream",fi})}
		}
	}
}

