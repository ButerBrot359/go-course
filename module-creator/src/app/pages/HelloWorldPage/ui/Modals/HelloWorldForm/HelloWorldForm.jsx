
	import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { useOutletContext, useParams } from "react-router-dom";

import { Button } from "@mui/material";
import { yupResolver } from "@hookform/resolvers/yup";

import { FormHeader } from "src/app/entities/FormViewWrappers";

import { useNavigateToRoute } from "src/app/shared/lib/hooks/useNavigateToRoute";
import { useFetchAndReturnData } from "src/app/shared/lib/hooks/useFetchAndReturnData";
import { HookFormDateInput } from "src/app/shared/ui/HookFormInputs/HookFormDateInput/HookFormDateInput";
import { HookFormTextField } from "src/app/shared/ui/HookFormInputs/HookFormTextField/HookFormTextField";
import { HookFormAutocomplete } from "src/app/shared/ui/HookFormInputs/HookFormAutocomplete/HookFormAutocomplete";

import { helloWorldFormSchema } from "../../../models/helloWorldFormSchema";

import stl from "./HelloWorldForm.module.scss";

const HelloWorldForm = () => {
  const { id: elementId } = useParams();
  const { openSidebar, closeSidebar, backUrl } = useOutletContext();
  const [navigateBack, navigateToElement] = useNavigateToRoute(backUrl);

  const [isSelectedDataLoading, fetchSelectedElement] = useFetchAndReturnData(
    {}
  );

  const { control, formState, handleSubmit, watch, reset } = useForm({
    mode: "onChange",
    resolver: yupResolver(helloWorldFormSchema),
  });

  const setResetData = async () => {
    const resetData = await fetchSelectedElement(() => {}, {
      id: elementId,
    });

    const updatedResetData = {
      ...resetData,
    };

    reset(updatedResetData);
  };

  const onSubmit = async (formData) => {
    let payload = {
      ...formData,
    };

    setIsRequestSent(true);

    try {
      if (elementId) {
        payload.id = elementId;

        // updateService
        await new Promise((res) => res())({ payload });
      } else {
        //create service
        await new Promise((res) => res())({ payload });
      }

      navigateBack({ refresh: true })();
    } catch (error) {
      console.log(error);
    } finally {
      setIsRequestSent(false);
      reset({});
    }
  };

  const { errors, isValid } = formState;

  useEffect(() => {
    openSidebar();

    return () => {
      closeSidebar();
    };
  }, []);

  useEffect(() => {
    if (!elementId) return;

    setResetData();
  }, [elementId]);

  return (
    <div className={stl["sidebar-form"]}>
      <FormHeader
        backAction={() => {
          if (elementId) {
            navigateToElement({ refresh: false })(elementId);
          } else {
            navigateBack({ refresh: false })();
          }
        }}
        closeAction={navigateBack({ refresh: false })}
      />
      <main className={stl["main"]}>
        <HookFormAutocomplete
          control={control}
          inputName="test1"
          label="test1"
          errors={errors}
          getOptionLabel={(option) => option.title || ""}
          options={[]}
        />
        <HookFormTextField
          type="text"
          control={control}
          inputName="test2"
          label="test2"
          errors={errors}
        />
        <HookFormDateInput
          control={control}
          label="test3"
          inputName="test3"
          errors={errors}
        />
      </main>
      <footer className={stl["footer"]}>
        <Button
          variant="contained"
          color="success"
          disabled={Object.keys(errors).length > 0 || !isValid}
          onClick={handleSubmit(onSubmit)}
        >
          {elementId ? "Обновить" : "Сохранить"}
        </Button>
      </footer>
    </div>
  );
};

export default HelloWorldForm;

	