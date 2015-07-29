package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	. "github.com/colefan/gsgo/tools/packettool"
	"github.com/colefan/gsgo/tools/utils"
)

func main() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	splitstrings := strings.Split(path, string(os.PathSeparator))
	size := len(splitstrings)
	splitstrings = strings.Split(path, splitstrings[size-1])

	xmlPath := splitstrings[0] + ".." + string(os.PathSeparator)
	storePath := xmlPath
	//	fmt.Println("xmlpath: " + xmlPath)
	//	fmt.Println("storepath: " + storePath)

	if len(xmlPath) < 0 {
		fmt.Println("protocol_desc_file is empty:(")
		return
	}

	if len(storePath) < 0 {
		fmt.Println("packet_store_path is empty:(")
		return
	}

	files, err := toolsutils.ListDir(xmlPath, ".xml")
	if err != nil {
		fmt.Println("xmlpath read error:", err)
		return
	}
	//fmt.Println("files,", len(files), "xmlpath=", xmlPath)
	for _, filepath := range files {
		//遍历目录，
		fmt.Println("storepath:", storePath)
		fmt.Println("filepath :", filepath)
		if perr := GenPacketForGoLang(storePath, filepath); perr != nil {
			fmt.Println("GenPacketForGoLang error,error:", perr)
		}
	}

}
