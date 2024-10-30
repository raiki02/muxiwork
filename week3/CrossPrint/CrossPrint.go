/*交叉打印下面两个字符串
"ABCDEFGHIJKLMNOPQRSTUVWXYZ" "0123..."
得到："AB01CD23EF34..."*/

package main

import (
	change "crossprint/change"
	process "crossprint/process"
	"fmt"
)

var (
	s1 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s2 = "01234567891011121314151617"
)

func main() {

	ss1 := make([]string, 0)
	ss2 := make([]string, 0)

	ss1 = change.Change(s1)
	ss2 = change.Change(s2)

	process.Crossprint(ss1, ss2)
	process.Wg.Wait()
	fmt.Println(process.Oc)

}
