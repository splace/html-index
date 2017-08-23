package dirlisting

import "os"
import "io"
import "testing"
import "fmt"
import "path"

func TestFlexFormatting(t *testing.T) {
	folder,_:=os.Open(".") 
	fileInfos,_:=folder.Readdir(0)
	io.Copy(os.Stdout,XMLEncode(fileInfos,NameOnly))
	io.Copy(os.Stdout,XMLEncode(fileInfos,NameSizeModTime))
	io.Copy(os.Stdout,XMLEncode(fileInfos,NameSizeModTimeMode))
	io.Copy(os.Stdout,XMLEncode(fileInfos,NameTypeSizeModTimeMode))
}

func ignoreError(fn func ()(string,error)) string{
	r,_:=fn()
	return r
}

func TestStringFormatting(t *testing.T) {
	folder,err:=os.Open(".") 
	if err!=nil{panic(err)}
	f,err:=os.Create("index.xml") 
	if err!=nil{panic(err)}
	fileInfos,err:=folder.Readdir(0)
	if err!=nil{panic(err)}
	fmt.Fprint(f,"<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n")
	fmt.Fprint(f,"<?xml-stylesheet type=\"text/xsl\" href=\"index.xsl\"?>\n")
	fmt.Fprint(f,"<index host=\""+ignoreError(os.Hostname)+"\" name=\""+path.Join(ignoreError(os.Getwd),folder.Name())+"\" >\n")
	io.Copy(f,XMLEncode(fileInfos,NameSizeModTime))
	fmt.Fprint(f,"</index>\n")
}


