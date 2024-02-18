package main

import (
	"fmt"
	"unicode/utf8"
)

/*
区别一：
与一般的c++不同，c++存储的是字符，是char类型
我们知道其实char类型也是数值类型，ascii

而golang是直接存储数值类型，但不是ascii,而是范围更大的unicode
字符叫做rune

unicode是全球统一码，任何一个字符都有一个数值来表示

但是一般不会直接用unicode
而会用他的编码版本，最常用的就是utf-8编码
UTF-8 是 Unicode 的一种实现方式
他的设计思想就是：将常用的用短字节表示，不常用的用长的字节来表示
对于常见的字符（如拉丁字母、数字等），使用一个字节表示。
对于较少常见的字符，使用两个或三个字节表示。
对于罕见的字符，可能需要四个字节表示。
他包含了所有的unicode，只不过他的字符串每个字符，所占的字节个数不定

现在使用utf-8作为字符串是主流，而不是ascii,因为这样很好的解决了中文乱码问题

区别二:
golang的字符串是不可变的，这和python,java是一样的，c++的字符串完全是用的数组来实现的
所以c++的字符串是可变的

c++使用了\0来标识字符串的结尾，这是他使用区别数组所采取的措施
但是golang没有
*/

func getlen(s string) {
	// len操作求的是字符串s的字节个数
	fmt.Println(s, "所占的字节个数  = ", len(s))

	// 求字符的个数
	fmt.Println(s, "字符的个数", utf8.RuneCountInString(s))
}

func main() {
	getlen("hello")
	getlen("hello你好") // 可见你，好各占了3个字节
}
