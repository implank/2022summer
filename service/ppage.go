package service

import (
	"2022summer/global"
	"2022summer/model/database"
)

func GetPPages(prototypeID uint64) (ppages []database.PPageID) {
	global.DB.Table("p_pages").Select("p_page_id, p_page_name").
		Where("prototype_id = ?", prototypeID).Find(&ppages).RecordNotFound()
	return ppages
}

/*func GetPPages(prototypeID uint64) (ppages []database.PPage) { // 如果需要把 data 一起传回去
	global.DB.Where("prototype_id = ?", prototypeID).Find(&ppages).RecordNotFound()
	return ppages
}*/

func QueryPPageByPPageName(pPageName string, prototypeID uint64) (ppage database.PPage, notFound bool) {
	notFound = global.DB.Where("p_page_name = ? and prototype_id = ?", pPageName, prototypeID).First(&ppage).RecordNotFound()
	return ppage, notFound
}

func QueryPPageByPPageID(pPageID uint64) (ppage database.PPage, notFound bool) {
	notFound = global.DB.Where("p_page_id = ?", pPageID).First(&ppage).RecordNotFound()
	return ppage, notFound
}

func CreatePPage(pPage *database.PPage) (err error) {
	if err = global.DB.Create(&pPage).Error; err != nil {
		return err
	}
	return
}

func UpdatePPage(pPage *database.PPage) (err error) {
	err = global.DB.Save(pPage).Error
	return err
}

func DeletePPage(pPage *database.PPage) (err error) {
	err = global.DB.Delete(&pPage).Error
	return err
}
