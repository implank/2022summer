package service

import (
	"2022summer/global"
	"2022summer/model/database"
	"github.com/jinzhu/gorm"
)

func GetProjUmls(projID uint64, status int) (umls []database.UmlID) {
	global.DB.Table("umls").Select("uml_id, uml_name").Order("uml_id DESC").
		Where("proj_id = ? and status = ?", projID, status).Find(&umls).RecordNotFound()
	return umls
}

func QueryUmlByUmlName(umlName string, projID uint64) (uml database.Uml, notFound bool) {
	// 同一项目中不能有同名绘制图
	notFound = global.DB.Where("uml_name = ? and proj_id = ?", umlName, projID).First(&uml).RecordNotFound()
	return uml, notFound
}

func QueryUmlByUmlID(umlID uint64) (uml database.Uml, notFound bool) {
	notFound = global.DB.Where("uml_id = ?", umlID).First(&uml).RecordNotFound()
	return uml, notFound
}

func CreateUml(uml *database.Uml) (err error) {
	global.DB.Model(database.Proj{}).Where("proj_id = ?", uml.ProjID).
		Update("edit_times", gorm.Expr("edit_times + ?", 1))
	if err = global.DB.Create(&uml).Error; err != nil {
		return err
	}
	return
}

func UpdateUml(uml *database.Uml) (err error) {
	global.DB.Model(database.Proj{}).Where("proj_id = ?", uml.ProjID).
		Update("edit_times", gorm.Expr("edit_times + ?", 1))
	err = global.DB.Save(uml).Error
	return err
}

func DeleteUml(uml *database.Uml) (err error) {
	err = global.DB.Delete(&uml).Error
	return err
}
