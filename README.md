# 小工具：period_time
### 用途：
* 聚合多个有交集的时间段

### 用法：

```go
package main

import (
	"fmt"
	"github.com/tiancheng92/period_time"
	"time"
)

func main() {
	p := period_time.New()
	p.Append(time.Now(), time.Now().Add(2*time.Hour))
	p.Append(time.Now().Add(-1*time.Hour), time.Now().Add(1*time.Hour))
	p.Append(time.Now().Add(10*time.Hour), time.Now().Add(20*time.Hour))
	res := p.Union()
	for _, v := range res {
		fmt.Println(v.StartTime.Format("2006-01-02 15:04:05"), v.EndTime.Format("2006-01-02 15:04:05"))
	}
}

```