package v2

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	WorkflowType_CreateAgent = "create-agent"
)

const (
	WorkflowActionType_CreateAgent   = "create-agent"
	WorkflowActionType_WaitForOnline = "wait-for-online"
	WorkflowActionType_AgentTask     = "agent-task"
)

type WorkflowContext struct {
	WorkflowID bson.ObjectID
	ActionIdx  int
}

type IWorkflowAction interface {
	Execute(action *WorkflowAction, workflowData interface{}, account *AccountSchema, wctx WorkflowContext) error
}

type BaseWorkflowData struct {
	AccountId bson.ObjectID `json:"accountid" bson:"accountId"`
}

type CreateAgentWorkflowData struct {
	BaseWorkflowData `bson:",inline"`
	AgentName        string `json:"servername" bson:"serverName"`
	Port             int    `json:"serverPort" bson:"serverPort"`
	Memory           int64  `json:"serverMemory" bson:"serverMemory"`
	AdminPass        string `json:"serverAdminPass" bson:"serverAdminPass"`
	ClientPass       string `json:"serverClientPass" bson:"serverClientPass"`
	APIKey           string `json:"serverAPIKey" bson:"serverApiKey"`
}

type ClaimServer_PostData struct {
	AdminPass  string `json:"adminPass"`
	ClientPass string `json:"clientPass"`
}

type WorkflowSchema struct {
	ID bson.ObjectID `json:"_id" bson:"_id"`
	// AgentId is zero until the workflow's create-agent action has run.
	AgentId bson.ObjectID    `json:"agentId" bson:"agentId"`
	Type    string           `json:"type" bson:"type"`
	Actions []WorkflowAction `json:"actions" bson:"actions"`
	Status  string           `json:"status" bson:"status"`
	Data    interface{}      `json:"data" bson:"data"`
}

type WorkflowAction struct {
	Type         string `json:"type" bson:"type"`
	Status       string `json:"status" bson:"status"`
	ErrorMessage string `json:"error" bson:"error"`
	RetryCount   int    `json:"retryCount" bson:"retryCount"`

	// Set when Type == WorkflowActionType_AgentTask.
	TaskAction string        `json:"taskAction,omitempty" bson:"taskAction,omitempty"`
	TaskData   interface{}   `json:"taskData,omitempty" bson:"taskData,omitempty"`
	TaskID     string        `json:"taskId,omitempty" bson:"taskId,omitempty"`
	Timeout    time.Duration `json:"timeout,omitempty" bson:"timeout,omitempty"`
}
