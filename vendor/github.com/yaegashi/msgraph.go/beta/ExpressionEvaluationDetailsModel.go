// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ExpressionEvaluationDetails undocumented
type ExpressionEvaluationDetails struct {
	// Object is the base model of ExpressionEvaluationDetails
	Object
	// ExpressionResult undocumented
	ExpressionResult *bool `json:"expressionResult,omitempty"`
	// Expression undocumented
	Expression *string `json:"expression,omitempty"`
	// ExpressionEvaluationDetails undocumented
	ExpressionEvaluationDetails []ExpressionEvaluationDetails `json:"expressionEvaluationDetails,omitempty"`
	// PropertyToEvaluate undocumented
	PropertyToEvaluate *PropertyToEvaluate `json:"propertyToEvaluate,omitempty"`
}