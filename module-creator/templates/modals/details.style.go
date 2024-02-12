package modals

import (
	"fmt"
	"model"
)

func CreateDetailsStyle(fileData model.FileData) string {
	return fmt.Sprintf(`.details-form-modal-content {
		padding-top: 40px;
		font-size: 16px;
	  }
	  
	  .details-form-modal-content--create,
	  .details-form-modal-content--edit {
		display: flex;
		flex-direction: column;
		row-gap: 20px;
	  }
	  `)
}
