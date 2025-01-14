import {PetListProps} from "./type";
import {useListImages} from "../../hooks";
import {Image} from "../../models/image";
import {useState} from "react";

type Upload = {
	id: string;
	file: File;
	createdAt: Date;
};

export function useImageList(props: PetListProps) {
	const imagesResult = useListImages({pageNumber: 1});
	const images: Image[] = imagesResult.data?.images ?? [];
	const [uploads, setUploads] = useState<Upload[]>([]);

	return {
		uploads,
		images,
		empty: !images.length && !uploads.length && !imagesResult.isLoading,
		uploadButtonProps: {
			onChange: async (files: File[]) => {
				if (files.length) {
					setUploads(prevUploads => {
						return [
							...files.map(file => (
								{
									id: Math.floor((Math.random() * 100000)).toString(),
									file,
									createdAt: new Date(),
								}
							)),
							...prevUploads,
						]; //.sort((a, b) => b.createdAt.getTime() - a.createdAt.getTime());
					})
				}
			}
		},
		getImageUploadProps: (upload: Upload) => {
			return {
				upload,
				onSuccess: (id: string) => {
					// setUploads(uploads => {
					// 	return uploads.filter(uploadItem => uploadItem.id !== id);
					// })
				},
			};
		},
	}
}
