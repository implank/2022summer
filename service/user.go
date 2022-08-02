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

func UpdateUser(user *model.User) (err error) {
	err = global.DB.Save(user).Error
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

func CreateGroup(group *model.Group) (err error) { // 先创建 Group, 再创建 Identity
	if err = global.DB.Create(&group).Error; err != nil {
		return err
	}
	identity := model.Identity{UserID: group.UserID, GroupID: group.GroupID, Status: 3}
	if err = CreateIdentity(&identity); err != nil {
		return err
	}
	return
}

func UpdateGroup(group *model.Group) (err error) {
	err = global.DB.Save(group).Error
	return err
}

func DeleteGroup(group *model.Group) (err error) { // 先删除 Identity, 再删除 Group
	global.DB.Where("group_id = ?", group.GroupID).Delete(model.Identity{})
	err = global.DB.Delete(&group).Error
	return err
}

func QueryIdentity(userID uint64, groupID uint64) (identity model.Identity, notFound bool) {
	notFound = global.DB.Where("user_id = ? and group_id = ?", userID, groupID).First(&identity).RecordNotFound()
	return identity, notFound
}

func GetUserHasGroups(userID uint64) (groups []model.Group, notFound bool) {
	notFound = global.DB.Raw(
		"SELECT * FROM identities, `groups` "+
			"WHERE ? = identities.user_id "+
			"AND `groups`.group_id = identities.group_id;", userID).Find(&groups).RecordNotFound()
	return groups, notFound
}

func GetGroupMembers(groupID uint64) (users []model.GroupMembers, notFound bool) {
	// GroupMembers 比 User 多了一个身份, 少了一些其他信息, 可以 "直观地查看到成员的昵称、真实姓名、邮箱、身份"
	notFound = global.DB.Raw(
		"SELECT users.user_id AS user_id, username, real_name, email, status "+
			"FROM identities, users "+
			"WHERE users.user_id = identities.user_id "+
			"AND ? = identities.group_id", groupID).Find(&users).RecordNotFound()
	return users, notFound
}

func CreateIdentity(identity *model.Identity) (err error) {
	if err = global.DB.Create(&identity).Error; err != nil {
		return err
	}
	return
}

func UpdateIdentity(identity *model.Identity) (err error) {
	err = global.DB.Save(identity).Error
	return err
}

func DeleteIdentity(identity *model.Identity) (err error) {
	err = global.DB.Delete(&identity).Error
	return err
}
