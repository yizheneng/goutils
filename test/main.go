// main
package main

import "log"
import "time"
import "github.com/yizheneng/goutils"

func main()  {
	elapsedTimer := goutils.NewElapsedTimerWithVal(10000)

	for{
		log.Printf("time:%v istimeOut:%v", goutils.SystemUpTime(), elapsedTimer.IsTimeout())
		time.Sleep(time.Millisecond * 1000)
	}
}