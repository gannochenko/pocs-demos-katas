import CircularProgress from '@mui/joy/CircularProgress';
import { Image } from '../Image';
import {Root, Inner, Status, Label, ErrorIcon} from "./style";
import {useImageUpload} from "./useImageUpload";
import { ImageUploadProps } from './type';

export function ImageUpload(props: ImageUploadProps) {
	const { progressProps, failed } = useImageUpload(props);

	return (
		<Root>
			<Image />
			<Inner>
				<Status>
					{
						!failed
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
