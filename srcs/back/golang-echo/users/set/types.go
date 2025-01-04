package set

// UpdateField represents the configuration for updating a specific field
type UpdateFieldConfig struct {
	FieldName      string
	TableName      string
	ValidateTag    string
	MessageSuccess string
}

// UpdateFieldConfigs は全てのフィールド更新設定を保持する構造体
type UpdateFieldConfigs struct {
	Area      UpdateFieldConfig
	Gender    UpdateFieldConfig
	Sexuality UpdateFieldConfig
	LastName  UpdateFieldConfig
	FirstName UpdateFieldConfig
	SelfIntro UpdateFieldConfig
	BirthDate UpdateFieldConfig
	IsGps     UpdateFieldConfig
}

// GenericUpdateRequest is a generic request structure
type GenericUpdateRequest struct {
	Value interface{} `json:"value"`
}

// NewUpdateFieldConfigs は全てのフィールド更新設定を初期化して返す
func NewUpdateFieldConfigs() UpdateFieldConfigs {
	return UpdateFieldConfigs{
		Area: UpdateFieldConfig{
			FieldName:      "area",
			TableName:      "user_info",
			ValidateTag:    "required,area",
			MessageSuccess: "Area updated successfully",
		},
		Gender: UpdateFieldConfig{
			FieldName:      "gender",
			TableName:      "user_info",
			ValidateTag:    "required,oneof=male female",
			MessageSuccess: "Gender updated successfully",
		},
		Sexuality: UpdateFieldConfig{
			FieldName:      "sexuality",
			TableName:      "user_info",
			ValidateTag:    "required,oneof=male female male/female",
			MessageSuccess: "Sexuality updated successfully",
		},
		LastName: UpdateFieldConfig{
			FieldName:      "lastname",
			TableName:      "user_info",
			ValidateTag:    "required,name",
			MessageSuccess: "Last name updated successfully",
		},
		FirstName: UpdateFieldConfig{
			FieldName:      "firstname",
			TableName:      "user_info",
			ValidateTag:    "required,name",
			MessageSuccess: "First name updated successfully",
		},
		SelfIntro: UpdateFieldConfig{
			FieldName:      "self_intro",
			TableName:      "user_info",
			ValidateTag:    "required,self_intro",
			MessageSuccess: "Self introduction updated successfully",
		},
		BirthDate: UpdateFieldConfig{
			FieldName:      "birthdate",
			TableName:      "user_info",
			ValidateTag:    "required,birthdate",
			MessageSuccess: "Birth date updated successfully",
		},
		IsGps: UpdateFieldConfig{
			FieldName:      "is_gps",
			TableName:      "user_info",
			ValidateTag:    "required,boolean",
			MessageSuccess: "GPS setting updated successfully",
		},
	}
}
