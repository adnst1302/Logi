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

	if checkMember == false {
		hash, _ := helper.HashPassword(pylCreateMember.Password)
		scmCreateMember := scm.Member{
			UserId:   pylCreateMember.UserId,
			Email:    pylCreateMember.Email,
			Role:     pylCreateMember.Role,
			Password: hash,
		}
		app.Instance.Create(&scmCreateMember)

		respCreateMember.Message = "Create member sukses"
		respCreateMember.Code = 200
		respCreateMember.Success = true
		log.Print("Create member sukses")
	} else {
		respCreateMember.Message = "Gagal create member, user id sudah terdaftar"
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
			UserId:  rMember.UserId,
			Email:   rMember.Email,
			Role:    rMember.Role,
			Profile: rMember.Profile,
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

func DetailMember(c echo.Context) resp.DetailMember {
	pylDetailMember := new(pyl.DetailMember)
	err := c.Bind(pylDetailMember)
	helper.Pie(err)

	checkMember := CheckMember(pylDetailMember.UserId)
	respDetailMember := resp.DetailMember{}
	if checkMember == false {
		respDetailMember.Message = "Failed, userId not found!!!"
		respDetailMember.Code = 400
		respDetailMember.Success = false
	} else {
		var rMember = GetDataMember(pylDetailMember.UserId)
		memberOneData := resp.DataGetAllMember{
			UserId:  rMember.UserId,
			Email:   rMember.Email,
			Role:    rMember.Role,
			Profile: GetProfileMember(pylDetailMember.UserId),
		}
		respDetailMember.Message = "Success get detail member."
		respDetailMember.Code = 200
		respDetailMember.Success = true
		respDetailMember.Data = memberOneData
	}
	return respDetailMember
}

func UpdateProfileMember(c echo.Context) resp.UpdateProfileMember {
	pylUpdateProfileMember := new(pyl.UpdateProfileMember)
	err := c.Bind(pylUpdateProfileMember)
	helper.Pie(err)

	checkMember := CheckMember(pylUpdateProfileMember.UserId)
	respUpdateProfileMember := resp.UpdateProfileMember{}

	if checkMember == false {
		respUpdateProfileMember.Code = 400
		respUpdateProfileMember.Message = "Gagal update member, userId tidak terdaftar!!!"
		respUpdateProfileMember.Succcess = false
	} else {
		SubDeleteProfileMember(pylUpdateProfileMember.UserId)
		scmMemberProfile := &scm.MemberProfile{
			UserId:         pylUpdateProfileMember.UserId,
			FullName:       pylUpdateProfileMember.FullName,
			PlaceOfBirth:   pylUpdateProfileMember.PlaceOfBirth,
			BirthDate:      pylUpdateProfileMember.BirthDate,
			PhoneNumber:    pylUpdateProfileMember.PhoneNumber,
			Address:        pylUpdateProfileMember.Address,
			EducationLevel: pylUpdateProfileMember.EducationLevel,
			Division:       pylUpdateProfileMember.Division,
			SignInDate:     pylUpdateProfileMember.SignInDate,
			StatusEmployee: pylUpdateProfileMember.StatusEmployee,
		}
		app.Instance.Create(scmMemberProfile)

		respUpdateProfileMember.Code = 200
		respUpdateProfileMember.Message = "Success update member."
		respUpdateProfileMember.Succcess = true
	}

	return respUpdateProfileMember
}

func DeleteMember(c echo.Context) resp.DeleteMember {
	pylDeleteMember := new(pyl.DeleteMember)
	err := c.Bind(pylDeleteMember)
	helper.Pie(err)
	checkMember := CheckMember(pylDeleteMember.UserId)
	respDeleteMember := resp.DeleteMember{}

	if checkMember == false {
		respDeleteMember.Code = 400
		respDeleteMember.Message = "Failed delete member data, userId not found!!!"
		respDeleteMember.Succcess = false
	} else {
		SubDeleteMember(pylDeleteMember.UserId)
		SubDeleteProfileMember(pylDeleteMember.UserId)
		respDeleteMember.Code = 200
		respDeleteMember.Message = "Success delete member."
		respDeleteMember.Succcess = true
	}

	return respDeleteMember
}

func SubDeleteMember(userId string) bool {
	var member scm.Member
	app.Instance.Where("user_id = ?", userId).Unscoped().Delete(&member)
	return true
}

func SubDeleteProfileMember(userId string) bool {
	var memberProfile scm.MemberProfile
	app.Instance.Where("user_id = ?", userId).Unscoped().Delete(&memberProfile)
	return true
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
	var memberProfile scm.MemberProfile
	app.Instance.Where("user_id = ?", userId).First(&memberProfile)
	respProfileDataGetAllMember := resp.ProfileDataGetAllMember{}
	respProfileDataGetAllMember.FullName = memberProfile.FullName
	respProfileDataGetAllMember.BirthDate = memberProfile.BirthDate
	respProfileDataGetAllMember.PhoneNumber = memberProfile.PhoneNumber
	respProfileDataGetAllMember.PlaceOfBirth = memberProfile.PlaceOfBirth
	respProfileDataGetAllMember.Address = memberProfile.Address
	return respProfileDataGetAllMember
}
