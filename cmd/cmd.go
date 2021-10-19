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
	f, err := os.Create(Path + fileName) //创建文件
	defer f.Close()
	cmd := exec.Command(commandName, params...)
	fmt.Println("CmdAndChangeDirToFile", dir, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
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
