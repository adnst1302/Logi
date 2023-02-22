package app

import "Logi/api/model/scm"

func CheckMember(userId string) bool {
	var members []scm.Member
	Instance.Where("user_id = ?", userId).Find(&members)

	if len(members) < 1 {
		return false
	} else {
		return true
	}

}

func DeleteMember(userId string) bool {
	var member scm.Member
	Instance.Where("user_id = ?", userId).Unscoped().Delete(&member)
	return true
}

func DeleteProfileMember(userId string) bool {
	var memberProfile scm.MemberProfile
	Instance.Where("user_id = ?", userId).Unscoped().Delete(&memberProfile)
	return true
}
