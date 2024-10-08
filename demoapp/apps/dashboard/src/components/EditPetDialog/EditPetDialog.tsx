import Modal from '@mui/joy/Modal';
import Sheet from '@mui/joy/Sheet';
import ModalClose from '@mui/joy/ModalClose';
import {useEditPetDialog} from "./hooks/useEditPetDialog";
import {Typography} from "../Typography";

export type EditPetDialogProps = Partial<{
	open: boolean;
	onClose: () => {};
	data: any;
	onDataSave: () => {};
}>;

export function EditPetDialog(props: EditPetDialogProps) {
	const { modalProps } = useEditPetDialog(props);

	return (
		<Modal
			aria-labelledby="modal-title"
			aria-describedby="modal-desc"
			sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}
			{...modalProps}
		>
			<Sheet
				variant="outlined"
				sx={{ maxWidth: 500, borderRadius: 'md', p: 3, boxShadow: 'lg' }}
			>
				<ModalClose variant="plain" sx={{ m: 1 }} />
				<Typography
					component="h2"
					id="modal-title"
					level="h4"
					textColor="inherit"
					sx={{ fontWeight: 'lg', mb: 1 }}
				>
					This is the modal title
				</Typography>
				<Typography id="modal-desc" textColor="text.tertiary">
					Make sure to use <code>aria-labelledby</code> on the modal dialog with an
					optional <code>aria-describedby</code> attribute.
				</Typography>
			</Sheet>
		</Modal>
	);
}
