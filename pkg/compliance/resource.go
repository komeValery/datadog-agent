// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

package compliance

// Resource describes supported resource types observed by a Rule
type Resource struct {
	File    *File           `yaml:"file,omitempty"`
	Process *Process        `yaml:"process,omitempty"`
	Group   *Group          `yaml:"group,omitempty"`
	Command *Command        `yaml:"command,omitempty"`
	Audit   *Audit          `yaml:"audit,omitempty"`
	Docker  *DockerResource `yaml:"docker,omitempty"`
	API     *API            `yaml:"api,omitempty"`
}

// File describes a file resource
type File struct {
	Path     string    `yaml:"path,omitempty"`
	PathFrom ValueFrom `yaml:"pathFrom,omitempty"`
	Glob     string    `yaml:"glob,omitempty"`

	Filter []Filter `yaml:"filter,omitempty"`

	Report Report `yaml:"report,omitempty"`
}

// Process describes a process resource
type Process struct {
	Name string `yaml:"name"`

	Filter []Filter `yaml:"filter,omitempty"`

	Report Report `yaml:"report,omitempty"`
}

// Group describes a group membership resource
type Group struct {
	Name string `yaml:"name"`

	Filter []Filter `yaml:"filter,omitempty"`

	Report Report `yaml:"report,omitempty"`
}

// Command describes a command resource usually reporting exit code or output
type Command struct {
	Run string `yaml:"run"`

	// TODO: generalize to use the same filter types
	Filter []CommandFilter `yaml:"filter,omitempty"`

	Report Report `yaml:"report,omitempty"`
}

// Audit describes an audited file resource
type Audit struct {
	Path     string    `yaml:"path,omitempty"`
	PathFrom ValueFrom `yaml:"pathFrom,omitempty"`

	Filter []Filter `yaml:"filter,omitempty"`

	Report Report `yaml:"report,omitempty"`
}

// DockerResource describes a resource from docker daemon
type DockerResource struct {
	Kind string `yaml:"kind"`

	Filter []Filter `yaml:"filter,omitempty"`

	Report Report `yaml:"report,omitempty"`
}

// API describes a generic API query resource
type API struct {
	Kind string `yaml:"kind"`
	Get  string `yaml:"get,omitempty"`

	Vars APIVars `yaml:"vars,omitempty"`

	Filter []Filter `yaml:"filter,omitempty"`

	Report Report `yaml:"report,omitempty"`
}

// ValueFrom provides a lookup list for substitution of a value in a Resource
type ValueFrom []ValueSource

// ValueSource provides a single lookup option for value substitution in a Resource
type ValueSource struct {
	Command string           `yaml:"command,omitempty"`
	File    ValueFromFile    `yaml:"file,omitempty"`
	Process ValueFromProcess `yaml:"process,omitempty"`
}

// ValueFromFile describes a value taken from properties of a file
type ValueFromFile struct {
	Path     string `yaml:"path"`
	Property string `yaml:"property"`
}

// ValueFromProcess describes a value taken from attributes of a process
type ValueFromProcess struct {
	Name string `yaml:"name"`
	Flag string `yaml:"flag"`
}

// Report defines a set of reported fields which are sent in a RuleEvent
type Report []ReportedField

const (
	// PropertyKindAttribute describes an attribute
	PropertyKindAttribute = "attribute"

	// PropertyKindJSONPath describes a JSONPath query
	PropertyKindJSONPath = "jsonpath"

	// PropertyKindFlag describes a process flag
	PropertyKindFlag = "flag"

	// PropertyKindTemplate describes a template
	PropertyKindTemplate = "template"
)

// ReportedField defines options for reporting various attributes of observed resources
type ReportedField struct {
	Property string `yaml:"property,omitempty"`
	Kind     string `yaml:"kind,omitempty"`
	As       string `yaml:"as,omitempty"`
	Value    string `yaml:"value,omitempty"`
}

// CommandFilter specifies filtering options to include or exclude a Command from reporting
type CommandFilter struct {
	Include *CommandCondition `yaml:"include,omitempty"`
	Exclude *CommandCondition `yaml:"exclude,omitempty"`
}

// CommandCondition specifies conditions to include or exclude a Command from reporting
type CommandCondition struct {
	ExitCode *int `yaml:"exitCode,omitempty"`
}

// Filter specifies filtering options to include or exclude a Docker resource from reporting
type Filter struct {
	Include *Condition `yaml:"include,omitempty"`
	Exclude *Condition `yaml:"exclude,omitempty"`
}

const (
	// OpExists defines an operation that checks for property presence
	OpExists = "exists"
	// OpEqual defines an operation that checks for property equality
	OpEqual = "equal"
)

// Condition defines a filter condition
type Condition struct {
	Operation string `yaml:"op,omitempty"`
	Property  string `yaml:"property,omitempty"`
	Kind      string `yaml:"kind,omitempty"`
	Value     string `yaml:"value,omitempty"`
}

// APIVars defines a list of variables substituted in generic API resource endpoint queries
type APIVars []APIVar

// APIVar defines a variable substitution in generic API resource endpoint query
type APIVar struct {
	Name  string       `yaml:"name"`
	List  *APIVarValue `yaml:"enumerate,omitempty"`
	Value *APIVarValue `yaml:"value,omitempty"`
}

// APIVarValue defines how an API variable is retrieved
type APIVarValue struct {
	Get      string `yaml:"get"`
	JSONPath string `yaml:"jsonpath"`
}