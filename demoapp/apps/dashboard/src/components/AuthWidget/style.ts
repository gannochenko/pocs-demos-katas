import {styled} from "@mui/joy";
import { Avatar, Typography } from '@mui/joy';
import {spacing} from "../../util/mixins";

export const AuthWidgetRoot = styled("div")`
    padding-left: ${spacing(5)};
    padding-bottom: ${spacing(0.5)};
`;

export const AuthWidgetAvatar = styled(Avatar)`
    width: ${spacing(7.5)};
    height: ${spacing(7.5)};
    cursor: pointer;
`;

export const AuthWidgetUserName = styled(Typography)`
    margin-bottom: 1rem !important;
`;
