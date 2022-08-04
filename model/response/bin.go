package response

type DeleteProjQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type DeleteProjA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MovePrototypeFromBinQ struct {
	PrototypeID uint64 `json:"prototype_id"`
}

type MovePrototypeFromBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveUmlFromBinQ struct {
	UmlID uint64 `json:"uml_id"`
}

type MoveUmlFromBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveDocumentFromBinQ struct {
	DocumentID int64 `json:"document_id"`
}

type MoveDocumentFromBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DeletePrototypeQ struct {
	PrototypeID uint64 `json:"prototype_id"`
}

type DeletePrototypeA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DeleteUmlQ struct {
	UmlID uint64 `json:"uml_id"`
}

type DeleteUmlA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DeleteDocumentQ struct {
	DocumentID uint64 `json:"document_id"`
}

type DeleteDocumentA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
