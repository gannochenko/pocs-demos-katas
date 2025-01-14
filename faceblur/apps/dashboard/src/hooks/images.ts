import {useQuery} from "react-query";
import {useAuth0} from "@auth0/auth0-react";
import {ListImages} from "../proto/image/v1/image";
import {ImageModel} from "../models/image";
import {isError} from "../util/fetch";

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
	return useQuery(
		[LIST_IMAGES_KEY, ...[request]],
		async (): Promise<ListImageResponse> => {
			const result = await ListImages({
				pageNavigation: {
					pageSize: PAGE_SIZE,
					pageNumber: request.pageNumber,
				},
			}, await getAccessTokenSilently());

			if (isError(result)) {
				// todo: send notification
				return {
					images: [],
				};
			}

			return {
				images: result.images,
			};
		},
		{
			keepPreviousData: true,
			refetchOnWindowFocus: false,
			onError: (error: any) => {
				// todo: show notification here
			},
		},
	);
}
