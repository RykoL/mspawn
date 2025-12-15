package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitCmd(t *testing.T) {

	t.Run("splits command and arguments", func(t *testing.T) {
		cmd := "python -m http.server"

		program, argument := splitCmd(cmd)
		
		assert.Equal(t, program, "python")
		assert.Equal(t, argument, []string{"-m", "http.server"})
	})

	t.Run("returns program only", func(t *testing.T) {
		cmd := "httpd"

		program, argument := splitCmd(cmd)
		
		assert.Equal(t, program, "httpd")
		assert.Empty(t, argument)
	})
	
}
