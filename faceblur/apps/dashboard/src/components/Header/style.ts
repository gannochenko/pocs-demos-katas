import { styled } from "@mui/joy";
import {spacing} from "../../util/mixins";

export const HeaderRoot = styled("header")`
    position: relative;
    min-width: 320px;
    flex-shrink: 0;
	z-index: 100;
`;

export const HeaderOffset = styled("div")`
    height: ${spacing(10)};
`;
