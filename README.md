# acme-audit

Append-only audit logging for the kuetix ERP. Captures state transitions
(actor, timestamp, workflow, state, transition, entity, before/after,
success/error) for compliance and observability.

## Quick start

```bash
./runner.sh build
./runner.sh run solutions/audit/record
```

## Transitions

| Method | Description |
|---|---|
| `Record(actor, workflow, state, transition, entityType, entityId, before, after, success, errorMessage)` | Append an entry |
| `List(entityType, entityId, actor, workflow, limit)` | Filter entries (most-recent-first) |
| `Get(entryId)` | Fetch a single entry |

## Cross-package helper

```go
import auditTransitions "github.com/acme-kuetix/acme-audit/modules/audit/transitions"

err := auditTransitions.RecordFromTransition(&auditTransitions.AuditEntry{
    Actor:      "alice",
    Workflow:   "invoice-create",
    Transition: "CreateInvoice",
    EntityType: "invoice",
    EntityID:   "1",
    After:      map[string]interface{}{"number": "INV/0001"},
    Success:    true,
})
```

The composition layer (`erp-app`) uses this helper to capture audit
entries from key lifecycle transitions without a WSL round-trip.

## Architecture

- `SequenceStore`-style interface pattern: `AuditStore` interface with
  in-memory default (swap via `SetStore` for SQL adapter).
- Entries are immutable and append-only — no Update or Delete.
- `List` returns most-recent-first to prioritize the operational
  use case ("what just happened?").

## Dependencies

- [engine](../engine) — workflow runtime
- [std-core](../std-core) — response helpers
