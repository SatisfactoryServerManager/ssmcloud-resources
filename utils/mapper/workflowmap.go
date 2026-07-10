package mapper

import (
	v2 "github.com/SatisfactoryServerManager/ssmcloud-resources/models/v2"
	pbModels "github.com/SatisfactoryServerManager/ssmcloud-resources/proto/generated/models"
)

func MapWorkflowToProto(workflow *v2.WorkflowSchema) *pbModels.Workflow {

	pbWorkflowActions := make([]*pbModels.WorkflowAction, 0, len(workflow.Actions))

	for i := range workflow.Actions {
		pbWorkflowActions = append(pbWorkflowActions, MapWorkflowActionToProto(&workflow.Actions[i]))
	}

	agentId := ""
	if !workflow.AgentId.IsZero() {
		agentId = workflow.AgentId.Hex()
	}

	return &pbModels.Workflow{
		Id:      workflow.ID.Hex(),
		AgentId: agentId,
		Type:    workflow.Type,
		Status:  workflow.Status,
		Actions: pbWorkflowActions,
	}
}

func MapWorkflowActionToProto(workflowAction *v2.WorkflowAction) *pbModels.WorkflowAction {

	return &pbModels.WorkflowAction{
		Type:         workflowAction.Type,
		Status:       workflowAction.Status,
		ErrorMessage: workflowAction.ErrorMessage,
		RetryCount:   int32(workflowAction.RetryCount),
		TaskAction:   workflowAction.TaskAction,
	}
}
