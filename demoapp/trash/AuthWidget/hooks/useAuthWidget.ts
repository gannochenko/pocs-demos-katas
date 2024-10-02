import { Ref, useCallback } from 'react';
import { AuthWidgetPropsType } from '../type';
import { AuthState } from '../../../apps/dashboard/src/states';

export const useAuthWidget = (
    ref: Ref<HTMLDivElement>,
    props: AuthWidgetPropsType,
) => {
    const { signOut, signIn, isAuthenticated, user } = AuthState.useContainer();

    // console.log('USER');
    // console.log(user);

    const color = 'primary';
    const variant: 'text' | 'outlined' | 'contained' = 'contained';

    const onSingUpButtonClick = useCallback(() => signIn(), [signIn]);
    const onSingOutButtonClick = useCallback(() => signOut(), [signOut]);

    return {
        rootProps: {
            ...props, // rest props go to the root node, as before
            ref, // same for the ref
        },
        signUpButtonProps: {
            onClick: onSingUpButtonClick,
            color,
            variant,
        },
        signOutButtonProps: {
            onClick: onSingOutButtonClick,
            // color,
            // variant,
        },
        getInvitationRequestButtonProps: () => ({
            // color,
            // variant,
        }),
        avatarProps: {
            src: '',
        },
        userNameProps: {
            // children: userId,
        },
        authenticated: isAuthenticated,
    };
};
