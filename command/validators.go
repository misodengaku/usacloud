package command

import (
	"fmt"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
	"strings"
)

func ValidateInStrValues(fieldName string, object interface{}, allowValues ...string) []error {
	return schema.ValidateInStrValues(allowValues...)(fieldName, object)
}

func ValidateRequired(fieldName string, object interface{}) []error {
	if IsEmpty(object) {
		return []error{fmt.Errorf("%q: is required", fieldName)}
	}
	return []error{}
}

func ValidateSakuraID(fieldName string, object interface{}) []error {
	return schema.ValidateSakuraID()(fieldName, object)
}

func ValidateSetProhibited(fieldName string, object interface{}) []error {
	if !IsEmpty(object) {
		return []error{fmt.Errorf("%q: can't set on current context", fieldName)}
	}
	return []error{}
}

func ValidateConflicts(fieldName string, object interface{}, values map[string]interface{}) []error {

	if !IsEmpty(object) {
		for _, v := range values {
			if !IsEmpty(v) {
				keys := []string{}
				for k := range values {
					keys = append(keys, fmt.Sprintf("%q", k))
				}
				return []error{fmt.Errorf("%q: is conflict with %s", fieldName, strings.Join(keys, " or "))}
			}
		}
	}
	return []error{}

}

func ValidateConflictValues(fieldName string, object interface{}, values map[string]interface{}) []error {

	if !IsEmpty(object) {
		for _, v := range values {
			if !IsEmpty(v) {
				keys := []string{}
				for k := range values {
					keys = append(keys, fmt.Sprintf("%q", k))
				}
				return []error{fmt.Errorf("%q(%#v): is conflict with %s", fieldName, object, strings.Join(keys, " or "))}
			}
		}
	}
	return []error{}

}

func ValidateBetween(fieldName string, object interface{}, min int, max int) []error {

	if object == nil {
		object = []int64{}
	}

	isSlice := func(object interface{}) bool {
		_, ok1 := object.([]int64)
		_, ok2 := object.([]string)

		return ok1 || ok2
	}

	if isSlice(object) {
		sliceLen := 0
		if s, ok := object.([]int64); ok {
			sliceLen = len(s)
		} else {
			s := object.([]string)
			sliceLen = len(s)
		}

		if max <= 0 {
			if sliceLen < min {
				return []error{fmt.Errorf("%q: slice length must be %d or more", fieldName, min)}
			}
		} else {
			if !(min <= sliceLen && sliceLen <= max) {
				return []error{fmt.Errorf("%q: slice length must be beetween %d and %d", fieldName, min, max)}
			}

		}
	}

	return []error{}
}

func ValidateOutputOption(o output.Option) []error {

	outputType := o.GetOutputType()
	columns := o.GetColumn()
	format := o.GetFormat()
	formatFile := o.GetFormatFile()
	quiet := o.GetQuiet()

	// format and format-file
	if format != "" && formatFile != "" {
		return []error{fmt.Errorf("%q: can't set with --format-file", "--format")}
	}

	// format(or format-file) with output-type
	if outputType != "" && format != "" {
		return []error{fmt.Errorf("%q: can't set with --output-type", "--format")}
	}
	if outputType != "" && formatFile != "" {
		return []error{fmt.Errorf("%q: can't set with --output-type", "--format-file")}
	}

	if formatFile != "" {
		errs := schema.ValidateFileExists()("--format-file", formatFile)
		if len(errs) > 0 {
			return errs
		}
	}

	if outputType != "" && quiet {
		return []error{fmt.Errorf("%q: can't set with --output-type", "--quiet")}
	}

	if outputType != "tsv" && outputType != "csv" && len(columns) > 0 {
		return []error{fmt.Errorf("%q: can't set when --output-type is csv/tsv", "--column")}
	}

	return []error{}

}
