package obj

import "reflect"

type Entry struct {
	Key   string
	Value interface{}
}

// Keys returns the field names of a struct as a slice of strings.
func Keys(s interface{}) []string {
	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil
	}

	keys := make([]string, 0, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		keys = append(keys, t.Field(i).Name)
	}

	return keys
}

// Values returns the field values of a struct
func Values(s interface{}) []interface{} {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil
	}

	values := make([]interface{}, 0, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values = append(values, v.Field(i).Interface())
	}

	return values
}

// Entries returns a slice of Entry for a struct
func Entries(s interface{}) []Entry {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
		t = t.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	entries := make([]Entry, 0, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		entries = append(entries, Entry{
			Key:   t.Field(i).Name,
			Value: v.Field(i).Interface(),
		})
	}
	return entries
}

// EntriesFlat returns JS-style 2D array [key, value] for a struct
func EntriesFlat(s interface{}) [][2]interface{} {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
		t = t.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	entries := make([][2]interface{}, 0, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		entries = append(entries, [2]interface{}{
			t.Field(i).Name,
			v.Field(i).Interface(),
		})
	}

	return entries
}

// FromEntries converts []Entry to map[string]interface{}
func FromEntries(entries []Entry) map[string]interface{} {
	result := make(map[string]interface{}, len(entries))

	for _, entry := range entries {
		result[entry.Key] = entry.Value
	}
	return result
}

// Pick selects specified keys from a struct and returns a map
func Pick(obj interface{}, keys ...string) map[string]interface{} {
	keySet := make(map[string]struct{}, len(keys))

	// create emmpty Hash Map instead of using slice string[]
	for _, key := range keys {
		keySet[key] = struct{}{} // empty struct
	}

	entries := Entries(obj)
	result := make(map[string]interface{}, len(keys)) // based on the picked up size

	for _, entry := range entries {

		// _, found := keySet[entry.Key]

		// if found {
		// 	result[entry.Key] = entry.Value
		// }

		// shorthand for code above
		if _, found := keySet[entry.Key]; found {
			result[entry.Key] = entry.Value
		}
	}

	return result
}

// Omit removes specified keys from a struct and returns a map
func Omit(obj interface{}, keys ...string) map[string]interface{} {
	keySet := make(map[string]struct{}, len(keys))

	for _, key := range keys {
		keySet[key] = struct{}{}
	}

	entries := Entries(obj)

	result := make(map[string]interface{}, len(entries))

	for _, entry := range entries {
		if _, found := keySet[entry.Key]; !found {
			result[entry.Key] = entry.Value
		}
	}
	return result
}

// Merge merges two structs/maps into a single map
func Merge(a, b interface{}) map[string]interface{} {
	// optional future implementation
	return nil
}
