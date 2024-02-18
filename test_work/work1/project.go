package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

/*

工作1：将代码运行

工作2：请结合整个目录文件以及输出来理解代码的功能

工作3：给代码加上注释

注意：你可以进行百度了解某个函数的功能
*/

var (
	fileCantOpen error = errors.New("文件打开错误")
	fileReadErr  error = errors.New("文件读取内容错误或两个文件行数不匹配")
)

func merge(path1 string, path2 string, path3 string) error {
	file1, err1 := os.OpenFile(path1, os.O_RDONLY, 0666)
	file2, err2 := os.OpenFile(path2, os.O_RDONLY, 0666)
	file3, err3 := os.OpenFile(path3, os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil || err2 != nil || err3 != nil {
		return fileCantOpen
	}
	defer file1.Close()
	defer file2.Close()
	defer file3.Close()
	reader1 := bufio.NewReader(file1)
	reader2 := bufio.NewReader(file2)
	writer := bufio.NewWriter(file3)
	for {
		line1, err1 := reader1.ReadString('\n')
		line2, err2 := reader2.ReadString('\n')
		if err1 != nil && err2 != nil && err1.Error() == "EOF" && err2.Error() == "EOF" {
			break
		}
		if err1 != nil || err2 != nil {
			return fileReadErr
		}
		writer.WriteString(line2)
		writer.WriteString(line1)
	}
	writer.Flush()
	return nil
}

func main() {
	// 测试
	path1 := "test_work/work1/chinese.txt"
	path2 := "test_work/work1/english.txt"
	path3 := "test_work/work1/merge.txt"
	e := merge(path1, path2, path3)
	// merge.txt之后应该有100行
	fmt.Print(e)
}
