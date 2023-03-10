/*
 * Competencies and Academic Standards Exchange (CASE) Service OpenAPI (YAML) Definition
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: lmattson@imsglobal.org
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// ImsxStatusInfoType - This is the container for the status code and associated information returned within the HTTP messages received from the Service Provider. For the CASE service this object will only be returned to provide information about a failed request i.e. it will NOT be in the payload for a successful request. See Appendix B for further information on the interpretation of the information contained within this class
type ImsxStatusInfoType struct {
	ImsxCodeMajor string `json:"imsx_codeMajor"`

	ImsxSeverity string `json:"imsx_severity"`

	// Model Primitive Datatype = String
	ImsxDescription string `json:"imsx_description,omitempty"`

	ImsxCodeMinor ImsxCodeMinorType `json:"imsx_codeMinor,omitempty"`
}
