package main

import (
	"fmt"
	"golangDemo/dirTest/logger"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	log.Println(GetCurrentDirectory())
	log.Println(GetRootDir())
	log.Println(GetExecFilePath())

	logger.InitLog("error", "./a.json")
	logger.Debug("hello")
	logger.Error("err")
}
func GetCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

func GetRootDir() string {
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file = fmt.Sprintf("%s%s", file, string(os.PathSeparator))
	}
	return file
}

func GetExecFilePath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file, _ = filepath.Abs(file)
	}
	return file
}
