package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Exec(fileName, dir, commandName string, params []string) (*exec.Cmd, error) {
	os.MkdirAll(dir, 755)
	f, err := os.Create(dir + fileName)
	defer f.Close()
	cmd := exec.Command(commandName, params...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe: ", err)
		return cmd, err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err = cmd.Start()
	if err != nil {
		return cmd, err
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		print(line)
		_, err = f.WriteString(line)
		f.Sync()
	}
	err = cmd.Wait()
	return cmd, err
}

func ReadLog(filePath string, lineNumber int) ([]string, int) {
	file, _ := os.Open(filePath)
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
