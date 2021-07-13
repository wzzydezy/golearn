package main

import(
	_ "5-init/lib1" //_表示该包可有可无
	mylib2 "5-init/lib2"
	. "5-init/lib2" //尽量不要使用，会造成不同包同名函数的歧义
)

func main(){
	//lib1.Lib1Test()
	//lib2.Lib2Test()
	mylib2.Lib2Test()
	Lib2Test()
}