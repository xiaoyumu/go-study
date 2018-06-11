package commandline

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// DefaultParameterPrefix identify the parameter in os.Args slice
const DefaultParameterPrefix = "-"

// DefaultNameValueDelimiter separates the name and value from
// a command line paramter, default is :
const DefaultNameValueDelimiter string = ":"

const minimumDelimiterIndex = 2

// ParameterPool contains the command line parameters
type ParameterPool struct {
	parameters map[string]string
	setting    *ParameterParseSetting
}

// ParameterParseSetting contains setting that used
// while parsing command line parameters
type ParameterParseSetting struct {
	Prefix                   string
	Delimiter                string
	RequiredParameters       []string
	ActionOnValidationFailed func() // When parameter validation failed, the action will be executed.
}

// ParameterError contains multiple parameter validation errors
type ParameterError struct {
	errors []string
}

// CreateDefault the ParameterPool instance and return the pointer
// with default parameter prefix and delimiter
func CreateDefault() (*ParameterPool, error) {
	return Create(&ParameterParseSetting{
		Prefix:    DefaultParameterPrefix,
		Delimiter: DefaultNameValueDelimiter,
	})
}

// Create the ParameterPool instance and return the pointer with
// given parameter parse setting
func Create(s *ParameterParseSetting) (*ParameterPool, error) {
	if s == nil {
		panic("parameter s cannot be nil while creating a ParameterPool")
	}

	if len(s.Prefix) <= 0 {
		s.Prefix = DefaultParameterPrefix
	}

	if len(s.Delimiter) <= 0 {
		s.Delimiter = DefaultNameValueDelimiter
	}

	pool := ParameterPool{
		parameters: getParameters(s.Prefix, s.Delimiter),
		setting:    s,
	}

	err := pool.validate()
	if err != nil {
		if s.ActionOnValidationFailed != nil {
			s.ActionOnValidationFailed()
		}

		return nil, err
	}

	return &pool, nil
}

func (p *ParameterPool) validate() error {
	// If no required parameter specified
	// just skip the validation
	if len(p.setting.RequiredParameters) <= 0 {
		return nil
	}

	parameterError := ParameterError{
		errors: []string{},
	}
	for _, requiredParameter := range p.setting.RequiredParameters {
		if !p.HasParameter(requiredParameter) {
			parameterError.appendError(
				fmt.Sprintf("the parameter %s is required, but not present", requiredParameter))
		}
	}

	if parameterError.HasError() {
		return parameterError
	}
	return nil
}

// The first arg of os.Args is the program itself.
func getParameters(prefix string, delimiter string) map[string]string {

	parameterPool := make(map[string]string)

	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, prefix) {
			continue
		}

		nameValueDelimiterIndex := strings.Index(arg, delimiter)

		// The parameter should be like -h:somehost
		// so the minimum index of the name value delimiter should be at
		if nameValueDelimiterIndex <= minimumDelimiterIndex {
			continue
		}

		parameterName := arg[1:nameValueDelimiterIndex]

		parameterValue := arg[nameValueDelimiterIndex+1:]

		// Ignore empty parameter value
		if len(parameterValue) <= 0 {
			continue
		}
		parameterPool[parameterName] = parameterValue
	}

	return parameterPool
}

// GetParameterValueString checks the parameter 'name' in the pool, return its value,
// And if not found, return the defaultValue
func (p *ParameterPool) GetParameterValueString(name string, defaultValue string) string {
	if p == nil {
		panic("parameter pool must not be nil")
	}

	var value string
	if parameterValue, exists := p.parameters[name]; !exists {
		log.Printf("(ParaParser) Parameter [%s] was not found, using default value [%s]", name, defaultValue)
		value = defaultValue
	} else {
		value = parameterValue
	}

	return value
}

// DumpParameters write all parameters along with its values to log from the ParameterPool
func (p *ParameterPool) DumpParameters() {
	log.Printf("Dumping parameters from the pool...")
	p.Iterate(func(parameterName string, parameterValue string) {
		log.Printf("(ParaParser) Parameter: %s = %s", parameterName, parameterValue)
	})
}

// Iterate parameter pool and retrieve all parameter name and its value
func (p *ParameterPool) Iterate(action func(string, string)) {
	if action == nil {
		return
	}

	for parameterName, parameterValue := range p.parameters {
		action(parameterName, parameterValue)
	}
}

// Count how mange parameters in the pool
func (p *ParameterPool) Count() int {
	return len(p.parameters)
}

// HasParameter checks if a given parameter exists in the pool
func (p *ParameterPool) HasParameter(name string) bool {
	for parameterName := range p.parameters {
		if parameterName == name {
			return true
		}
	}

	return false
}

func (e ParameterError) Error() string {
	errorMessages := make([]string, len(e.errors))

	for i, err := range e.errors {
		errorMessages[i] = err
	}

	return strings.Join(errorMessages, "\n")
}

// HasError checks if the ParameterError has any error
func (e *ParameterError) HasError() bool {
	return len(e.errors) > 0
}

func (e *ParameterError) appendError(errMsg string) {
	e.errors = append(e.errors, errMsg)
}
