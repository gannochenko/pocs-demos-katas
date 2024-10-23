import CheckCircle from '@mui/icons-material/CheckCircle';
import Chip from "@mui/joy/Chip";
import {Tag} from "../../models/tag";
import {TagViewerRoot} from "./style";

type TagViewerProps = Partial<{
	tags: Tag[];
}>;

export function TagViewer({tags}: TagViewerProps) {
	return (
		<TagViewerRoot>
			{tags?.map(tag => (
				<Chip
					key={tag.id}
					variant="soft"
					size="lg"
					color="success"
				>
					{tag.name}
				</Chip>
			))}
		</TagViewerRoot>
	);
}
