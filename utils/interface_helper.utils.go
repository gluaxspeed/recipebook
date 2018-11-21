package utils

func GetString(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	default:
		return ""
	}
}

func GetInt(i interface{}) int {
	switch i.(type) {
	case int:
		return i.(int)
	default:
		return 0
	}
}

func GetStringSlice(i interface{}) []string {
	switch i.(type) {
	case []interface{}:
		l := i.([]interface{})
		sl := make([]string, len(l))
		for i, item := range l {
			sl[i] = item.(string)
		}

		return sl
	default:
		return make([]string, 0)
	}
}
