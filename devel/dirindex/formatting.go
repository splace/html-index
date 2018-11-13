package dirindex

import "os"
import "io"
import "fmt"
import "time"
import "path"

type details uint8

const (
	NameOnly details = iota
	NameSizeModTime
	NameSizeModTimeMode
	NameTypeSizeModTimeMode
)

// format by using stringers embedding os.FileInfo, each for the different arrangements of information for a file.
// hierarically use a stringer inside a stringer to only have one underlying stringer for each type.
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

// WriteXML writes formatted XML to the provided writer.
// the provided Tag writing function formats each file in the provided directory to an XML Tag.
func WriteXML(w io.Writer, dir string, tagWriter func(io.Writer, []os.FileInfo)) (err error){
	folder,err:=os.Open(dir) 
	if err!=nil{return}
	fileInfos,err:=folder.Readdir(0)
	if err!=nil{return}
	fmt.Fprint(w,"<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n")
	fmt.Fprint(w,"<?xml-stylesheet type=\"text/xsl\" href=\"index.xsl\"?>\n")
	fmt.Fprint(w,"<index host=\""+ignoreError(os.Hostname)+"\" name=\""+path.Join(ignoreError(os.Getwd),folder.Name())+"\" >\n")
	tagWriter(w,fileInfos)
	fmt.Fprint(w,"</index>\n")
	return
}

func ignoreError(fn func ()(string,error)) string{
	r,_:=fn()
	return r
}

// create a closure for a specific tag structure.
func TagWriter(d details, DirFirst ...bool) func(io.Writer, []os.FileInfo){
	return func(w io.Writer, fis []os.FileInfo){ WriteTags(w, fis, d, DirFirst...)}
}

// WriteTags writes the directory listing as XML Tags, with the required details, and optionally folders first.
func WriteTags(w io.Writer, fis []os.FileInfo,d details,dirFirstFlag ...bool) {
	switch d {
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
				// TODO retrieve actual type,(and so if Text) maybe from xattribs?
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

