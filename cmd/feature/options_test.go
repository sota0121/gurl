package feature

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test Suite for the package feature.OptValue
func TestOptValue(t *testing.T) {
	testMap := map[string]func(t *testing.T){
		"New OptValue Instance": testNewOptValue,
	}

	for scenario, fn := range testMap {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testNewOptValue(t *testing.T) {
	// string value pattern
	v := NewOptValue(String)
	v.strValue = "test"
	require.Equal(t, String, v.GetType())
	s := v.String()
	require.Equal(t, "test", s)

	// bool value pattern
	v = NewOptValue(Bool)
	v.boolValue = true
	require.Equal(t, Bool, v.GetType())
	b := v.Bool()
	require.Equal(t, true, b)

	// int value pattern
	v = NewOptValue(Int)
	v.intValue = 123
	require.Equal(t, Int, v.GetType())
	i := v.Int()
	require.Equal(t, 123, i)
}

// Test Suite for the package feature.Opt
func TestOpt(t *testing.T) {
	testMap := map[string]func(t *testing.T){
		"New Opt Instance":        testNewOpt,
		"Constant Options Access": testConstantOptionsAccess,
	}

	for scenario, fn := range testMap {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testNewOpt(t *testing.T) {
	// string value pattern
	v := Opt{
		Name:      data.Name,
		Help:      "help",
		Usage:     "usage",
		Required:  false,
		Supported: true,
		Value:     *NewOptValue(String),
	}
	require.Equal(t, data.Name, v.Name)
}

func testConstantOptionsAccess(t *testing.T) {
	optSize := GetOptSize()
	require.Equal(t, len(options), optSize)

	for i := 0; i < optSize; i++ {
		opt, err := GetOptWithIndex(i)
		require.NoError(t, err)
		require.Equal(t, options[i], opt)
	}
	outOfIndex := optSize + 1
	_, err := GetOptWithIndex(outOfIndex)
	require.Error(t, err)

	for _, opt := range options {
		newOpt, err := GetOptWithName(opt.Name)
		require.NoError(t, err)
		require.Equal(t, opt, newOpt)
	}
	_, err = GetOptWithName("not-exist")
	require.Error(t, err)

	supportedOptSize := GetSupportedOptSize()
	supportedOptions := []Opt{}
	for _, opt := range options {
		if opt.Supported {
			supportedOptions = append(supportedOptions, opt)
		}
	}
	require.Equal(t, len(supportedOptions), supportedOptSize)
}
