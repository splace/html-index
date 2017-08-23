package  dirindex

import "os"
import "io/ioutil"
import "testing"

func TestFormatStream(t *testing.T) {
	fileInfos,err:=ioutil.ReadDir(".")
	if err!=nil{panic(err)}
	WriteTags(os.Stdout,fileInfos,NameOnly)
	WriteTags(os.Stdout,fileInfos,NameSizeModTime)
	WriteTags(os.Stdout,fileInfos,NameSizeModTimeMode)
	WriteTags(os.Stdout,fileInfos,NameTypeSizeModTimeMode)
	WriteTags(os.Stdout,fileInfos,NameOnly,false)
}

func TestFormatting(t *testing.T) {
	f,err:=os.Create("index.xml") 
	if err!=nil{panic(err)}
	WriteXML(f,".",TagWriter(NameSizeModTime))
}


