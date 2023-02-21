package service

import (
	"Logi/api/model/pyl"
	"Logi/api/model/resp"
	"Logi/api/model/scm"
	"Logi/app"
	"Logi/helper"
	"github.com/labstack/echo/v4"
	"log"
)

func CreateMember(c echo.Context) resp.CreateMember {
	pylCreateMember := new(pyl.CreateMember)
	err := c.Bind(pylCreateMember)
	helper.Pie(err)

	checkMember := CheckMember(pylCreateMember.UserId)
	respCreateMember := resp.CreateMember{}
	respDataCreateMember := resp.DataCreateMember{}

	if checkMember == false {
		hash, _ := helper.HashPassword(pylCreateMember.Password)
		scmCreateMember := scm.Member{
			UserId:   pylCreateMember.UserId,
			Email:    pylCreateMember.Email,
			Role:     pylCreateMember.Role,
			Password: hash,
		}
		app.Instance.Create(&scmCreateMember)

		respDataCreateMember.Message = "Create member sukses"
		respCreateMember.Data = respDataCreateMember
		respCreateMember.Code = 200
		respCreateMember.Success = true
		log.Print("Create member sukses")
	} else {
		respDataCreateMember.Message = "Gagal create member, user id sudah terdaftar"
		respCreateMember.Data = respDataCreateMember
		respCreateMember.Code = 400
		respCreateMember.Success = false
		log.Print("Gagal create member, user id sudah terdaftar")
	}

	return respCreateMember
}

func GetAllMembers(c echo.Context) resp.GetAllMember {
	var members []scm.Member
	var respDataGetAllMember []resp.DataGetAllMember
	app.Instance.Find(&members)

	for x := 0; x < len(members); x++ {
		var userId = members[x].UserId
		var rMember = GetDataMember(userId)

		memberOneData := resp.DataGetAllMember{
			UserId: rMember.UserId,
			Email:  rMember.Email,
			Role:   rMember.Role,
		}
		respDataGetAllMember = append(respDataGetAllMember, memberOneData)
	}

	app.Instance.Find(&members)

	respGetAllMember := resp.GetAllMember{
		Success: true,
		Code:    200,
		Data:    respDataGetAllMember,
		Message: "Success get data all member",
	}

	return respGetAllMember
}

func CheckMember(userId string) bool {
	var members []scm.Member
	app.Instance.Where("user_id = ?", userId).Find(&members)

	if len(members) < 1 {
		return false
	} else {
		return true
	}
}

func GetDataMember(userId string) resp.DataGetAllMember {
	var member = scm.Member{}
	app.Instance.Where("user_id = ?", userId).Find(&member)

	respMemberData := resp.DataGetAllMember{
		UserId:  member.UserId,
		Email:   member.Email,
		Role:    member.Role,
		Profile: GetProfileMember(userId),
	}
	return respMemberData
}

func GetProfileMember(userId string) resp.ProfileDataGetAllMember {
	respProfileDataGetAllMember := resp.ProfileDataGetAllMember{}
	respProfileDataGetAllMember.FullName = "Aditia"
	return respProfileDataGetAllMember
}
