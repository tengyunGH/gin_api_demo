// gin_api_demo   tengyun；2021/8/13 4:53 下午
package main

import "fmt"

type another struct {
	Hu int64
}

type test struct {
	AInt    int64
	Another another
}

// map中的指针，没有初始化空间或者没有相应的key的时候，都返回了nil和false
func main() {
	var stringMap map[string]*string

	fmt.Println(stringMap)

	bString, ok := stringMap["te"]
	fmt.Println(bString)
	fmt.Println(ok)

	fmt.Println(bString) // 增加修改main方法

	stringSlice := make([]string, 2)
	aString := stringSlice[0]
	fmt.Println(aString)

	var a []test
	fmt.Println(a)
	b := make([]test, 2)

	fmt.Println(fmt.Sprintf("aaa %v, bbb %v", a, b))

	var aMap map[string]*test

	aValue, ok := aMap["ty"]
	fmt.Println(fmt.Sprintf("1111 %v, %v", aValue, ok))

	aMap = make(map[string]*test)
	aMap["tengyun"] = &test{AInt: 89}

	bValue, ok := aMap["ty"]
	fmt.Println(fmt.Sprintf("2222 %v, %v", bValue, ok))

	cValue, ok := aMap["tengyun"]
	fmt.Println(fmt.Sprintf("3333 %v, %v", cValue, ok))
}
