package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSpawnProcess(t *testing.T) {
	err := SpawnProcess(t.Context(), "go", []string{"version"})

	assert.NoError(t, err)

	t.Run("return err if process could not be started", func(t *testing.T) {
		err := SpawnProcess(t.Context(), "some binary that doesn't exists", []string{"version"})

		assert.Error(t, err)
	})
}
