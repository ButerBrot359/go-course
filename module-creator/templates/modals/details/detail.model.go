package details

import "model"

func CreateDetailModel(fileData model.FileData) string {
	return `import * as yup from "yup";

	export const DetailsFormSchema = yup.object().shape({});
	`
}
