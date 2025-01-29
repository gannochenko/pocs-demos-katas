import {useEffect, useState} from "react";
import {useAuth0} from "@auth0/auth0-react";
import {PetListProps} from "./type";
import {useListImages} from "@/hooks";
import {ImageModel, Upload} from "@/models/image";
import {isErrorResponse, uploadFile} from "@/util/fetch";
import {useNotification} from "@/hooks/notification";
import {useWebsocketContext} from "../WebsocketProvider";
import {ServerMessage, ServerMessageType} from "@/proto/websocket/v1/websocket";
import {GetUploadURL, SubmitImage} from "@/proto/image/v1/image";

const useUploader = () => {
	const [uploads, setUploads] = useState<Upload[]>([]);
	const { getAccessTokenSilently } = useAuth0();
	const {showError} = useNotification();
	const {addEventListener, removeEventListener} = useWebsocketContext();

	useEffect(() => {
		const handler = (payload: ServerMessage) => {
			console.log("RECEIVED");
			console.log(payload);
		};

		addEventListener(ServerMessageType.SERVER_MESSAGE_TYPE_IMAGE_LIST, handler);

		return () => removeEventListener(ServerMessageType.SERVER_MESSAGE_TYPE_IMAGE_LIST, handler);
	}, [addEventListener, removeEventListener]);

	const updateUpload = (id: string, updateCb: (upload: Upload) => Upload) => {
		setUploads(prevState => {
			return prevState.map(upload => {
				if (upload.id === id) {
					return updateCb(upload);
				}

				return upload;
			});
		});
	};

	const showUploadError = (reason: string) => {
		showError("Error uploading file", reason);
	};

	return {
		uploads,
		submit: async (newUploads: Upload[]) => {
			newUploads = newUploads.sort((a, b) => (b.uploadedAt?.getDate() ?? 0) - (a.uploadedAt?.getDate() ?? 0));

			setUploads(prevUploads => {
				return [
					...newUploads,
					...prevUploads,
				];
			});

			for (const upload of newUploads) {
				const getUploadULRResponse = await GetUploadURL({}, await getAccessTokenSilently())
				if (isErrorResponse(getUploadULRResponse)) {
					showUploadError(getUploadULRResponse.error);
					updateUpload(upload.id, (upload) => ({
						...upload,
						failed: true,
					}))
				} else {
					await uploadFile(getUploadULRResponse.url, upload.file!, (newProgress) => {
						updateUpload(upload.id, (upload) => ({
							...upload,
							progress: newProgress,
						}))
					});
					const submitImageResponse = await SubmitImage({
						image: {
							objectName: getUploadULRResponse.objectName,
							uploadedAt: upload.uploadedAt!,
						},
					}, await getAccessTokenSilently());
					if (isErrorResponse(submitImageResponse)) {
						showUploadError(submitImageResponse.error);
						updateUpload(upload.id, (upload) => ({
							...upload,
							failed: true,
						}))
					} else {
						updateUpload(upload.id, (upload) => ({
							...upload,
							image: submitImageResponse.image,
						}))
					}
				}
			}
		},
	};
};

export function useImageList(props: PetListProps) {
	const imagesResult = useListImages({pageNumber: 1});
	const images: ImageModel[] = imagesResult.data?.images ?? [];

	const {uploads, submit} = useUploader();

	// todo: reconcile images and uploads by id. if id is present in both, take the image, not the upload

	return {
		uploads,
		images,
		empty: !images.length && !uploads.length && !imagesResult.isLoading,
		uploadButtonProps: {
			onChange: async (files: File[]) => {
				if (files.length) {
					submit(files.map(file => (
						{
							id: Math.floor((Math.random() * 100000)).toString(),
							file,
							uploadedAt: new Date(),
							progress: 0,
						}
					)))
				}
			}
		},
		getImageUploadProps: (upload: Upload) => {
			return {
				upload,
			};
		},
		getImageUploadPropsByImage: (image: ImageModel) => {
			return {
				upload: {
					id: image.id,
					file: undefined,
					image,
					uploadedAt: image.updatedAt,
					failed: image.isFailed,
					progress: 100,
				},
			};
		},
	}
}
