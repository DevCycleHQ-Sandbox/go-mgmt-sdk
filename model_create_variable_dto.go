/*
 * DevCycle Management API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go_mgmt_sdk

type CreateVariableDto struct {
	// Variable name
	Name string `json:"name,omitempty"`
	// A description of the Variable
	Description string `json:"description,omitempty"`
	// Unique Variable identifier, can be used in the SDK / API to reference by key rather then ID. Must only contain lower-case characters and `_` or `-`.
	Key string `json:"key"`
	// The key or ID of the Feature this Variable is associated with
	Feature string `json:"_feature,omitempty"`
	// The type of Variable. Must be one of [String | Boolean | Number | JSON]
	Type_ string `json:"type"`
	// A default value for the Variable
	DefaultValue *interface{} `json:"defaultValue,omitempty"`
}