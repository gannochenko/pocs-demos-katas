import {Root} from "./style";
import {ErrorIcon, Inner, Label, Status, ImageContainer, Underlay} from "./style";
import CircularProgress from "@mui/joy/CircularProgress";
import {Upload} from "../../models/image";

type ImageProps = Partial<{
	upload: Upload;
}>;

export function Image({ upload }: ImageProps) {
	const {failed, image, progress} = upload ?? {};

	const progressProps = {
		determinate: true,
		value: progress,
	};
	const imageProps = {
		url: image?.url ?? "",
		isProcessed: !!image?.isProcessed,
	};

	return (
		<Root>
			<ImageContainer {...imageProps} />
			<Inner>
				<Status>
					{
						(!failed && !image)
						&&
                        <Underlay>
                            <CircularProgress {...progressProps} />
                            <Label>
                                Uploading
                            </Label>
                        </Underlay>
					}
					{
						(!failed && !image?.isProcessed)
						&&
                        <Underlay>
                            <CircularProgress />
                            <Label>
                                Processing
                            </Label>
                        </Underlay>
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
