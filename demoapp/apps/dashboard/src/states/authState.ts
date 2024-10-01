import { useCallback, useEffect, useState } from 'react';
import { createContainer } from 'unstated-next';
import { useQuery } from 'react-query';
import { getUserByToken } from '../services/auth';
import { UserRoleEnum, UserType } from '../services/user';
import { getToken, revokeToken, storeToken } from '../util/token';

const noop = () => {};

const useAuth = () => {
    const [, setSerial] = useState(1);

    const token = getToken();
    const hasToken = !!token;

    const { isSuccess, data: userQueryData } = useQuery(
        ['userData', token],
        // @ts-ignore
        getUserByToken,
        {
            enabled: hasToken,
            refetchOnMount: false,
            refetchOnReconnect: false,
            refetchOnWindowFocus: false,
            retry: false,
        },
    );

    const revoke = userQueryData?.revoke ?? false;
    const userData = userQueryData?.data;
    const userId = userData?.id;
    const userRoles = userData?.attributes?.roles ?? [];

    useEffect(() => {
        if (isSuccess && revoke) {
            revokeToken();
            setSerial((previousSerial) => previousSerial + 1);
        }
    }, [isSuccess, revoke]);

    const user: UserType | undefined =
        isSuccess && !revoke && !!userId
            ? {
                  id: userId,
                  roles: userRoles,
              }
            : undefined;

    const setToken = useCallback((token: string) => {
        storeToken(token);
        setSerial((previousSerial) => previousSerial + 1);
    }, []);

    const signOut = useCallback(() => {
        revokeToken();
        setSerial((previousSerial) => previousSerial + 1);
    }, []);

    return {
        user,
        signOut,
        signIn: noop,
        setToken,
        isAuthenticated: !!user?.id,
        isEditor:
            userRoles.includes(UserRoleEnum.contributor) ||
            userRoles.includes(UserRoleEnum.admin),
    };
};

export const AuthState = createContainer(useAuth);
