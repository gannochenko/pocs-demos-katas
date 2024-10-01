import styled from '@emotion/styled';
import { Avatar, Popover, Typography } from '@mui/joy';
import { AuthWidgetRootPropsType } from './type';

export const AuthWidgetRoot = styled.div<AuthWidgetRootPropsType>`
    padding-left: ${spa(5)};
`;

export const AuthWidgetAvatar = styled(Avatar)`
    width: ${spacing(7.5)};
    height: ${muiSpacing(7.5)};
    cursor: pointer;
`;

export const AuthWidgetPopover = styled(Popover)`
    .MuiPopover-paper {
        margin-top: 0.7rem;
        width: 15rem;
        padding: 1rem;
    }
`;

export const AuthWidgetUserName = styled(Typography)`
    margin-bottom: 1rem !important;
`;
