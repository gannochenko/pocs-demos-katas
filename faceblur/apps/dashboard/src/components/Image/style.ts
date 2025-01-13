import {styled} from "@mui/joy";

type RootProps = {
	url: string;
};

export const Root = styled("div")<RootProps>`
    width: 100%;
    padding-bottom: 100%;
    position: relative;
    background-image: url(${(props) => props.url});
    background-position: center; /* Centers the image */
    background-repeat: no-repeat; /* Prevents the image from repeating */
    background-size: cover;
`;
