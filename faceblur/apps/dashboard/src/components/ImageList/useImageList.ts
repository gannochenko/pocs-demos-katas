import {PetListProps} from "./type";
import {useListImages} from "../../hooks";
import {Image} from "../../models/image";

export function useImageList(props: PetListProps) {
	const imagesResult = useListImages({pageNumber: 1});
	const images: Image[] = imagesResult.data?.images ?? [];

	return {
		images,
		uploadButtonProps: {
			onChange: (files: File[]) => {
				console.log(files);
			}
		},
	}
}
