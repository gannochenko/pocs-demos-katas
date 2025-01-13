import {styled} from "@mui/joy";
import ReportIcon from '@mui/icons-material/Report';
import {spacing} from "../../util/mixins";

export const Root = styled("div")`
    width: 100%;
    position: relative;
`;

export const Inner = styled("div")`
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
	display: flex;
	justify-content: center;
	align-items: center;
`;

export const Status = styled("div")`
	display: flex;
	flex-direction: column;
    align-items: center;
`;

export const Label = styled("div")`
	margin-top: ${spacing(2)};
`;

export const ErrorIcon = styled(ReportIcon)`
    font-size: 34px;
    color: #e83b3b;
`;