package flowcontext

import (
	"github.com/pugdag/pugdagd/domain"
)

// Domain returns the Domain object associated to the flow context.
func (f *FlowContext) Domain() domain.Domain {
	return f.domain
}
