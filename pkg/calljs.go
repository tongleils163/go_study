package pkg

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

func CallJs(word string) string {
	filePath := "./baidu_word.js"
	// 先读入文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	vm := otto.New()

	_, err = vm.Run(string(bytes))
	if err != nil {
		panic(err)
	}

	data := word
	// encodeInp是JS函数的函数名
	value, err := vm.Call("tk", nil, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(value.String())
	return value.String()
}
