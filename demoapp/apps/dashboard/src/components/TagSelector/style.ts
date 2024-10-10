import {Box, styled} from "@mui/joy";
import {spacing} from "../../util/mixins";
import Chip from "@mui/joy/Chip";

export const TagSelectorRoot = styled(Box)`
	display: flex;
	flex-wrap: wrap;
	gap: ${spacing(1)};
`;

export const TagSelectorChip = styled(Chip)`
    cursor: pointer;
`;

