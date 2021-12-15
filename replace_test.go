package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceAll(t *testing.T) {
	regex := Constant(`(\[?([0-9A-Fa-f]{1,4}:){7}[0-9A-Fa-f]{1,4}|(\d{1,3}\.){3}\d{1,3})(\]?:[0-9]+)?`)
	badRegex := Constant(`(()))`)
	nonConstantRegex := Param("regex")

	assert.Equal(t, "Dialed <addr>", msgpacked(t, ReplaceAll(Constant("Dialed 127.0.0.1"), regex, Constant("<addr>"))).Eval(nil))
	assert.Equal(t, "Dialed <addr>", msgpacked(t, ReplaceAll(Constant("Dialed 127.0.0.1:8789"), regex, Constant("<addr>"))).Eval(nil))
	assert.Equal(t, "Dialed <addr>", msgpacked(t, ReplaceAll(Constant("Dialed 2605:a601:41f3:9900:4c57:94a4:ac48:db3a"), regex, Constant("<addr>"))).Eval(nil))
	assert.Equal(t, "Dialed <addr>", msgpacked(t, ReplaceAll(Constant("Dialed [2605:a601:41f3:9900:4c57:94a4:ac48:db3a]:8789"), regex, Constant("<addr>"))).Eval(nil))
	assert.Equal(t, "Input", msgpacked(t, ReplaceAll(Constant("Input"), badRegex, Constant("replacement"))).Eval(nil))
	assert.Equal(t, "Input", msgpacked(t, ReplaceAll(Constant("Input"), nonConstantRegex, Constant("replacement"))).Eval(nil))
	assert.Equal(t, nil, msgpacked(t, ReplaceAll(Constant(nil), regex, Constant("<addr>"))).Eval(nil))

}
