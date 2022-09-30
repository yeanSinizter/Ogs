package usermodel

type User struct {
	Id         int    `json:"id" param:"id" query:"id" form:"id"`
	FirstName  string `json:"first_name" query:"first_name"`
	LastName   string `json:"last_name" query:"last_name"`
	NamePrefix string `json:"name_prefix" query:"name_prefix"`
	SumGrade   string `json:"sum_grade" query:"sum_grade"`
	AGrade     string `json:"a_grade" query:"a_grade"`
	BGrade     string `json:"b_grade" query:"b_grade"`
	CGrade     string `json:"c_grade" query:"c_grade"`
}

func (User) TableName() string {
	return "users"
}
