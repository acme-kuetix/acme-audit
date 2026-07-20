package transitions

import (
	"github.com/kuetix/engine/engine/domain/interfaces"
	"github.com/kuetix/engine/engine/workflow"
)

// auditCollection is the persistence/store collection where audit entries live.
const auditCollection = "audit_entries"

type auditTransitions struct {
	workflow.BaseServiceTransition
}

func NewAuditTransitions() interfaces.ServiceTransitions {
	return &auditTransitions{}
}
