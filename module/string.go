package module

import (
	"strings"
	"github.com/ekilie/vint-lang/object"
)

var StringFunctions = map[string]object.ModuleFunction{}

func init() {
	StringFunctions["trim"] = trim
	StringFunctions["contains"] = contains
	StringFunctions["toUpper"] = toUpper
	StringFunctions["toLower"] = toLower
	StringFunctions["replace"] = replace
	StringFunctions["split"] = split
	StringFunctions["join"] = join
	StringFunctions["substring"] = substring
	StringFunctions["length"] = length
	StringFunctions["indexOf"] = indexOf
}

// trim removes leading and trailing whitespaces
func trim(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "We need one argument: the string to trim"}
	}

	str := args[0].Inspect()
	return &object.String{Value: strings.TrimSpace(str)}
}

// contains checks if the substring exists in the string
func contains(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "We need two arguments: the string and the substring"}
	}

	str := args[0].Inspect()
	substr := args[1].Inspect()
	return &object.Boolean{Value: strings.Contains(str, substr)}
}

// toUpper converts the string to uppercase
func toUpper(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "We need one argument: the string to convert"}
	}

	str := args[0].Inspect()
	return &object.String{Value: strings.ToUpper(str)}
}

// toLower converts the string to lowercase
func toLower(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "We need one argument: the string to convert"}
	}

	str := args[0].Inspect()
	return &object.String{Value: strings.ToLower(str)}
}

// replace replaces occurrences of the old substring with the new one
func replace(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "We need three arguments: the string, the substring to replace, and the new substring"}
	}

	str := args[0].Inspect()
	old := args[1].Inspect()
	new := args[2].Inspect()

	return &object.String{Value: strings.ReplaceAll(str, old, new)}
}

// split splits a string into a slice based on the delimiter
func split(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "We need two arguments: the string and the delimiter"}
	}

	str := args[0].Inspect()
	delimiter := args[1].Inspect()

	// Split the string by the delimiter
	parts := strings.Split(str, delimiter)
	// Convert the parts into a list of strings
	list := &object.Array{}
	for _, part := range parts {
		list.Elements = append(list.Elements, &object.String{Value: part})
	}

	return list
}

// join joins a slice of strings into a single string with a delimiter
func join(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "We need two arguments: the list of strings and the delimiter"}
	}

	list := args[0].(*object.Array)
	delimiter := args[1].Inspect()

	var result string
	for i, elem := range list.Elements {
		if i > 0 {
			result += delimiter
		}
		result += elem.(*object.String).Value
	}

	return &object.String{Value: result}
}

// substring extracts a substring from the string
func substring(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "We need three arguments: the string, the start index, and the end index"}
	}

	str := args[0].Inspect() // Get the string value
	start := args[1].(*object.Integer).Value // Get the start index as int64
	end := args[2].(*object.Integer).Value // Get the end index as int64

	// Convert int64 to int for string indexing
	startIdx := int(start)
	endIdx := int(end)

	// Ensure indices are valid
	if startIdx < 0 || endIdx > len(str) || startIdx >= endIdx {
		return &object.Error{Message: "Invalid start or end index"}
	}

	return &object.String{Value: str[startIdx:endIdx]}
}


// length returns the length of the string
func length(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "We need one argument: the string to measure"}
	}

	str := args[0].Inspect()
	return &object.Integer{Value: int64(len(str))}
}

// indexOf finds the index of the first occurrence of a substring
func indexOf(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "We need two arguments: the string and the substring to find"}
	}

	str := args[0].Inspect()
	substr := args[1].Inspect()

	index := strings.Index(str, substr)
	if index == -1 {
		return &object.Error{Message: "Substring not found"}
	}

	return &object.Integer{Value: int64(index)}
}