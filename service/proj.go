package service

import (
	"2022summer/global"
	"2022summer/model/database"
	"github.com/jinzhu/gorm"
)

func QueryProjByProjName(projName string) (proj database.Proj, notFound bool) {
	notFound = global.DB.Where("proj_name = ?", projName).First(&proj).RecordNotFound()
	return proj, notFound
}

func GetProjsByProjNameBur(projName string) (projs []database.Proj) { // 模糊搜索
	global.DB.Where("proj_name like '%" + projName + "%' and status = 1").Find(&projs).RecordNotFound()
	return projs
}

func QueryProjByProjID(projID uint64) (proj database.Proj, notFound bool) {
	notFound = global.DB.Where("proj_id = ?", projID).First(&proj).RecordNotFound()
	return proj, notFound
}

func CreateProj(proj *database.Proj) (err error) {
	if err = global.DB.Create(&proj).Error; err != nil {
		return err
	}
	return
}

func UpdateProj(proj *database.Proj) (err error) {
	proj.EditTimes += 1
	err = global.DB.Save(proj).Error
	return err
}

func UpdateProjStatus(proj *database.Proj) (err error) {
	// 移入或移出回收站: 先改变 PPage, Uml, Document 的 Status, 再保存 Proj
	// 【注意】传入的 Proj 的 Status 应该是已被修改的
	global.DB.Model(database.PPage{}).Where("proj_id = ?", proj.ProjID).Updates(database.PPage{Status: proj.Status})
	global.DB.Model(database.Uml{}).Where("proj_id = ?", proj.ProjID).Updates(database.Uml{Status: proj.Status})
	global.DB.Model(database.Document{}).Where("proj_id = ?", proj.ProjID).Updates(database.Document{Status: proj.Status})
	proj.EditTimes += 1
	err = global.DB.Save(proj).Error
	return err
}

func DeleteProj(proj *database.Proj) (err error) {
	// 删除回收站中的 Proj: 先删除 PPage, Uml, Document, 再删除 Proj
	global.DB.Where("proj_id = ?", proj.ProjID).Delete(database.PPage{})
	global.DB.Where("proj_id = ?", proj.ProjID).Delete(database.Uml{})
	global.DB.Where("proj_id = ?", proj.ProjID).Delete(database.Document{})
	err = global.DB.Delete(&proj).Error
	return err
}

func GetUserProjs(userID uint64, status int, op int) (projs []database.Proj) {
	// 查找该用户在所有团队中的项目
	// op = 1 时查找"我创建的", op = 2 时查找"我参与的", op = 3 时查找"全部项目"
	if op == 1 {
		global.DB.Where("user_id = ? and status = ?", userID, status).Find(&projs).RecordNotFound()
	} else if op == 2 {
		global.DB.Raw("SELECT * FROM projs, identities "+
			"WHERE projs.group_id = identities.group_id "+
			"AND identities.user_id = ? AND projs.user_id != ? "+
			"AND projs.status = ?;", userID, userID, status).Find(&projs).RecordNotFound()
	} else {
		global.DB.Raw("SELECT * FROM projs, identities "+
			"WHERE projs.group_id = identities.group_id "+
			"AND identities.user_id = ? AND projs.status = ?;", userID, status).Find(&projs).RecordNotFound()
	}
	return projs
}

func GetUserProjsInGroup(userID uint64, groupID uint64, status int, op int, orderBy int, isDesc bool) (projs []database.Proj) {
	// 查找该用户在某个团队中的项目
	// op = 1 时查找"我创建的", op = 2 时查找"我参与的", op = 3 时查找"全部项目"
	var query *gorm.DB
	if op == 1 {
		query = global.DB.Where("user_id = ? and group_id = ? and status = ?",
			userID, groupID, status)
	} else if op == 2 {
		query = global.DB.Raw("SELECT * FROM projs, identities "+
			"WHERE identities.user_id = ? AND projs.user_id != ? "+
			"AND projs.group_id = ? AND identities.group_id = ? "+
			"AND projs.status = ?", userID, userID, groupID, groupID, status)
	} else {
		query = global.DB.Raw("SELECT * FROM projs, identities "+
			"WHERE identities.user_id = ? "+
			"AND projs.group_id = ? AND identities.group_id = ? "+
			"AND projs.status = ?", userID, groupID, groupID, status)
	}
	query = query.Order("top DESC")
	var str string
	if orderBy == 1 {
		str = "created_at"
	} else if orderBy == 2 {
		str = "updated_at "
	} else {
		str = "edit_times"
	}
	if isDesc {
		str += " DESC"
	}
	query = query.Order(str)
	query.Find(&projs).RecordNotFound()
	return projs
}
