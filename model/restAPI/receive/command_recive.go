/*
 * air_condition
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package receive

type CommandRecive struct {
	Command string `json:"command"`

	Args []string `json:"args"`
}