/*
 * Email Validator
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CheckIfEmailExistSyntax struct {
	Address string `json:"address,omitempty"`

	Domain string `json:"domain,omitempty"`

	IsValidSyntax bool `json:"is_valid_syntax,omitempty"`

	Username string `json:"username,omitempty"`
}
