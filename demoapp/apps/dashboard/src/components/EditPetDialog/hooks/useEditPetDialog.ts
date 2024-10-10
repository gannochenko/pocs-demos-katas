import {FieldValues, useForm} from 'react-hook-form';
import {EditPetDialogProps} from "../EditPetDialog";
import {Pet} from "../../../models/pet";
import {useGetPet} from "../../../hooks/pets/useGetPet";
import {useState, useEffect, ChangeEvent, useMemo, useCallback} from "react";
import {useUpdatePet} from "../../../hooks/pets/useUpdatePet";

type FormState = {
	name: string;
	status: string;
	categoryID: string;
	tagIDs: string[];
};

const emptyValue: FormState = {
	name: "",
	status: "",
	categoryID: "",
	tagIDs: [],
};

export const useEditPetDialog = ({open, onClose, petID}: EditPetDialogProps) => {
	const petResult = useGetPet({
		id: petID ?? "",
	});
	const pet: Pet | undefined = petResult.data?.pet;

	const updatePet = useUpdatePet();

	const [formState, setFormState] = useState<FormState>(emptyValue);
	useEffect(() => {
		if (pet) {
			setFormState({
				name: pet.name,
				status: pet.status,
				categoryID: "",
				tagIDs: [],
			});
			setValue("name", pet.name);
		} else {
			resetForm();
		}
	}, [pet]);

	const { register, setValue, handleSubmit, formState: { errors } } = useForm();
	const resetForm = useCallback(() => {
		setFormState(emptyValue);
		setValue("name", "");
		setValue("status", "");
	}, [setValue, setFormState]);

	const onSubmit = (data: FieldValues) => {
		console.log(data);
		updatePet.mutate({
			pet: {
				id: petID ?? "",
				name: data.name,
				status: data.status,
			},
		}, {
			onSuccess: () => {
				resetForm();
				onClose?.();
			},
		})
	};

	return {
		modalProps: {
			open: !!open,
			onClose: () => {
				resetForm();
				onClose?.();
			},
		},
		title: petID ? `Pet ${petID}` : "New pet",
		formProps: {
			onSubmit: handleSubmit(onSubmit),
		},
		nameInputProps: {
			value: formState.name,
			...register("name"),
			onChange: (e: ChangeEvent<HTMLInputElement>) => {
				const newName = e.target.value;
				setFormState((prevState) => {
					return {
						...prevState,
						name: newName,
					};
				});
				setValue("name", newName);
			},
		},
		statusSelectorProps: {
			value: formState.status,
			...register("status"),
			onChange: (e: ChangeEvent<HTMLInputElement>) => {
				const newName = e.target.value;
				setFormState((prevState) => {
					return {
						...prevState,
						name: newName,
					};
				});
				setValue("name", newName);
			},
		},
	};
};
