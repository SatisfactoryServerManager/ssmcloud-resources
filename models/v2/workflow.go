package v2

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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




// func (obj *WorkflowAction) CreateAgent(workflowData CreateAgentWorkflowData, theAccount *AccountSchema) error {
// 	newAgent := NewAgent(workflowData.AgentName, workflowData.Port, workflowData.Memory, workflowData.APIKey)

// 	if _, err := mongoose.InsertOne(&newAgent); err != nil {
// 		return fmt.Errorf("error inserting new agent with error: %s", err.Error())
// 	}

// 	theAccount.AgentIds = append(theAccount.AgentIds, newAgent.ID)

// 	dbUpdate := bson.M{
// 		"agents":    theAccount.AgentIds,
// 		"updatedAt": time.Now(),
// 	}

// 	if err := mongoose.UpdateModelData(*theAccount, dbUpdate); err != nil {
// 		return fmt.Errorf("error updating account AgentSchema with error: %s", err.Error())
// 	}

// 	theAccount.AddAudit("CREATE_AGENT", fmt.Sprintf("New agent created (%s)", workflowData.AgentName))
// 	return nil
// }

// func (obj *WorkflowAction) WaitForOnline(workflowData CreateAgentWorkflowData) (bool, error) {
// 	var theAgent AgentSchema

// 	if err := mongoose.FindOne(bson.M{"apiKey": workflowData.APIKey}, &theAgent); err != nil {
// 		return false, err
// 	}

// 	fmt.Printf("waiting for agent: %s to be online \n", theAgent.AgentName)

// 	if !theAgent.Status.Online {
// 		obj.RetryCount += 1
// 		if obj.RetryCount > 120 {
// 			return false, fmt.Errorf("timeout waiting for agent to start")
// 		}

// 		return false, nil
// 	}

// 	return true, nil
// }

// func (obj *WorkflowAction) InstallSFServer(workflowData CreateAgentWorkflowData) error {

// 	var theAgent AgentSchema

// 	if err := mongoose.FindOne(bson.M{"apiKey": workflowData.APIKey}, &theAgent); err != nil {
// 		return err
// 	}

// 	newTask := NewAgentTask("installsfserver", nil)

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

// func (obj *WorkflowAction) WaitForInstalled(workflowData CreateAgentWorkflowData) (bool, error) {
// 	var theAgent AgentSchema

// 	if err := mongoose.FindOne(bson.M{"apiKey": workflowData.APIKey}, &theAgent); err != nil {
// 		return false, err
// 	}

// 	fmt.Printf("waiting for agent: %s to install sf server \n", theAgent.AgentName)

// 	if !theAgent.Status.Installed {
// 		obj.RetryCount += 1
// 		if obj.RetryCount > 120 {
// 			return false, fmt.Errorf("timeout waiting for agent to install sf server")
// 		}

// 		return false, nil
// 	}

// 	return true, nil
// }

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
