package bug

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/daedaleanai/git-ticket/identity"
	"github.com/stretchr/testify/assert"
)

func TestSetStatusSerialize(t *testing.T) {
	var rene = identity.NewBare("René Descartes", "rene@descartes.fr")
	unix := time.Now().Unix()
	before := NewSetStatusOp(rene, unix, MergedStatus)

	data, err := json.Marshal(before)
	assert.NoError(t, err)

	var after SetStatusOperation
	err = json.Unmarshal(data, &after)
	assert.NoError(t, err)

	// enforce creating the IDs
	before.Id()
	rene.Id()

	assert.Equal(t, before, &after)
}
