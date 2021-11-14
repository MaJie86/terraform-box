package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Exec(fileName, dir, commandName string, params []string) error {
	Path := "d://logs/cmd/"
	os.MkdirAll(Path, 755)
	f, err := os.Create(Path + fileName)
	defer f.Close()
	cmd := exec.Command(commandName, params...)
	fmt.Println("CmdAndChangeDirToFile", dir, cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err = cmd.Start()
	if err != nil {
		return err
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		_, err = f.WriteString(line)
		f.Sync()
	}
	_, err = f.WriteString("=================end========================")
	f.Sync()
	err = cmd.Wait()
	return err
}

func ReadLog(lineNumber int) ([]string, int) {
	file, _ := os.Open("d://logs/cmd/test.log")
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	var lines []string
	for fileScanner.Scan() {
		if lineCount >= lineNumber {
			lines = append(lines, fileScanner.Text())
		}
		lineCount++
	}
	defer file.Close()
	return lines, lineCount - 1
}
