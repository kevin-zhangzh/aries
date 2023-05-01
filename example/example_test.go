package example

import (
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/didexchange"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	fram, err := aries.New()
	ctx, err := fram.Context()
	cli, err := didexchange.New(ctx)
	assert.NoError(t, err)
	cli.Name()
}
