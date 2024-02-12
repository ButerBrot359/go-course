import { lazy } from "react";
import authRoles from "src/app/auth/authRoles";

import { ACTS_JOURNAL_PAGE_URL_PATH } from "../consts/urlPaths";

const ActsJournalPage = lazy(() =>
  import("../ui/ActsJournalPage/ActsJournalPage")
);
const ActsJournalView = lazy(() =>
  import("../ui/Modals/ActsJournalView/ActsJournalView")
);
const ActsJournalForm = lazy(() =>
  import("../ui/Modals/ActsJournalForm/ActsJournalForm")
);

export const ActsJournalRoutesConfig = {
  auth: authRoles.userView,
  routes: [
    {
      path: ACTS_JOURNAL_PAGE_URL_PATH,
      element: <ActsJournalPage />,
      exact: true,
      children: [
        {
          path: "new",
          element: <ActsJournalForm />,
        },
        {
          path: ":id",
          element: <ActsJournalView />,
        },
        {
          path: ":id/edit",
          element: <ActsJournalForm />,
        },
      ],
    },
  ],
};
