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

// make a hierarchy of formatting, by using separate stringers, all embedding FileInfo, for the different arrangements of information.
// a hierarchy fits XML structure, and means element id's are all just strings.   

var FileFormatting = "\t<txt %s/>\n"
var DirFormatting = "\t<dir %s/>\n"
var DefaultFormatting = "\t<raw %s/>\n"

type NameStringer struct {
	os.FileInfo
}

func (fi NameStringer) String() string {
	return fmt.Sprintf("name=\"%s\"", escape(fi.Name()))
}

type FileNameSizeModTimeStringer struct {
	os.FileInfo
}

func (fi FileNameSizeModTimeStringer) String() string {
	return fmt.Sprintf("name=\"%s\" size=\"%d\" modified=\"%s\"", escape(fi.Name()), fi.Size(), fi.ModTime().Format(time.RFC3339))
}

type DirNameSizeModTimeStringer struct {
	os.FileInfo
}

func (fi DirNameSizeModTimeStringer) String() string {
	return fmt.Sprintf("name=\"%s\" modified=\"%s\"", escape(fi.Name()), fi.ModTime().Format(time.RFC3339))
}

type FileNameSizeModTimeModeStringer struct {
	os.FileInfo
}

func (fi FileNameSizeModTimeModeStringer) String() string {
	return fmt.Sprintf("name=\"%s\" size=\"%d\" modified=\"%s\" mode=\"%s\"", escape(fi.Name()), fi.Size(), fi.ModTime().Format(time.RFC3339), fi.Mode())
}

type FileNameTypeSizeModTimeModeStringer struct {
	mime string
	os.FileInfo
}

func (fi FileNameTypeSizeModTimeModeStringer) String() string {
	return fmt.Sprintf("name=\"%s\" type=\"%s\" size=\"%d\" modified=\"%s\" mode=\"%s\"", escape(fi.Name()),fi.mime, fi.Size(), fi.ModTime().Format(time.RFC3339), fi.Mode())
}

type DirNameSizeModTimeModeStringer struct {
	os.FileInfo
}

func (fi DirNameSizeModTimeModeStringer) String() string {
	return fmt.Sprintf("name=\"%s\" modified=\"%s\" mode=\"%s\"", escape(fi.Name()), fi.ModTime().Format(time.RFC3339), fi.Mode())
}

// WriteXML writes the directory listing as XML, using the indicated Tag writer.
func WriteXML(w io.WriteCloser, dir string,tagWriter func(io.WriteCloser, []os.FileInfo)) (err error){
	folder,err:=os.Open(dir) 
	if err!=nil{return}
	fileInfos,err:=folder.Readdir(0)
	if err!=nil{return}
	fmt.Fprint(w,"<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n")
	fmt.Fprint(w,"<?xml-stylesheet type=\"text/xsl\" href=\"index.xsl\"?>\n")
	fmt.Fprint(w,"<index host=\""+ignoreError(os.Hostname)+"\" name=\""+path.Join(ignoreError(os.Getwd),folder.Name())+"\" >\n")
	tagWriter(w,fileInfos)
	fmt.Fprint(w,"</index>\n")
	w.Close()
	return
}

func ignoreError(fn func ()(string,error)) string{
	r,_:=fn()
	return r
}

// create a closure for a specific tag structure
func TagWriter(details uint,DirFirst ...bool) func(io.WriteCloser, []os.FileInfo){
	return func(w io.WriteCloser, fis []os.FileInfo){ WriteTags(w, fis,details,DirFirst...)}
}

// WriteTags writes the directory listing as XML Tags, with the required details, and optionally folders first.
func WriteTags(w io.WriteCloser, fis []os.FileInfo,details uint,dirFirstFlag ...bool) {
	switch details {
	case NameOnly:
		if len(dirFirstFlag)==0 || dirFirstFlag[0]{
			for _, fi := range fis {
				if fi.IsDir() {fmt.Fprintf(w,DirFormatting, NameStringer{fi})}
			}
			for _, fi := range fis {
				if fi.Mode().IsRegular() {fmt.Fprintf(w,FileFormatting, NameStringer{fi})}
			}
		}else{
			for _, fi := range fis {
				if fi.IsDir(){
					fmt.Fprintf(w,DirFormatting, NameStringer{fi})
				}else if fi.Mode().IsRegular() {
					fmt.Fprintf(w,FileFormatting, NameStringer{fi})
				} 
			}
		}
	case NameSizeModTime:
		if len(dirFirstFlag)==0 || dirFirstFlag[0]{
			for _, fi := range fis {
				if fi.IsDir() {fmt.Fprintf(w,DirFormatting, DirNameSizeModTimeStringer{fi})}
			}
			for _, fi := range fis {
				if fi.Mode().IsRegular() {fmt.Fprintf(w,FileFormatting, FileNameSizeModTimeStringer{fi})}
			}
		}else{
			for _, fi := range fis {
				if fi.IsDir(){
					fmt.Fprintf(w,DirFormatting,DirNameSizeModTimeStringer{fi})
				}else if fi.Mode().IsRegular() {
					fmt.Fprintf(w,FileFormatting,FileNameSizeModTimeStringer{fi})
				}
			}
		}
	case NameSizeModTimeMode:
		if len(dirFirstFlag)==0 || dirFirstFlag[0]{
			for _, fi := range fis {
				if fi.IsDir() {fmt.Fprintf(w,DirFormatting, DirNameSizeModTimeModeStringer{fi})}
			}
			for _, fi := range fis {
				if fi.Mode().IsRegular() {fmt.Fprintf(w,FileFormatting, FileNameSizeModTimeModeStringer{fi})}
			}
		}else{
			for _, fi := range fis {
				if fi.IsDir(){
					fmt.Fprintf(w,DirFormatting, DirNameSizeModTimeModeStringer{fi})
				}else if fi.Mode().IsRegular() {
					fmt.Fprintf(w,FileFormatting, FileNameSizeModTimeModeStringer{fi})
				} 
			}
		}
	case NameTypeSizeModTimeMode:
		if len(dirFirstFlag)==0 || dirFirstFlag[0]{
			for _, fi := range fis {
				if fi.IsDir() {fmt.Fprintf(w,DirFormatting, DirNameSizeModTimeModeStringer{fi})}
			}
			for _, fi := range fis {
				// TODO retrieve actual type,(and so weather Text) maybe from xattribs?
				//if fi.Mode().IsRegular() {fmt.Fprintf(w,FileFormatting, FileNameTypeSizeModTimeModeStringer{"text/*",fi})}
				if fi.Mode().IsRegular() {fmt.Fprintf(w,DefaultFormatting, FileNameTypeSizeModTimeModeStringer{"application/octet-stream",fi})}
			}
		}else{
			for _, fi := range fis {
				if fi.IsDir(){
					fmt.Fprintf(w,DirFormatting, DirNameSizeModTimeModeStringer{fi})
				}else if fi.Mode().IsRegular() {
					//fmt.Fprintf(w,FileFormatting, FileNameTypeSizeModTimeModeStringer{"text/*",fi})
					fmt.Fprintf(w,DefaultFormatting, FileNameTypeSizeModTimeModeStringer{"application/octet-stream",fi})
				}
			}
		}
	}
}

