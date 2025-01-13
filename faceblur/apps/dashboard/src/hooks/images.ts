import {useMutation, useQuery} from "react-query";
import {useAuth0} from "@auth0/auth0-react";
import {GetUploadURL, ListImages, SubmitImage} from "../proto/image/v1/image";
import {Image} from "../models/image";
import {isError} from "../util/fetch";
import axios from "axios";
import {useState} from "react";

export const LIST_IMAGES_KEY = "list_images";
export const PAGE_SIZE = 9;

type ListImagesRequest = {
	pageNumber: number;
};

type ListImageResponse = {
	images: Image[];
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

type GetUploadURLResponse = {
	url: string;
	objectName: string;
};

export function useGetUploadURL(onSuccess: (response: GetUploadURLResponse) => void, onFailure: () => void) {
	const { getAccessTokenSilently } = useAuth0();
	return useQuery(
		[],
		async (): Promise<GetUploadURLResponse> => {
			const result = await GetUploadURL({}, await getAccessTokenSilently());

			if (isError(result)) {
				// todo: send notification
				return {
					url: "",
					objectName: "",
				};
			}

			return result;
		},
		{
			onError: (error: any) => {
				// todo: show notification here
				onFailure();
			},
			onSuccess,
		},
	);
}

type UploadResponse = {
};

const uploadImage = async (
	url: string,
	file: File,
	onProgress: (progress: number) => void
): Promise<UploadResponse> => {
	const formData = new FormData();
	formData.append("file", file);

	// using axios because it supports progress out of the box
	const response = await axios.put(url, file, {
		headers: {
			"Content-Type": "application/octet-stream",
		},
		onUploadProgress: (progressEvent) => {
			const progress = Math.round(
				(progressEvent.loaded / progressEvent.total!) * 100
			);
			onProgress(progress);
		},
	});

	return response.data;
};

type ImageUploadRequest = {
	url: string;
	file: File;
};

export const useImageUpload = (onProgressChange: (value: number) => void) => {
	return useMutation(
		({url, file}: ImageUploadRequest) => uploadImage(url, file, onProgressChange),
		{
			onSuccess: () => {
				// todo: invalidate cache here
			},
			onError: (error: any) => {
				// todo: show notification here
			},
		}
	);
};

type SubmitImageRequest = {
	objectName: string;
};

export const useSubmitImage = () => {
	const { getAccessTokenSilently } = useAuth0();
	return useMutation(
		async ({objectName}: SubmitImageRequest) => SubmitImage({
			image: {
				objectName,
			},
		}, await getAccessTokenSilently()),
		{
			onSuccess: () => {
				// todo: invalidate cache here
			},
			onError: (error: any) => {
				// todo: show notification here
			},
		}
	);
};
