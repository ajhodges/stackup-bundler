package entrypoint

import (
	"math/big"
	"strings"

	"github.com/ajhodges/stackup-bundler/internal/utils"
	"github.com/ajhodges/stackup-bundler/pkg/tracer"
	"github.com/ethereum/go-ethereum/common"
)

type callEntry struct {
	To     common.Address
	Type   string
	Method string
	Revert any
	Return any
	Value  *big.Int
}

func newCallStack(calls []tracer.CallInfo) []*callEntry {
	out := []*callEntry{}
	stack := utils.NewStack[tracer.CallInfo]()
	for _, call := range calls {
		if call.Type == revertOpCode || call.Type == returnOpCode {
			top, _ := stack.Pop()

			if strings.Contains(top.Type, "CREATE") {
				// TODO: implement...
			} else if call.Type == revertOpCode {
				// TODO: implement...
			} else {
				out = append(out, &callEntry{
					To:     top.To,
					Type:   top.Type,
					Method: top.Method,
					Return: call.Data,
				})
			}
		} else {
			stack.Push(call)
		}
	}

	return out
}
