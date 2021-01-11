/*
 * Email Validator
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CheckIfEmailExistResult struct {
	Input string `json:"input,omitempty"`

	IsReachable string `json:"is_reachable,omitempty"`

	Misc CheckIfEmailExistMisc `json:"misc,omitempty"`

	Mx CheckIfEmailExistMx `json:"mx,omitempty"`

	Smtp CheckIfEmailExistSmtp `json:"smtp,omitempty"`

	Syntax CheckIfEmailExistSyntax `json:"syntax,omitempty"`

	Error string `json:"error,omitempty"`
}