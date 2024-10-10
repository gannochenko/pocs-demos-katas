import {FieldValues, useForm} from 'react-hook-form';
import {EditPetDialogProps} from "../EditPetDialog";
import {Pet} from "../../../models/pet";
import {useGetPet} from "../../../hooks/pets/useGetPet";
import {ChangeEvent, useCallback, useEffect, useState} from "react";
import {useUpdatePet} from "../../../hooks/pets/useUpdatePet";
import {Status} from "../../StatusSelector";

type FormState = {
	name: string;
	status: Status;
	category: string;
	tagIDs: string[];
};

const emptyValue: FormState = {
	name: "",
	status: Status.Pending,
	category: "",
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
			const newValue = {
				name: pet.name,
				status: pet.status as Status,
				category: pet.category?.id ?? "",
				tagIDs: [],
			}

			setFormState(newValue);
			setValue("name", newValue.name);
			setValue("status", newValue.status);
			setValue("category", newValue.category);
		} else {
			resetForm();
		}
	}, [pet]);

	const { register, setValue, handleSubmit, control, formState: { errors } } = useForm();
	const resetForm = useCallback(() => {
		setFormState(emptyValue);
		setValue("name", "");
		setValue("status", "");
		setValue("category", "");
	}, [setValue, setFormState]);

	const onSubmit = (data: FieldValues) => {
		console.log(data);
		updatePet.mutate({
			pet: {
				id: petID ?? "",
				name: data.name,
				status: data.status,
				category: {
					id: data.category as string,
					name: "",
				},
				tags: [],
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
			onChange: (newValue: Status) => {
				setFormState((prevState) => {
					return {
						...prevState,
						status: newValue,
					};
				});
				setValue("status", newValue);
			},
		},
		categorySelectorProps: {
			value: formState.category,
			...register("category"),
			onChange: (newValue: string) => {
				setFormState((prevState) => {
					return {
						...prevState,
						category: newValue,
					};
				});
				setValue("category", newValue);
			},
		},
	};
};
