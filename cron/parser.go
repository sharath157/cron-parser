package cron

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	minMinutes = 0
	maxMinutes = 59
	minHours   = 0
	maxHours   = 23
	minMonths  = 1
	maxMonths  = 12
	minDays    = 1
	maxDays    = 31
	minWeekDay = 1
	maxWeekDay = 7
)

type Parser struct {
	Minutes     []int
	Hours       []int
	DaysOfMonth []int
	Months      []int
	weekNumber  []int
}

func NewParser(exp []string) (Parser, error) {
	var p Parser
	if len(exp) != 5 {
		return p, fmt.Errorf("invalid cron expression,shoud be in format: [minutes] [hours] [Days of the months] [Months]  [day of week] ")
	}

	if data, err := p.Parse(exp[0], minMinutes, maxMinutes); err != nil {
		return p, err
	} else {
		p.Minutes = data
	}

	if data, err := p.Parse(exp[1], minHours, maxHours); err != nil {
		return p, err
	} else {
		p.Hours = data
	}
	if data, err := p.Parse(exp[2], minDays, maxDays); err != nil {
		return p, err
	} else {
		p.DaysOfMonth = data
	}
	if data, err := p.Parse(exp[3], minMonths, maxMonths); err != nil {
		return p, err
	} else {
		p.Months = data
	}
	if data, err := p.Parse(exp[4], minWeekDay, maxWeekDay); err != nil {
		return p, err
	} else {
		p.weekNumber = data
	}

	return p, nil
}

func (p *Parser) Parse(expr string, minRange, maxRange int) ([]int, error) {

	mins := make([]int, 0)
	if expr == "*" {
		for i := minRange; i <= maxRange; i++ {
			mins = append(mins, i)
		}
		return mins, nil
	} else if strings.Contains(expr, "/") {
		data := strings.SplitN(expr, "/", 2)
		if len(data) != 2 {
			return nil, fmt.Errorf("invalid cron parser at : %v", expr)
		}
		step, err := strconv.Atoi(data[1])
		if err != nil {
			return nil, fmt.Errorf("invalid  parser at %v ,%w", expr, err)
		}
		startRange, endRange := minRange, maxRange
		if strings.Contains(data[0], "-") {
			startRange, endRange, err = ParseRange(data[0], minRange, maxRange)
			if err != nil {
				return nil, err
			}
		}
		for i := startRange; i <= endRange; i = i + step {
			mins = append(mins, i)
		}
		return mins, nil
	} else if strings.Contains(expr, "-") {
		startRange, endRange, err := ParseRange(expr, minRange, maxRange)
		if err != nil {
			return nil, err
		}
		for i := startRange; i <= endRange; i++ {
			mins = append(mins, i)
		}
		return mins, nil
	} else if strings.Contains(expr, ",") {
		list := strings.Split(expr, ",")
		for _, data := range list {
			if value, err := strconv.Atoi(data); err != nil {
				return nil, fmt.Errorf("invalid data in parser %v %w", data, err)
			} else {
				if value < minRange || value > maxRange {
					return nil, fmt.Errorf("invalid data in parser %v minRange:%v maxRange:%v", data, minRange, maxRange)
				}
				mins = append(mins, value)
			}
		}
		return mins, nil
	}
	if data, err := strconv.Atoi(expr); err != nil {
		return nil, fmt.Errorf("invalid data in Parser:%w", err)
	} else {
		if data < minRange || data > maxRange {
			return nil, fmt.Errorf("invalid data in parser %v minRange:%v maxRange:%v", data, minRange, maxRange)
		}
		return []int{data}, nil
	}

	return nil, fmt.Errorf("unrecognized format in parser:%v", expr)
}

func ParseRange(expr string, minRange, maxRange int) (int, int, error) {
	data := strings.SplitN(expr, "-", 2)
	startRange, err := strconv.Atoi(data[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid startRange in Parser:%v %w", expr, err)
	}
	endRange, err := strconv.Atoi(data[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid endRange in Parser:%v %w", expr, err)
	}
	if startRange > endRange || startRange < minRange || endRange > maxRange {
		return 0, 0, fmt.Errorf("invalid startRange-endRange in Parser %v %v", startRange, endRange)
	}
	return startRange, endRange, nil
}
