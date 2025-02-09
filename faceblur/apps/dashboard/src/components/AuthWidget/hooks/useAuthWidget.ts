import { MouseEvent } from 'react';
import { useAuth0 } from "@auth0/auth0-react";
import { AuthWidgetPropsType } from '../type';
import {useNavigate} from "react-router-dom";

export const useAuthWidget = (
    props: AuthWidgetPropsType,
) => {
    const navigate = useNavigate();
    const { loginWithRedirect, logout, isAuthenticated, user } = useAuth0();
    
    return {
        rootProps: props,
        singInButtonProps: {
            onClick: async (e: MouseEvent<HTMLAnchorElement>) => {
                e.preventDefault();
                await loginWithRedirect({
                    appState: {
                        returnTo: "/",
                    },
                    authorizationParams: {
                        prompt: "login",
                    },
                });
            }
        },
        singOutButtonProps: {
            onClick: async (e: MouseEvent<HTMLAnchorElement>) => {
                e.preventDefault();
                await logout({
                    openUrl: () => {
                        navigate("/");
                    },
                });
            }
        },
        isAuthenticated,
        firstName: user?.given_name ?? "",
    };
};
