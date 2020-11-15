package basic

import (
	"fmt"
)

func test_scan() {
	// 测试键盘输入
	var (
        name    string
        age     int
        married bool
    )
    fmt.Scan(&name, &age, &married)
    fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

// func test_fscan() {
    //测试Fscan
// }

// func test_sscan() {
   //测试Sscan
// }