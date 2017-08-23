package  dirindex

import "os"
import "io/ioutil"
import "testing"

func TestFormatStream(t *testing.T) {
	fileInfos,err:=ioutil.ReadDir(".")
	if err!=nil{panic(err)}
	TagEncode(os.Stdout,fileInfos,NameOnly)
	TagEncode(os.Stdout,fileInfos,NameSizeModTime)
	TagEncode(os.Stdout,fileInfos,NameSizeModTimeMode)
	TagEncode(os.Stdout,fileInfos,NameTypeSizeModTimeMode)
	TagEncode(os.Stdout,fileInfos,NameOnly,false)
}

func TestFormatting(t *testing.T) {
	f,err:=os.Create("index.xml") 
	if err!=nil{panic(err)}
	XMLEncode(f,".",NameSizeModTime)
}


