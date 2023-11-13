package main

import (
	"cron-parser/cron"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	arguments := flag.Args()
	cronJob, err := cron.NewCronScheduler(arguments)
	if err != nil {
		fmt.Println(err)
		return
	}
	cronJob.Print()
}



```
```