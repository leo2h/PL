// 每个 Go 程序都是由包构成的
// 程序从 main 包开始运行
package main

import "fmt" // 导入 fmt 包
import ( // 分组导入
  "math"
  "time"
)

func Swap(a, b int) (int, int) {
  fmt.Println(a, b, time.Now())
  return b, a
}

func add(x, y float64) (r float64) {
  r = math.Abs(x + y)
  return
}
