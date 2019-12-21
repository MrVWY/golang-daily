package main
//regexp包
import (
	"fmt"
	"regexp"
)

func main() {
	_, _ = regexp.MatchString("H(.*)d!","Hello world!")
	_, _ = regexp.Match("H(.*)d!",[]byte("Hello world!"))

	r, _ := regexp.Compile("H(.*)d!")
	fmt.Println(r.MatchString("Hello world!"))//true
	fmt.Println(r.FindString(" Hello world! Hello"))//Hello world!
	fmt.Println(r.FindStringIndex("Hello World! worl"))//[0 12]
	fmt.Println(r.FindStringSubmatch("Hello World! world")) //[`H(.*)d!` : Hello World!    `(.*)`:  ello Worl]
	fmt.Println(r.FindStringSubmatchIndex("Hello World! world")) //[起始索引: 0 12  结束索引: 1 10]
	fmt.Println(r.FindAllString("Hello World! Held! world", -1)) //[Hello World! Held!]  返回所有正则匹配的字符，不仅仅是第一个

	res, _ := regexp.Compile("H([a-z]+)d!")
	fmt.Println(res.FindAllString("Hello World! Held! Hellowrld! world", 2)) //[Held! Hellowrld!]

	fmt.Println(r.ReplaceAllString("Hello World! Held! world", "html")) //html world   代替

	reg := regexp.MustCompile(`(\w+),(\w+)`)
	src := []byte("Golang,World!")           // 源文本
	dst := []byte("Say: ")                   // 目标文本
	template := []byte("Hello $1, Hello $2") // 模板
	m := reg.FindSubmatchIndex(src)          // 解析源文本
	fmt.Printf("%q", reg.Expand(dst, template, src, m))// "Say: Hello Golang, Hello World"
}