import {useQuery} from "react-query";
import {useAuth0} from "@auth0/auth0-react";
import {useNotification} from "@/hooks/notification";
import {ListImages} from "@/proto/image/v1/image";
import {ImageModel} from "@/models/image";
import {isErrorResponse} from "@/util/fetch";

export const LIST_IMAGES_KEY = "list_images";
export const PAGE_SIZE = 9;

type ListImagesRequest = {
	pageNumber: number;
};

type ListImageResponse = {
	images: ImageModel[];
};

export function useListImages(request: ListImagesRequest) {
	const { getAccessTokenSilently } = useAuth0();
	const {showError} = useNotification();
	return useQuery(
		[LIST_IMAGES_KEY, ...[request]],
		async (): Promise<ListImageResponse> => {
			let token = "";
			try {
				token = await getAccessTokenSilently()
			} catch (e) {
				showError("Unauthorized");
				return {
					images: [],
				};
			}

			const response = await ListImages({
				pageNavigation: {
					pageSize: PAGE_SIZE,
					pageNumber: request.pageNumber,
				},
			}, token);

			if (isErrorResponse(response)) {
				showError("Error fetching images", response.error);
				return {
					images: [],
				};
			}

			return {
				images: response.images,
			};
		},
		{
			keepPreviousData: true,
			refetchOnWindowFocus: false,
			onError: (error: any) => {
				showError("Error fetching images", error);
			},
		},
	);
}
