package commandline

import (
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
}

// CreateDefault the ParameterPool instance and return the pointer
// with default parameter prefix and delimiter
func CreateDefault() *ParameterPool {
	return Create(DefaultParameterPrefix, DefaultNameValueDelimiter)
}

// Create the ParameterPool instance and return the pointer with
// given parameter prefix and delimiter
func Create(prefix string, delimiter string) *ParameterPool {
	pool := ParameterPool{
		parameters: getParameters(prefix, delimiter),
	}

	return &pool
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
	for parameterName, parameterValue := range p.parameters {
		log.Printf("(ParaParser) Parameter: %s = %s", parameterName, parameterValue)
	}
}
