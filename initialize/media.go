package initialize

import (
	"2022summer/global"
	"os"
)

func InitMedia() {
	_, err := os.Stat("./media")
	if os.IsNotExist(err) {
		_ = os.MkdirAll("./media", 0755)
	}
	_, err = os.Stat("./media/ppages")
	if os.IsNotExist(err) {
		_ = os.MkdirAll("./media/ppages", 0755)
	}
	_, err = os.Stat("./media/umls")
	if os.IsNotExist(err) {
		_ = os.MkdirAll("./media/umls", 0755)
	}
	_, err = os.Stat("./media/documents")
	if os.IsNotExist(err) {
		_ = os.MkdirAll("./media/documents", 0755)
	}
	_, err = os.Stat(global.VP.Get("avatar_dir").(string))
	if os.IsNotExist(err) {
		_ = os.MkdirAll(global.VP.Get("avatar_dir").(string), 0755)
	}
	_, err = os.Stat(global.VP.Get("image_dir").(string))
	if os.IsNotExist(err) {
		_ = os.MkdirAll(global.VP.Get("image_dir").(string), 0755)
	}
	return
}
