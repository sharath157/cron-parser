package cron

import (
	"fmt"
	"strings"
)

type CronScheduler struct {
	Scheduler Parser
	Cmd       string
}

func NewCronScheduler(arguments []string) (CronScheduler, error) {
	var scheduler CronScheduler
	if len(arguments) != 6 {
		return scheduler, fmt.Errorf("invalid expression to cron scheduler")
	}
	parser, err := NewParser(arguments[0:5])
	if err != nil {
		return CronScheduler{}, err
	}
	scheduler.Scheduler = parser
	scheduler.Cmd = arguments[5]
	return scheduler, nil
}

func (c *CronScheduler) Print() {
	fmt.Printf("%-14s %v \n", "minute", JoinArray(c.Scheduler.Minutes, " "))
	fmt.Printf("%-14s %v \n", "hour", JoinArray(c.Scheduler.Hours, " "))
	fmt.Printf("%-14s %v \n", "day of month", JoinArray(c.Scheduler.DaysOfMonth, " "))
	fmt.Printf("%-14s %v \n", "month", JoinArray(c.Scheduler.Months, " "))
	fmt.Printf("%-14s %v \n", "day of week", JoinArray(c.Scheduler.weekNumber, " "))
	fmt.Printf("%-14s %v \n", "command", c.Cmd)
}

func JoinArray[T any](data []T, sep string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data)), sep), "[]")
}
