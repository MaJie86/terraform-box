package cmd

import (
	"bufio"
	"fmt"
	"github.com/majie86/terraform-box/taskpool"
	"io"
	"os"
	"os/exec"
)

func Exec(fileName, dir, commandName string, params []string, task *taskpool.Task) error {
	os.MkdirAll(dir, 755)
	f, err := os.Create(dir + fileName)
	defer f.Close()
	cmd := exec.Command(commandName, params...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err = cmd.Start()
	task.Command = cmd
	if err != nil {
		return err
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
	return err
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
