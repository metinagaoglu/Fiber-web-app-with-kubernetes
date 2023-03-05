package validations

func GetBroadcastValidationRules() map[string]interface{} {
	return map[string]interface{}{
		"room": "required,min=3,max=6",
		"message": "required,min=2,max=20",
	}
}
