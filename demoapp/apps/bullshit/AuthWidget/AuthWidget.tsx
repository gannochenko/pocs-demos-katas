import React, { forwardRef } from 'react';
import PopupState, { bindTrigger, bindPopover } from 'material-ui-popup-state';
import { Button } from '@mui/joy';

import { AuthWidgetPropsType } from './type';
import {
    AuthWidgetRoot,
    AuthWidgetAvatar,
    AuthWidgetPopover,
    AuthWidgetUserName,
} from './style';
import { useAuthWidget } from './hooks/useAuthWidget';
import { Link } from '../../dashboard/src/components/Link';

const anchorOrigin = {
    vertical: 'bottom',
    horizontal: 'right',
};

const transformOrigin = {
    vertical: 'top',
    horizontal: 'right',
};

export const AuthWidget = forwardRef<HTMLDivElement, AuthWidgetPropsType>(
    function AuthWidget(props, ref) {
        const {
            rootProps,
            signUpButtonProps,
            signOutButtonProps,
            avatarProps,
            userNameProps,
            authenticated,
            getInvitationRequestButtonProps,
        } = useAuthWidget(ref, props);

        return (
            <AuthWidgetRoot {...rootProps}>
                <PopupState variant="popover" popupId="auth-popover" disableAutoFocus>
                    {(popupState) => (
                        <>
                            <AuthWidgetAvatar
                                alt="Remy Sharp"
                                {...bindTrigger(popupState)}
                                {...avatarProps}
                            />
                            <AuthWidgetPopover
                                {...bindPopover(popupState)}
                                anchorOrigin={anchorOrigin}
                                transformOrigin={transformOrigin}
                            >
                                {authenticated && (
                                    <>
                                        <AuthWidgetUserName {...userNameProps}>
                                            Вы авторизованы
                                        </AuthWidgetUserName>
                                        <Button {...signOutButtonProps}>
                                            Выйти
                                        </Button>
                                    </>
                                )}
                                {!authenticated && (
                                    // <Button {...signUpButtonProps}>
                                    //     Log in / Sign up
                                    // </Button>
                                    <div>
                                        Вы можете войти по приглашению.
                                        <br />
                                        <br />
                                        <Link to="/contacts">
                                            <Button
                                                {...getInvitationRequestButtonProps()}
                                            >
                                                Получить приглашение
                                            </Button>
                                        </Link>
                                    </div>
                                )}
                            </AuthWidgetPopover>
                        </>
                    )}
                </PopupState>
            </AuthWidgetRoot>
        );
    },
);
