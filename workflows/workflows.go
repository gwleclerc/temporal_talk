package workflows

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.temporal.io/sdk/workflow"
)

type RenewInput struct {
	Domain                string    `json:"domain" validate:"required"`
	Period                int64     `json:"period"  validate:"required"`
	CurrentExpirationDate time.Time `json:"current_expiration_date" validate:"required"`
}

func WorkflowDefinition(ctx workflow.Context, workflowInput RenewInput) error {
	err := validator.New().Struct(&workflowInput)
	if err != nil {
		return err
	}

	var domainStatus DomainStatusOutput
	err = workflow.ExecuteActivity(ctx, GetDomainStatus, DomainStatusInput{Domain: workflowInput.Domain}).Get(ctx, &domainStatus)
	if err != nil {
		return err
	}

	var premiumStatus CheckPremiumStatusOutput
	err = workflow.ExecuteActivity(ctx, CheckPremiumStatus, CheckPremiumStatusInput{Domain: workflowInput.Domain}).Get(ctx, &premiumStatus)
	if err != nil {
		return err
	}

	...

	return nil
}
