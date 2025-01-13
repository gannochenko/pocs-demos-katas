import Grid from '@mui/joy/Grid';
import {Root, ImageItem} from "./style";
import React from "react";
import {PortalToID} from "../PortalToID/PortalToID";
import {UploadButton} from "../UploadButton";
import {ImageUpload} from "../ImageUpload";
import {PetListProps} from "./type";
import {useImageList} from "./useImageList";

export function ImageList(props: PetListProps) {
	const { uploads, images, uploadButtonProps, getImageUploadProps } = useImageList(props);

	return (
		<>
			<PortalToID id="page-header-portal">
				<UploadButton {...uploadButtonProps} />
			</PortalToID>
			<Root>
				<Grid container spacing={2} sx={{ flexGrow: 1 }}>
					{
						uploads.map(upload => (
							<Grid xs={4} id={upload.id}>
								<ImageUpload {...getImageUploadProps(upload)} />
							</Grid>
						))
					}
					{
						images.map(image => (
							<Grid xs={4} id={image.id}>
								<ImageItem image={image.url} />
							</Grid>
						))
					}
				</Grid>
			</Root>
		</>
	);
}
