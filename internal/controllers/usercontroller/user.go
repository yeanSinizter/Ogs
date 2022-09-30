package usercontroller

import (
	"golang/configs/dbconfig"
	"golang/configs/responseconfig"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CalGrade(grade string) int {
	var score = 0
	if grade == "A" {
		score = 4
	} else if grade == "B" {
		score = 2
	} else if grade == "C" {
		score = 1
	} else if grade == "D" {
		score = 0
	}
	return score
}

func GetUser(c echo.Context) error {
	users := []User{}
	filter := new(Filter)
	err := c.Bind(filter)
	if err != nil {
		log.Println(err.Error())
	}

	// h.DB.Debug().Where("first_name LIKE ? OR first_name IS NULL AND last_name LIKE ?", "%"+searchFirstName+"%", "%"+searchLastName+"%").Find(&users)
	// OR columname IS NULL คือตัวอย่างแต่ไม่ควรทำ **

	dbconfig.MySQL.Debug().Where("first_name LIKE ?", "%"+filter.FirstName+"%").
		Where("last_name LIKE ? ", "%"+filter.LastName+"%").
		Where("sum_grade LIKE ? ", "%"+filter.SumGrade+"%").
		Find(&users)

	return responseconfig.Handler(c).Success(users)
}

func GetUserParam(c echo.Context) error {
	users := User{}
	userId := c.Param("userId")
	getUserId, err := strconv.Atoi(userId)
	if err != nil {
		return responseconfig.Handler(c).BadRequest("GetUserParam")

	}
	users.Id = getUserId
	dbconfig.MySQL.Find(&users)
	return responseconfig.Handler(c).Success(users)

}

func CreateUser(c echo.Context) error {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		return responseconfig.Handler(c).BadRequest("CreateUser")
	}
	var a_grade = CalGrade(user.AGrade)
	var b_grade = CalGrade(user.BGrade)
	var c_grade = CalGrade(user.CGrade)
	var sum_grade = a_grade + b_grade + c_grade

	var score = ""
	if sum_grade >= 9 && sum_grade <= 12 {
		score = "A"
	} else if sum_grade >= 5 && sum_grade <= 8 {
		score = "B"
	} else if sum_grade >= 2 && sum_grade <= 4 {
		score = "C"
	} else {
		score = "D"
	}
	log.Println(user.SumGrade)
	user.SumGrade = score
	dbconfig.MySQL.Create(&user).Debug()
	return responseconfig.Handler(c).Success(user)
}

func UpdateData(c echo.Context) error {
	users := User{}
	userId := c.Param("userId")
	err := c.Bind(&users)
	if err != nil {
		return responseconfig.Handler(c).BadRequest("UpdateData")

	}
	getUserId, err := strconv.Atoi(userId)
	if err != nil {
		return c.String(400, err.Error())
	}
	var a_grade = CalGrade(users.AGrade)
	var b_grade = CalGrade(users.BGrade)
	var c_grade = CalGrade(users.CGrade)
	var sum_grade = a_grade + b_grade + c_grade

	var score = ""
	if sum_grade >= 9 && sum_grade <= 12 {
		score = "A"
	} else if sum_grade >= 5 && sum_grade <= 8 {
		score = "B"
	} else if sum_grade >= 2 && sum_grade <= 4 {
		score = "C"
	} else {
		score = "D"
	}
	users.SumGrade = score
	users.Id = getUserId
	dbconfig.MySQL.Updates(&users).Debug()
	return responseconfig.Handler(c).Success("UpdateData")

}
func DeleteData(c echo.Context) error {
	users := User{}
	userId := c.Param("userId")
	getUserId, err := strconv.Atoi(userId)
	if err != nil {
		return responseconfig.Handler(c).BadRequest("DeleteData")

	}
	users.Id = getUserId
	dbconfig.MySQL.Delete(&users).Debug()
	return responseconfig.Handler(c).Success("DeleteData")

}

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

type Filter struct {
	BGrade    string `json:"b_grade" query:"b_grade"`
	SumGrade  string `json:"sum_grade" query:"sum_grade"`
	FirstName string `json:"first_name" query:"first_name"`
	LastName  string `json:"last_name" query:"last_name"`
}

type Sumgrade struct {
	FirstName string `json:"first_name"`
	SumGrade  string `json:"sum_grade"`
	AGrade    string `json:"a_grade"`
	BGrade    string `json:"b_grade"`
	CGrade    string `json:"c_grade"`
}
