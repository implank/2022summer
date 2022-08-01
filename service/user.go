package service

import (
	"2022summer/global"
	"2022summer/model"
)

func QueryUserByUsername(username string) (user model.User, notFound bool) {
	notFound = global.DB.Where("username = ?", username).First(&user).RecordNotFound()
	return user, notFound
}

func QueryUserByUserID(userID uint64) (user model.User, notFound bool) {
	notFound = global.DB.Where("user_id = ?", userID).First(&user).RecordNotFound()
	return user, notFound
}

func QueryUserByEmail(email string) (user model.User, notFound bool) {
	notFound = global.DB.Where("email = ?", email).First(&user).RecordNotFound()
	return user, notFound
}

func CreateUser(user *model.User) (err error) {
	if err = global.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}

func UpdateUser(user *model.User) error {
	err := global.DB.Save(user).Error
	return err
}

func QueryGroupByGroupName(groupName string) (group model.Group, notFound bool) {
	notFound = global.DB.Where("group_name = ?", groupName).First(&group).RecordNotFound()
	return group, notFound
}

func QueryGroupByGroupID(groupID uint64) (group model.Group, notFound bool) {
	notFound = global.DB.Where("group_id = ?", groupID).First(&group).RecordNotFound()
	return group, notFound
}

func CreateGroup(group *model.Group) (err error) {
	if err = global.DB.Create(&group).Error; err != nil {
		return err
	}
	return
}

func UpdateGroup(group *model.Group) error {
	err := global.DB.Save(group).Error
	return err
}

func QueryIdentity(userID uint64, groupID uint64) (identity model.Identity, notFound bool) {
	notFound = global.DB.Where("user_id = ? and group_id = ?", userID, groupID).First(&identity).RecordNotFound()
	return identity, notFound
}

func GetUserHasGroups(userID uint64) (groups []model.Group, notFound bool) {
	global.DB.Raw(
		"SELECT * FROM identities, `groups` "+
			"WHERE ? = identities.user_id "+
			"AND `groups`.group_id = identities.group_id;", userID).Find(&groups).RecordNotFound()
	return groups, notFound
}

func GetGroupMembers(groupID uint64) (users []model.User, notFound bool) {
	global.DB.Raw(
		"SELECT * FROM identities, users "+
			"WHERE users.user_id = identities.user_id "+
			"AND ? = identities.group_id;", groupID).Find(&users).RecordNotFound()
	return users, notFound
}
