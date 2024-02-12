package details

import (
	"fmt"
	"model"
)

func CreateDetailProxyStyle(fileData model.FileData) string {
	return fmt.Sprintf(`.modal-content {
		display: flex;
		flex-direction: column;
		row-gap: 20px;
	  }
	  
	  .modal-content__field {
		border-bottom: 1px solid grey;
		padding: 15px 10px;
	  }
	  
	  .strong-text {
		font-weight: 600;
	  }
	  `)
}
