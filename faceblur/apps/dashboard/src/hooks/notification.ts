import {useSnackbar} from "notistack";

export const useNotification = () => {
	const { enqueueSnackbar } = useSnackbar();

	return {
		showError: (message: string, reason?: string) => {
			let notificationMessage = message;
			if (reason) {
				notificationMessage = `${notificationMessage}: ${reason}`;
			}

			enqueueSnackbar(
				notificationMessage,
				{
					variant: "error",
				}
			);
		},
	};
};
