package details

import (
	"fmt"
	"model"
)

func CreateDetailProxy(fileData model.FileData) string {
	return fmt.Sprintf(`import { DetailsCreateForm } from "./DetailsCreateForm";
	import { DetailsEditForm } from "./DetailsEditForm";
	import { DetailsViewForm } from "./DetailsViewForm";
	
	export const DetailsFormProxy = ({ type, ...props }) => {
	  const types = {
		create: DetailsCreateForm,
		edit: DetailsEditForm,
		view: DetailsViewForm,
	  };
	
	  const Component = types[type];
	
	  return <Component {...props} />;
	};
	`)
}
