import {Card, styled} from "@mui/joy";
import {spacing} from "../../util/mixins";
import ReportIcon from "@mui/icons-material/Report";

export const Root = styled("div")`
    width: 100%;
    position: relative;
`;

type ImageContainerProps = {
	url: string;
	isProcessed: boolean;
};

export const ImageContainer = styled("div")<ImageContainerProps>`
    width: 100%;
    padding-bottom: 100%;
    position: relative;
    background-image: url(${(props) => props.url});
    filter: blur(5px);
    background-position: center; /* Centers the image */
    background-repeat: no-repeat; /* Prevents the image from repeating */
    background-size: cover;
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

export const Underlay = styled(Card)`
    display: flex;
    flex-direction: column;
	justify-content: center;
    align-items: center;
	width: ${spacing(15)};
	height: ${spacing(15)};
`;