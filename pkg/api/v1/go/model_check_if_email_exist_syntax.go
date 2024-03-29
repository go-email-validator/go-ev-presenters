/*
 * Email Validator
 *
 * All timeouts are set in seconds with nanosecond precision. For example, 1.505402 is 1 second, 505 milliseconds and 402 microseconds.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CheckIfEmailExistSyntax struct {
	Address *string `json:"address,omitempty"`

	Domain string `json:"domain,omitempty"`

	IsValidSyntax bool `json:"is_valid_syntax,omitempty"`

	Username string `json:"username,omitempty"`
}
