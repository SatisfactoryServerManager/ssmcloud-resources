package v2

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	WorkflowType_CreateAgent = "create-agent"
)

const (
	WorkflowActionType_CreateAgent      = "create-agent"
	WorkflowActionType_WaitForOnline    = "wait-for-online"
	WorkflowActionType_InstallServer    = "install-server"
	WorkflowActionType_WaitForInstalled = "wait-for-installed"
	WorkflowActionType_StartServer      = "start-server"
	WorkflowActionType_WaitForRunning   = "wait-for-running"
	WorkflowActionType_ClaimServer      = "claim-server"
)

type IWorkflowAction interface {
	Execute(action *WorkflowAction, workflowData interface{}, account *AccountSchema) error
}

type BaseWorkflowData struct {
	AccountId primitive.ObjectID `json:"accountid" bson:"accountId"`
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
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Type    string             `json:"type" bson:"type"`
	Actions []WorkflowAction   `json:"actions" bson:"actions"`
	Status  string             `json:"status" bson:"status"`
	Data    interface{}        `json:"data" bson:"data"`
}

type WorkflowAction struct {
	Type         string `json:"type" bson:"type"`
	Status       string `json:"status" bson:"status"`
	ErrorMessage string `json:"error" bson:"error"`
	RetryCount   int    `json:"retryCount" bson:"retryCount"`
}
