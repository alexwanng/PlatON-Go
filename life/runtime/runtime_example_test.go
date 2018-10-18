package runtime_test

import (
	"fmt"

	"Platon-go/life/runtime"
)

func ExampleExecute() {
	// code, abi, input, cfg
	abi := []byte{}
	code := []byte{}
	ret, _, err := runtime.Execute(code, abi, nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
	// Output:
	// [96 96 96 64 82 96 8 86 91 0]
}
