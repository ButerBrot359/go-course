package details

import (
	"helpers"
	"model"
)

func CreateDetailView(fileData model.FileData) string {
	template := `import { isObjectEmpty } from "src/app/helpers/check";
	import s from "./index.module.scss";
	
	export const DetailsViewForm = ({
	  elementId,
	  data,
	  className,
	  closeModal = () => {},
	}) => {
	  if (!data) return <></>;
	
	  const isValidData = !isObjectEmpty(data);
	
	  const { mock = "mock data" } = data;
	
	  return (
		<>
		  {isValidData ? (
			<div className={[back-quote]${className} ${s["modal-content"]}[back-quote]}> // Change quotes type
			  <div className={s["modal-content__field"]}>
				Mock field:
				<strong className={s["strong-text"]}>{mock}</strong>
			  </div>
			</div>
		  ) : (
			<></>
		  )}
		</>
	  );
	};
	`

	return helpers.ModifyTemplateString(template, fileData)
}
