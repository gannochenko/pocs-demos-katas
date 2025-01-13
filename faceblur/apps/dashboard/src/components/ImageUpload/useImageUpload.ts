import {useState} from "react";
import {useGetUploadURL, useImageUpload as useImageUploadMutation, useSubmitImage as useSubmitImageMutation} from "../../hooks";
import {ImageUploadProps} from "./type";

export const useImageUpload = ({upload, onSuccess}: ImageUploadProps) => {
	const {file} = upload!;
	const [progress, setProgress] = useState<number>(0);
	const uploadImageMutation = useImageUploadMutation(setProgress);
	const submitImageMutation = useSubmitImageMutation();
	const [failed, setFailed] = useState(false);

	useGetUploadURL((response) => {
		uploadImageMutation.mutate({
			url: response.url,
			file,
		}, {
			onSuccess: () => {
				submitImageMutation.mutate({
					objectName: response.objectName,
				}, {
					onError: () => setFailed(true),
					onSuccess,
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
	};
};
