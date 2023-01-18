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

// LinkGenUriType - A container for the information that is used to achieve the link data reference.
type LinkGenUriType struct {

	// Model Primitive Datatype = NormalizedString
	Title string `json:"title"`

	// Model Primitive Datatype = NormalizedString
	Identifier string `json:"identifier"`

	// Model Primitive Datatype = AnyURI
	Uri string `json:"uri"`
}