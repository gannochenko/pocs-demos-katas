import { useForm } from 'react-hook-form';
import {EditPetDialogProps} from "../EditPetDialog";

export const useEditPetDialog = ({open, onClose}: EditPetDialogProps) => {
	const { register, handleSubmit, formState: { errors } } = useForm();

	return {
		modalProps: {
			open,
			onClose,
		},
	};
};
