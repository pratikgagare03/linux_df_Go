package sorting

import (
	"testing"
	"project/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestGetTop2Avail(t *testing.T) {
	t.Run("TestEmptyInput", func(t *testing.T) {
		var system []parser.Fs

		_, err := GetTop2Avail(system)

		require.NotNil(t, err, "Expected error for empty input")
		
	})

	t.Run("TestValidInput", func(t *testing.T) {
		input := []parser.Fs{
			{
				Name:      "def",
				Size:      "100G",
				Used:      "50G",
				Avail:     "50G",
				Use:       "50%",
				MountedOn: "/",
			},
			{
				Name:      "abc",
				Size:      "250G",
				Used:      "50G",
				Avail:     "200G",
				Use:       "20%",
				MountedOn: "/dev/shm",
			},
		}
		expected := []parser.Fs{
			{
				Name:      "abc",
				Size:      "250G",
				Used:      "50G",
				Avail:     "200G",
				Use:       "20%",
				MountedOn: "/dev/shm",
			},
			{
				Name:      "def",
				Size:      "100G",
				Used:      "50G",
				Avail:     "50G",
				Use:       "50%",
				MountedOn: "/",
			},
		}

		output, err := GetTop2Avail(input)
		require.Nil(t, err)
		assert.Equal(t, expected, output)
	})
}


func TestGetTop2Size(t *testing.T) {
	t.Run("TestEmptyInput", func(t *testing.T) {
		var system []parser.Fs

		_, err := GetTop2Size(system)

		require.NotNil(t, err, "Expected error for empty input")

	})

	t.Run("TestValidInput", func(t *testing.T) {
		input := []parser.Fs{
			{
				Name:      "abc",
				Size:      "100G",
				Used:      "50G",
				Avail:     "50G",
				Use:       "50%",
				MountedOn: "/",
			},
			{
				Name:      "def",
				Size:      "250G",
				Used:      "50G",
				Avail:     "200G",
				Use:       "20%",
				MountedOn: "/dev/shm",
			},
		}
		expected := []parser.Fs{
			{
				Name:      "def",
				Size:      "250G",
				Used:      "50G",
				Avail:     "200G",
				Use:       "20%",
				MountedOn: "/dev/shm",
			},
			{
				Name:      "abc",
				Size:      "100G",
				Used:      "50G",
				Avail:     "50G",
				Use:       "50%",
				MountedOn: "/",
			},
		}

		output, err := GetTop2Size(input)

		require.Nil(t, err, "Unexpected error for valid input")
		assert.Equal(t, expected, output, "Output should match expected")
	})
}

func TestGetTop2Use(t *testing.T) {
	t.Run("TestEmptyInput", func(t *testing.T) {
		var system []parser.Fs

		_, err := GetTop2Use(system)

		require.NotNil(t, err, "Expected error for empty input")

	})

	t.Run("TestValidInput", func(t *testing.T) {
		input := []parser.Fs{
			{
				Name:      "abc",
				Size:      "100G",
				Used:      "50G",
				Avail:     "50G",
				Use:       "50%",
				MountedOn: "/",
			},
			{
				Name:      "def",
				Size:      "250G",
				Used:      "50G",
				Avail:     "200G",
				Use:       "20%",
				MountedOn: "/dev/shm",
			},
		}
		expected := []parser.Fs{
			{
				Name:      "abc",
				Size:      "100G",
				Used:      "50G",
				Avail:     "50G",
				Use:       "50%",
				MountedOn: "/",
			},
			{
				Name:      "def",
				Size:      "250G",
				Used:      "50G",
				Avail:     "200G",
				Use:       "20%",
				MountedOn: "/dev/shm",
			},
			
		}

		output, err := GetTop2Use(input)
		require.Nil(t, err, "Unexpected error for valid input")
		assert.Equal(t, expected, output, "Output should match expected")
	})
}