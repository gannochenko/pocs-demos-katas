import {styled} from "@mui/joy";
import {spacing} from "../../util/mixins";

export const Root = styled("div")`
	margin-top: ${spacing(2)};
	margin-bottom: ${spacing(6)};
`;

type ImageItemProps = {
	image: string;
};

export const ImageItem = styled("div")<ImageItemProps>`
    width: 100%;
    padding-bottom: 100%;
    position: relative;
    background-image: url(${(props) => props.image});
    background-position: center; /* Centers the image */
    background-repeat: no-repeat; /* Prevents the image from repeating */
    background-size: cover;
`;
