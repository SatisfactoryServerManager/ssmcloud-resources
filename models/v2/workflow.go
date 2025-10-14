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

type IWorkflowAction interface{
    Execute() error
}

type BaseWorkflowData struct {
	AccountId primitive.ObjectID `json:"accountid" bson:"accountId"`
}

type CreateAgentWorkflowData struct {
	BaseWorkflowData
	AgentName  string `json:"servername" bson:"serverName"`
	Port       int    `json:"serverPort" bson:"serverPort"`
	Memory     int64  `json:"serverMemory" bson:"serverMemory"`
	AdminPass  string `json:"serverAdminPass" bson:"serverAdminPass"`
	ClientPass string `json:"serverClientPass" bson:"serverClientPass"`
	APIKey     string `json:"serverAPIKey" bson:"serverApiKey"`
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


// func (obj *WorkflowAction) StartSFServer(workflowData CreateAgentWorkflowData) error {

// 	var theAgent AgentSchema

// 	if err := mongoose.FindOne(bson.M{"apiKey": workflowData.APIKey}, &theAgent); err != nil {
// 		return err
// 	}

// 	newTask := NewAgentTask("startsfserver", nil)

// 	theAgent.Tasks = append(theAgent.Tasks, newTask)

// 	dbUpdate := bson.M{
// 		"tasks":     theAgent.Tasks,
// 		"updatedAt": time.Now(),
// 	}

// 	if err := mongoose.UpdateModelData(&theAgent, dbUpdate); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (obj *WorkflowAction) WaitForRunning(workflowData CreateAgentWorkflowData) (bool, error) {
// 	var theAgent AgentSchema

// 	if err := mongoose.FindOne(bson.M{"apiKey": workflowData.APIKey}, &theAgent); err != nil {
// 		return false, err
// 	}

// 	fmt.Printf("waiting for agent: %s to start sf server \n", theAgent.AgentName)

// 	if !theAgent.Status.Running {
// 		obj.RetryCount += 1
// 		if obj.RetryCount > 120 {
// 			return false, fmt.Errorf("timeout waiting for agent to start sf server")
// 		}

// 		return false, nil
// 	}

// 	return true, nil
// }

// func (obj *WorkflowAction) ClaimServer(workflowData CreateAgentWorkflowData) error {
// 	var theAgent AgentSchema

// 	if err := mongoose.FindOne(bson.M{"apiKey": workflowData.APIKey}, &theAgent); err != nil {
// 		return err
// 	}

// 	data := ClaimServer_PostData{
// 		AdminPass:  workflowData.AdminPass,
// 		ClientPass: workflowData.ClientPass,
// 	}

// 	newTask := NewAgentTask("claimserver", data)

// 	theAgent.Tasks = append(theAgent.Tasks, newTask)

// 	dbUpdate := bson.M{
// 		"tasks":     theAgent.Tasks,
// 		"updatedAt": time.Now(),
// 	}

// 	if err := mongoose.UpdateModelData(&theAgent, dbUpdate); err != nil {
// 		return err
// 	}

// 	return nil
// }
