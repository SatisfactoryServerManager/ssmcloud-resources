package v2

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	TaskStatusPending   = "pending"
	TaskStatusRunning   = "running"
	TaskStatusCompleted = "completed"
	TaskStatusDead      = "dead"
	TaskStatusCancelled = "cancelled"
)

const (
	TaskTriggerUser     = "user"
	TaskTriggerWorkflow = "workflow"
	TaskTriggerSystem   = "system"
)

const TaskDefaultMaxAttempts = 5

// TaskTrigger records who caused a task to exist, so the UI can attribute it.
type TaskTrigger struct {
	Type       string         `json:"type" bson:"type"`
	ExternalID string         `json:"externalId,omitempty" bson:"externalId,omitempty"`
	WorkflowID *bson.ObjectID `json:"workflowId,omitempty" bson:"workflowId,omitempty"`
}

// AgentTaskSchema is one unit of work dispatched to one agent.
//
// Active is present exactly when Status is pending or running. It exists because
// mongo's partialFilterExpression supports $exists but not $in, and the dedupe
// index must cover both non-terminal states.
type AgentTaskSchema struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	AgentID   bson.ObjectID `json:"agentId" bson:"agentId"`
	AccountID bson.ObjectID `json:"accountId" bson:"accountId"`

	Action string `json:"action" bson:"action"`
	Data   string `json:"data" bson:"data"`

	Status string `json:"status" bson:"status"`
	Active *bool  `json:"-" bson:"active,omitempty"`

	Attempts    int `json:"attempts" bson:"attempts"`
	MaxAttempts int `json:"maxAttempts" bson:"maxAttempts"`

	LeaseToken     string    `json:"-" bson:"leaseToken,omitempty"`
	LeaseExpiresAt time.Time `json:"-" bson:"leaseExpiresAt,omitempty"`
	NextAttemptAt  time.Time `json:"nextAttemptAt" bson:"nextAttemptAt"`

	CancelRequested bool   `json:"cancelRequested" bson:"cancelRequested"`
	DedupeKey       string `json:"-" bson:"dedupeKey,omitempty"`

	// DependsOn gates the claim: this task is not claimable until its parent has
	// completed. The parent's terminal transition clears the field (on success) or
	// cancels the child (on death), so the claim itself needs no $lookup.
	DependsOn *bson.ObjectID `json:"dependsOn,omitempty" bson:"dependsOn,omitempty"`

	// RequiresServerStopped gates the claim on the agent reporting the SF server
	// stopped. Without it a deferred syncmods would be claimed and run immediately,
	// because the dispatcher knows nothing about the SF process.
	RequiresServerStopped bool `json:"requiresServerStopped,omitempty" bson:"requiresServerStopped,omitempty"`

	LastError string `json:"lastError,omitempty" bson:"lastError,omitempty"`
	Progress  int32  `json:"progress" bson:"progress"`
	Message   string `json:"message,omitempty" bson:"message,omitempty"`

	TriggeredBy TaskTrigger `json:"triggeredBy" bson:"triggeredBy"`

	CreatedAt  time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt" bson:"updatedAt"`
	StartedAt  *time.Time `json:"startedAt,omitempty" bson:"startedAt,omitempty"`
	FinishedAt *time.Time `json:"finishedAt,omitempty" bson:"finishedAt,omitempty"`
}

// AgentTaskOpts carries the optional gates. Zero value means "claimable as soon
// as the agent is idle", which is what every existing caller wants.
type AgentTaskOpts struct {
	DependsOn             *bson.ObjectID
	RequiresServerStopped bool
}

func NewAgentTaskDoc(agentID, accountID bson.ObjectID, action, data, dedupeKey string, trigger TaskTrigger, opts AgentTaskOpts) AgentTaskSchema {
	now := time.Now()
	active := true

	return AgentTaskSchema{
		ID:                    bson.NewObjectID(),
		AgentID:               agentID,
		AccountID:             accountID,
		Action:                action,
		Data:                  data,
		Status:                TaskStatusPending,
		Active:                &active,
		Attempts:              0,
		MaxAttempts:           TaskDefaultMaxAttempts,
		NextAttemptAt:         now,
		DedupeKey:             dedupeKey,
		DependsOn:             opts.DependsOn,
		RequiresServerStopped: opts.RequiresServerStopped,
		TriggeredBy:           trigger,
		CreatedAt:             now,
		UpdatedAt:             now,
	}
}

// IsTerminal reports whether a task will never run again.
func (t *AgentTaskSchema) IsTerminal() bool {
	return t.Status == TaskStatusCompleted ||
		t.Status == TaskStatusDead ||
		t.Status == TaskStatusCancelled
}
