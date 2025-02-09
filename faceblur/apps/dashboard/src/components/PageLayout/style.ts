import {styled} from "@mui/joy";
import {typography} from "../../util/mixins";

export const PageLayoutRoot = styled("div")`
	${typography('body-md')};
    flex-grow: 1;
`;

// export const PageLayoutBackLink = styled(Link)`
//     text-decoration: none;
//     ${muiTypography('caption')};
// `;
