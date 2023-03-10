// format will handle formating of diffrent kinds.
package format

import (
	"fmt"
	"reflect"
	"strings"
)

// FormatProm will format the struct according to the label tag and return a string in the format promethues expects.
func FormatProm(obj any) (string, error) {
	rv := reflect.ValueOf(obj)

	// Make sure we get only pointer.
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return "", fmt.Errorf("invalid type %s", reflect.TypeOf(obj))
	}

	tagsMapping := make(map[string]string)
	CollectTagsRec(tagsMapping, reflect.ValueOf(obj))
	if _, ok := tagsMapping["metric_name"]; !ok {
		return "", fmt.Errorf("missing metric_name  tag in struct")
	}
	if _, ok := tagsMapping["metric_value"]; !ok {
		return "", fmt.Errorf("missing metric_value tag in struct")
	}

	var res string

	// Add all the common labels to the string.
	for labelName, value := range tagsMapping {
		if labelName != "metric_name" && labelName != "metric_value" {
			res = fmt.Sprintf("%s %s=%q", res, labelName, value)
		}
	}
	//
	res = strings.TrimSpace(res)

	// Add the metric name and value to the result in the format promethues expects.
	res = fmt.Sprintf("%s{%s} %s", formatString(tagsMapping["metric_name"]), res, tagsMapping["metric_value"])
	return res, nil
}

// CollectTagsRec will collect the values of the fields recursvly from
// the struct according the label tag.
func CollectTagsRec(tagsMapping map[string]string, val reflect.Value) {

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		switch f.Kind() {
		case reflect.Struct:

			// Get to the underlying fields of the struct.
			CollectTagsRec(tagsMapping, f)
		case reflect.Map:

			// Go over all the map add add the key and value.
			for _, key := range f.MapKeys() {
				k := key.String()
				val := f.MapIndex(key).String()
				k = formatString(k)
				tagsMapping[k] = val
			}
		case reflect.Slice:
			// Loop every member of the slice
			for j := 0; j < f.Len(); j++ {
				CollectTagsRec(tagsMapping, f.Index(i))
			}

			// Assume we only send those types.
		case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
			reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32,
			reflect.Float64, reflect.String:

			label := val.Type().Field(i).Tag.Get("label")

			// If the label is empty skip
			if label == "" {
				continue
			}
			value := fmt.Sprintf("%v", val.Field(i).Interface())

			// Add to the mapping.
			tagsMapping[formatString(label)] = value

		}

	}

}
