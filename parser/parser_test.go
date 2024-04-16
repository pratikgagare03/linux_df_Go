package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetResults(t *testing.T) {
	t.Run("TestEmptyInput", func(t *testing.T) {
		var input string

		_, err := GetResults(input)

		require.NotNil(t, err, "Expected error for empty input")

	})

	t.Run("TestValidInput", func(t *testing.T) {
		input := "Filesystem      Size  Used Avail Use% Mounted on\ntmpfs           1.6G  2.1M  1.6G   1% /run\n/dev/sda3       457G   19G  416G   5% /\ntmpfs           7.8G   70M  7.7G   1% /dev/shm\ntmpfs           5.0M  4.0K  5.0M   1% /run/lock\nefivarfs        128K   91K   33K  74% /sys/firmware/efi/efivars\ntmpfs           7.8G     0  7.8G   0% /run/qemu\n/dev/sda2       512M  6.1M  506M   2% /boot/efi\ntmpfs           1.6G  6.3M  1.6G   1% /run/user/1000"
		expected := []Fs{
			{Name:"tmpfs", Size:"1.6G", Used:"2.1M", Avail:"1.6G", Use:"1%", MountedOn:"/run"},
			{Name:"/dev/sda3", Size:"457G", Used:"19G", Avail:"416G", Use:"5%", MountedOn:"/"}, {Name:"tmpfs", Size:"7.8G", Used:"70M" ,Avail:"7.7G" ,Use:"1%" ,MountedOn:"/dev/shm"} ,{Name:"tmpfs", Size:"5.0M" ,Used:"4.0K" ,Avail:"5.0M" ,Use:"1%", MountedOn:"/run/lock"} ,{Name:"efivarfs", Size:"128K", Used:"91K", Avail:"33K", Use:"74%", 
			       MountedOn:"/sys/firmware/efi/efivars"} ,
			{Name:"tmpfs" ,Size:"7.8G", Used:"0", Avail:"7.8G", Use:"0%" ,MountedOn:"/run/qemu"} ,{Name:"/dev/sda2", Size:"512M", Used:"6.1M", Avail:"506M", Use:"2%", MountedOn:"/boot/efi"}, 
			{Name:"tmpfs", Size:"1.6G", Used:"6.3M", Avail:"1.6G", Use:"1%", MountedOn:"/run/user/1000"},
		}
			

		output, err := GetResults(input)
		require.Nil(t, err, "Unexpected error for valid input")
		assert.Equal(t, expected, output, "Output should match expected")
	})
}