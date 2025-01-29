import React, { forwardRef } from 'react';
import {Button, Typography} from '@mui/joy';

import { AuthWidgetPropsType } from './type';
import {
    AuthWidgetRoot,
    AuthWidgetAvatar,
    AuthWidgetUserName,
} from './style';
import { useAuthWidget } from './hooks/useAuthWidget';
import {Link} from "../Link";

export const AuthWidget = forwardRef<HTMLDivElement, AuthWidgetPropsType>(
    function AuthWidget(props, ref) {
        const {
            singInButtonProps,
            singOutButtonProps,
            isAuthenticated,
            rootProps,
            firstName,
        } = useAuthWidget(props);

        return (
            <AuthWidgetRoot {...rootProps}>
                {
                    !isAuthenticated
                    &&
                    <Link {...singInButtonProps}>Sign in</Link>
                }
                {
                    isAuthenticated
                    &&
                    <Typography>Welcome back, {firstName} | <Link {...singOutButtonProps}>Sign out</Link></Typography>
                }
            </AuthWidgetRoot>
        );
    },
);
