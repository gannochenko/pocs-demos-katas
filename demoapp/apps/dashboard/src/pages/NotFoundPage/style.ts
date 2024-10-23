import {styled} from "@mui/joy";
import {breakpointDown, spacing, typography} from "../../util/mixins";

const image01 = require('./assets/404.jpg') as string;

export const NotFoundRoot = styled("div")`
    display: flex;
    margin-top: ${spacing(16)};
`;

export const Image = styled("div")`
    background: url(${image01}) no-repeat;
    background-size: auto;
    width: 500px;
    height: 667px;
`;

export const Message = styled("div")`
    padding-left: ${spacing(8)};
`;

export const Code = styled("div")`
    font-size: ${spacing(40)};
    line-height: 0.8;
`;

export const Explanation = styled("div")`
    ${typography('body-md')};
    margin-top: ${spacing(4)};
`;

export const Left = styled("div")`
    position: relative;
    ${breakpointDown('sm')} {
        display: none;
    }
`;
