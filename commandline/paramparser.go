package commandline

import (
	"log"
	"os"
	"strings"
)

// ParameterPool contains the command line parameters
type ParameterPool struct {
	parameters map[string]string
}

// Create the ParameterPool instance and return the pointer
func Create() *ParameterPool {
	pool := ParameterPool{
		parameters: getParameters(),
	}

	return &pool
}

// The first arg of os.Args is the program itself.
func getParameters() map[string]string {

	const argumentPrefix = "-"
	const nameValueDelimiter string = ":"
	const minimumDelimiterIndex = 2

	parameterPool := make(map[string]string)

	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, argumentPrefix) {
			continue
		}

		nameValueDelimiterIndex := strings.Index(arg, nameValueDelimiter)

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
		log.Printf("%s was not found, using default value %s", name, defaultValue)
		value = defaultValue
	} else {
		value = parameterValue
	}

	return value
}
