package main

import (
	"fmt"
	s "strings"
)

/*
 */
var p = fmt.Println

func main() {
	p("Add:       ", "abc"+"edf")                     // 字符串加法
	p("Contains:  ", s.Contains("test", "es"))        // 是否包含
	p("Count:     ", s.Count("test", "t"))            //统计有几个子串
	p("HasPrefix: ", s.HasPrefix("test", "te"))       // 是否有前缀
	p("HasSuffix: ", s.HasSuffix("test", "st"))       // 是否有后缀
	p("Index:     ", s.Index("test", "e"))            // 字符串查找
	p("Join:      ", s.Join([]string{"a", "b"}, "-")) //字符串拼接
	p("Repeat:    ", s.Repeat("a", 5))                // 字符串重复再拼接
	p("Replace:   ", s.Replace("foo", "o", "0", -1))  //字符串替换，第4个参数是替换个数
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-")) //切分字符串
	p("ToLower:   ", s.ToLower("TEST"))         // 转换成小写
	p("ToUpper:   ", s.ToUpper("test"))         // 转换成大写
	p()

	p("Len: ", len("hello")) // 求字符串字节数量
	p("Char:", "hello"[1])   // 根据index取元素
}
