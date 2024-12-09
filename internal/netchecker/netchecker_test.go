package netchecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	exList := []struct {
		name    string
		subnet  string
		inpIP   string
		result  bool
		withErr bool
	}{
		{
			name:    "happy_path_contains",
			subnet:  "192.168.1.1/16",
			inpIP:   "192.168.1.2",
			result:  true,
			withErr: false,
		},
		{
			name:    "when_not_contains",
			subnet:  "192.168.1.1/16",
			inpIP:   "5.5.5.5",
			result:  false,
			withErr: false,
		},
		{
			name:    "when_empty_subnet",
			subnet:  "",
			inpIP:   "5.5.5.5",
			result:  false,
			withErr: true,
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			h, err := Make(ex.subnet)
			if ex.withErr {
				assert.NotEqual(t, nil, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, ex.result, h.Contains(ex.inpIP))
			}

		})
	}
}
