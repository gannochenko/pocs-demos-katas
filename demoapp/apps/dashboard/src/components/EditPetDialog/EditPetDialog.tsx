import ModalClose from '@mui/joy/ModalClose';
import { Box, Button, FormControl, FormLabel, Input } from '@mui/joy';
import {useEditPetDialog} from "./hooks/useEditPetDialog";
import {Typography} from "../Typography";
import {PetDialogModal, PetDialogSheet} from "./style";
import {TagSelector} from "../TagSelector";
import {CategorySelector} from "../CategorySelector";
import {StatusSelector} from "../StatusSelector";

export type EditPetDialogProps = Partial<{
	open: boolean;
	onClose: () => void;
	petID: string;
}>;

export function EditPetDialog(props: EditPetDialogProps) {
	const {
		title,
		modalProps,
		formProps,
		nameInputProps,
		statusSelectorProps,
		categorySelectorProps,
	} = useEditPetDialog(props);

	return (
		<PetDialogModal
			aria-labelledby="modal-title"
			aria-describedby="modal-desc"
			{...modalProps}
		>
			<PetDialogSheet
				variant="outlined"
				sx={{ borderRadius: 'md', p: 3, boxShadow: 'lg' }}
			>
				<ModalClose variant="plain" sx={{ m: 1 }} />
				<Box
					component="form"
					sx={{
						display: 'flex',
						flexDirection: 'column',
						gap: 2,
						width: '1000px',
						mx: 'auto',
						mt: 4,
					}}
					{...formProps}
				>
					<Typography
						component="h2"
						level="h2"
						sx={{ fontWeight: 'lg', mb: 1 }}
					>
						{title}
					</Typography>
					<FormControl>
						<FormLabel>Name</FormLabel>
						<Input
							type="text"
							{...nameInputProps}
						/>
					</FormControl>
					<FormControl>
						<FormLabel>Status</FormLabel>
						<StatusSelector {...statusSelectorProps} />
					</FormControl>
					<FormControl>
						<FormLabel>Category</FormLabel>
						<CategorySelector {...categorySelectorProps} />
					</FormControl>
					<FormControl>
						<FormLabel>Tags</FormLabel>
						<TagSelector
							value={[
								'd08334b5-42f4-4fef-adb8-b28bedf254d4'
							]}
						/>
					</FormControl>
					<Button type="submit" variant="solid" color="primary">
						Save
					</Button>
				</Box>
			</PetDialogSheet>
		</PetDialogModal>
	);
}
