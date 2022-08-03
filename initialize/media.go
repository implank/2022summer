package initialize

import "os"

func InitMedia() {
	_, err := os.Stat("./media")
	if os.IsNotExist(err) {
		_ = os.MkdirAll("./media", 0755)
	}
	_, err = os.Stat("./media/prototypes")
	if os.IsNotExist(err) {
		_ = os.MkdirAll("./media/prototypes", 0755)
	}
	_, err = os.Stat("./media/umls")
	if os.IsNotExist(err) {
		_ = os.MkdirAll("./media/umls", 0755)
	}
	_, err = os.Stat("./media/documents")
	if os.IsNotExist(err) {
		_ = os.MkdirAll("./media/documents", 0755)
	}
	return
}
