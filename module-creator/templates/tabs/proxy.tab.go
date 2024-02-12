package tabs

import (
	"helpers"
	"model"
)

func CreateProxyTab(fileData model.FileData) string {
	template := `import { FirstTab } from "./FirstTab";
	import { SecondTab } from "./SecondTab";
	
	export const [PascalCaseModuleName]TabsProxy = ({ type, ...props }) => {
	  const types = {
		firstTab: FirstTab,
		secondTab: SecondTab,
	  };
	
	  const Component = types[type];
	
	  return <Component {...props} />;
	};
	`

	return helpers.ModifyTemplateString(template, fileData)
}
