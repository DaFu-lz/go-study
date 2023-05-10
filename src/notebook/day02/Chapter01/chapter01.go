package main

//场合一,使用net/http/pprof的初始化函数init()
import (
	"fmt"
	_ "net/http/pprof"
)

//自定义函数，设置两个返回值
func myfunc() (int, string) {
	a := 10
	b := "golang"
	return a, b
}

//定义接口
type Foo interface {
	Say()
}

//定义结构体
type Dog struct {
	name string
}

//结构体实现接口Foo的使用
func (d Dog) Say() {
	fmt.Println(d.name + "say hello!")
}

func main() {
	//场合二：调用函数myfunc()并只返回第一个值
	a, _ := myfunc()
	fmt.Printf("返回值：%d\n", a)

	//场合三：判断结构体Dog是否实现接口Foo的使用
	//等同于判定有没有定义func (d Dog) Say(){}，如果Dog没有实现Foo，则会报编译错误
	var _ Foo = Dog{"black dog"}
}
