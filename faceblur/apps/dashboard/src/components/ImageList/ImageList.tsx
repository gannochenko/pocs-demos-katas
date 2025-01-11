import ReactDOM from 'react-dom';
import Grid from '@mui/joy/Grid';
import {Root, ImageItem} from "./style";
import {Image} from "../../models/image";
import {Link} from "../Link";
import {useListImages} from "../../hooks";
import Button from "@mui/joy/Button";
import React from "react";
import {PortalToID} from "../PortalToID/PortalToID";

type PetListProps = Partial<{
	onRowClick: (petID: string) => void;
}>;

export function ImageList({onRowClick}: PetListProps) {
	const imagesResult = useListImages({pageNumber: 1});
	const images: Image[] = imagesResult.data?.images ?? [];

	return (
		<>
			<PortalToID id="page-header-portal">
				<Button>
					Upload new photo
				</Button>
			</PortalToID>
			<Root>
				<Grid container spacing={2} sx={{ flexGrow: 1 }}>
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
