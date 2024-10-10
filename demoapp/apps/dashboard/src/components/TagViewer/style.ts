import {Box, styled} from "@mui/joy";
import {spacing} from "../../util/mixins";
import Chip from "@mui/joy/Chip";

export const TagViewerRoot = styled(Box)`
	display: flex;
	flex-wrap: wrap;
	gap: ${spacing(1)};
`;
