package workflows

import (
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/workflow"
)

type syncWorkflowInterceptor struct{}

func (i *syncWorkflowInterceptor) PreWorkflowExecutionHook(ctx workflow.Context, args ...any) (workflow.Context, error) {
	return ctx, nil
}

func (i *syncWorkflowInterceptor) PostWorkflowExecutionHook(ctx workflow.Context, res any) error {
	return nil
}

func (i *syncWorkflowInterceptor) WorflowExecutionErrorHook(ctx workflow.Context, err error, args ...any) error {
	return nil
}

type syncWorkflowInboundInterceptor struct {
	interceptor.WorkflowInboundInterceptorBase
	root syncWorkflowInterceptor
}

func (i *syncWorkflowInboundInterceptor) ExecuteWorkflow(ctx workflow.Context, input *interceptor.ExecuteWorkflowInput) (any, error) {
	ctx, err := i.root.PreWorkflowExecutionHook(ctx, input.Args...)
	if err != nil {
		return nil, err
	}
	res, err := i.Next.ExecuteWorkflow(ctx, input)
	if err != nil {
		return res, i.root.WorflowExecutionErrorHook(ctx, err, input.Args...)
	}
	return res, i.root.PostWorkflowExecutionHook(ctx, res)
}
