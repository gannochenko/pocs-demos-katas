import CircularProgress from '@mui/joy/CircularProgress';
import { Image } from '../Image';
import {Root, Inner, Status, Label, ErrorIcon} from "./style";
import {useImageUpload} from "./useImageUpload";
import { ImageUploadProps } from './type';

export function ImageUpload(props: ImageUploadProps) {
	const { progressProps, failed, imageProps, image } = useImageUpload(props);

	return (
		<Root>
			<Image {...imageProps} />
			<Inner>
				<Status>
					{
						(!failed && !image)
						&&
						<>
                            <CircularProgress {...progressProps} />
                            <Label>
                                Uploading
                            </Label>
						</>
					}
					{
						failed
						&&
						<>
                            <ErrorIcon color="error" />
                            <Label>
                                Failed
                            </Label>
						</>
					}
				</Status>
			</Inner>
		</Root>
	);
}
