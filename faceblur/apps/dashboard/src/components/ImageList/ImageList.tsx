import Table from '@mui/joy/Table';
import Grid from '@mui/joy/Grid';
import {PetListRoot} from "./style";
import {Image} from "../../models/image";
import {Link} from "../Link";
import {useListImages} from "../../hooks";

type PetListProps = Partial<{
	onRowClick: (petID: string) => void;
}>;

export function ImageList({onRowClick}: PetListProps) {
	const imagesResult = useListImages({pageNumber: 1});
	const images: Image[] = imagesResult.data?.images ?? [];

	return (
		<div>
			<Grid container spacing={2} sx={{ flexGrow: 1 }}>
				<Grid xs={4}>
					1
				</Grid>
				<Grid xs={4}>
					2
				</Grid>
				<Grid xs={4}>
					3
				</Grid>
				<Grid xs={4}>
					1
				</Grid>
				<Grid xs={4}>
					2
				</Grid>
				<Grid xs={4}>
					3
				</Grid>
			</Grid>
		</div>
	);
}
