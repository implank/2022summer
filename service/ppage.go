package service

import (
	"2022summer/global"
	"2022summer/model/database"
	"github.com/jinzhu/gorm"
)

func GetProjPPages(projID uint64, status int) (ppages []database.PPageID) {
	global.DB.Table("p_pages").Select("p_page_id, p_page_name").Order("p_page_id DESC").
		Where("proj_id = ? and status = ?", projID, status).Find(&ppages).RecordNotFound()
	return ppages
}

/*func GetProjPPages(projID uint64, status int) (ppages []database.PPage) { // 如果需要把 data 一起传回去
	global.DB.Where("proj_id = ? and status = ?", projID, status).Find(&ppages).RecordNotFound()
	return ppages
}*/

func QueryPPageByPPageName(pPageName string, projID uint64) (ppage database.PPage, notFound bool) {
	// 同一项目的设计原型中不能有同名页面
	notFound = global.DB.Where("p_page_name = ? and proj_id = ?", pPageName, projID).First(&ppage).RecordNotFound()
	return ppage, notFound
}

func QueryPPageByPPageID(pPageID uint64) (ppage database.PPage, notFound bool) {
	notFound = global.DB.Where("p_page_id = ?", pPageID).First(&ppage).RecordNotFound()
	return ppage, notFound
}

func CreatePPage(pPage *database.PPage) (err error) {
	global.DB.Model(database.Proj{}).Where("proj_id = ?", pPage.ProjID).
		Update("edit_times", gorm.Expr("edit_times + ?", 1))
	if err = global.DB.Create(&pPage).Error; err != nil {
		return err
	}
	return
}

func UpdatePPage(pPage *database.PPage) (err error) {
	global.DB.Model(database.Proj{}).Where("proj_id = ?", pPage.ProjID).
		Update("edit_times", gorm.Expr("edit_times + ?", 1))
	err = global.DB.Save(pPage).Error
	return err
}

func DeletePPage(pPage *database.PPage) (err error) {
	err = global.DB.Delete(&pPage).Error
	return err
}

func CreateSharedPPage(sharedPPage *database.SharedPPage) (err error) {
	err = global.DB.Create(&sharedPPage).Error
	return err
}

func DeleteSharedPPage(sharedPPage *database.SharedPPage) (err error) {
	err = global.DB.Delete(&sharedPPage).Error
	return err
}

func QuerySharedPPageByToken(token string) (sharedPPage database.SharedPPage, notFound bool) {
	notFound = global.DB.Where("token = ?", token).First(&sharedPPage).RecordNotFound()
	return sharedPPage, notFound
}

func QuerySharedPPageByPPageID(pPageID uint64) (sharedPPage database.SharedPPage, notFound bool) {
	notFound = global.DB.Where("p_page_id = ?", pPageID).First(&sharedPPage).RecordNotFound()
	return sharedPPage, notFound
}
