package service

import (
	"2022summer/global"
	"2022summer/model/database"
)

func QueryGroupByGroupName(groupName string) (group database.Group, notFound bool) {
	notFound = global.DB.Where("group_name = ?", groupName).First(&group).RecordNotFound()
	return group, notFound
}

func QueryGroupByGroupID(groupID uint64) (group database.Group, notFound bool) {
	notFound = global.DB.Where("group_id = ?", groupID).First(&group).RecordNotFound()
	return group, notFound
}

func CreateGroup(group *database.Group) (err error) { // 先创建 Group, 再创建 Identity
	if err = global.DB.Create(&group).Error; err != nil {
		return err
	}
	identity := database.Identity{UserID: group.UserID, GroupID: group.GroupID, Status: 3}
	if err = CreateIdentity(&identity); err != nil {
		return err
	}
	return
}

func UpdateGroup(group *database.Group) (err error) {
	err = global.DB.Save(group).Error
	return err
}

func DeleteGroup(group *database.Group) (err error) { // 先删除 Identity, 再删除 Group
	global.DB.Where("group_id = ?", group.GroupID).Delete(database.Identity{})
	err = global.DB.Delete(&group).Error
	return err
}

func QueryIdentity(userID uint64, groupID uint64) (identity database.Identity, notFound bool) {
	notFound = global.DB.Where("user_id = ? and group_id = ?", userID, groupID).First(&identity).RecordNotFound()
	return identity, notFound
}

func GetUserHasGroups(userID uint64) (groups []database.Group) {
	global.DB.Raw(
		"SELECT * FROM identities, `groups` "+
			"WHERE ? = identities.user_id "+
			"AND `groups`.group_id = identities.group_id;", userID).Find(&groups).RecordNotFound()
	return groups
}

func GetGroupMembers(groupID uint64) (users []database.GroupMember) {
	// GroupMembers 比 User 多了一个身份, 少了一些其他信息, 可以 "直观地查看到成员的昵称、真实姓名、邮箱、身份"
	global.DB.Raw(
		"SELECT users.user_id AS user_id, username, real_name, email, status "+
			"FROM identities, users "+
			"WHERE users.user_id = identities.user_id "+
			"AND ? = identities.group_id", groupID).Find(&users).RecordNotFound()
	return users
}

func CreateIdentity(identity *database.Identity) (err error) {
	if err = global.DB.Create(&identity).Error; err != nil {
		return err
	}
	return
}

func UpdateIdentity(identity *database.Identity) (err error) {
	err = global.DB.Save(identity).Error
	return err
}

func DeleteIdentity(identity *database.Identity) (err error) {
	err = global.DB.Delete(&identity).Error
	return err
}
