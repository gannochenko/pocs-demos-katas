import {useEffect, useState} from "react";
import {useGetUploadURL, useImageUpload as useImageUploadMutation, useSubmitImage as useSubmitImageMutation} from "../../hooks";
import {ImageUploadProps} from "./type";
import {isError} from "../../util/fetch";
import {Image} from "../../models/image";

export const useImageUpload = ({upload, onSuccess}: ImageUploadProps) => {
	const {file} = upload!;
	const [progress, setProgress] = useState<number>(0);
	const uploadImageMutation = useImageUploadMutation(setProgress);
	const submitImageMutation = useSubmitImageMutation();
	const [failed, setFailed] = useState(false);
	const [image, setImage] = useState<Image>();
	const [queryEnabled, setQueryEnabled] = useState(true);

	useEffect(() => {
		return () => {
			console.log('UNMOUNT');
		};
	}, []);

	useGetUploadURL(queryEnabled, (response) => {
		uploadImageMutation.mutate({
			url: response.url,
			file,
		}, {
			onSuccess: () => {
				setQueryEnabled(false);
				submitImageMutation.mutate({
					objectName: response.objectName,
				}, {
					onError: () => setFailed(true),
					onSuccess: (data) => {
						if (!isError(data)) {
							onSuccess?.(upload?.id ?? "");
							setImage(data.image);
						}
					},
				});
			},
			onError: () => setFailed(true),
		});
	}, () => {
		setFailed(true);
	});

	return {
		progressProps: {
			determinate: true,
			value: progress,
		},
		failed,
		imageProps: {
			url: image?.url ?? undefined,
		},
		image,
	};
};
