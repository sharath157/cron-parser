package cron

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestParserExpression(t *testing.T) {

	testData := []struct {
		TestName string
		Exp      string
		MinRange int
		MaxRange int
		Expected []int
		IsError  bool
	}{
		{
			TestName: "should accept all values ",
			Exp:      "*",
			MinRange: 0,
			MaxRange: 5,
			Expected: []int{0, 1, 2, 3, 4, 5},
			IsError:  false,
		},
		{
			TestName: "should accept Range",
			Exp:      "1-2",
			MinRange: 1,
			MaxRange: 2,
			Expected: []int{1, 2},
			IsError:  false,
		},
		{
			TestName: "Invalid Range , should fail",
			Exp:      "1-2",
			MinRange: 1,
			MaxRange: 1,
			Expected: []int{},
			IsError:  true,
		},
		{
			TestName: "Range by step - should fail - invalid range",
			Exp:      "1-2/1",
			MinRange: 1,
			MaxRange: 1,
			Expected: []int{1},
			IsError:  true,
		},
		{
			TestName: "* range with step - should success-corner case",
			Exp:      "*/1",
			MinRange: 1,
			MaxRange: 1,
			Expected: []int{1},
			IsError:  false,
		},
		{
			TestName: "* range with step - should success",
			Exp:      "*/3",
			MinRange: 1,
			MaxRange: 10,
			Expected: []int{1, 4, 7, 10},
			IsError:  false,
		},
		{
			TestName: "finite value - should success",
			Exp:      "1",
			MinRange: 1,
			MaxRange: 10,
			Expected: []int{1},
			IsError:  false,
		},
		{
			TestName: "finite values - should success",
			Exp:      "1,2",
			MinRange: 1,
			MaxRange: 10,
			Expected: []int{1, 2},
			IsError:  false,
		},
		{
			TestName: "finite values - should fail - out of range",
			Exp:      "1,12",
			MinRange: 1,
			MaxRange: 10,
			Expected: []int{1},
			IsError:  true,
		},
	}

	parser := Parser{}
	for _, tt := range testData {
		t.Run(tt.TestName, func(t *testing.T) {
			res, err := parser.Parse(tt.Exp, tt.MinRange, tt.MaxRange)
			if tt.IsError {
				assert.NotNil(t, err)
			} else {
				assert.True(t, reflect.DeepEqual(tt.Expected, res))
			}

		})
	}

}

func TestNewParser(t *testing.T) {

	parse, _ := NewParser([]string{"*/15", "0-23", "1,15", "*", "1-2"})
	expectedParser := Parser{
		Minutes:     []int{0, 15, 30, 45},
		Hours:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
		DaysOfMonth: []int{1, 15},
		Months:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		weekNumber:  []int{1, 2},
	}
	assert.True(t, reflect.DeepEqual(expectedParser, parse))
}

func TestNewErrParser(t *testing.T) {
	_, err := NewParser([]string{"*/15", "0-28", "1,15", "*", "1-2"})
	assert.NotNil(t, err)

	_, err = NewParser([]string{"*/15", "0-a", "1,15", "*", "1-2"})
	assert.NotNil(t, err)
}
