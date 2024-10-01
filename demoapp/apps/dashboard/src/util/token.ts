const TOKEN_LS_KEY = 'demoapp:token';

export const storeToken = (token: string) => {
    if (typeof window === 'undefined' || !token) {
        return;
    }

    window.localStorage.setItem(TOKEN_LS_KEY, token);
};

export const getToken = () => {
    if (typeof window === 'undefined') {
        return null;
    }

    return window.localStorage.getItem(TOKEN_LS_KEY);
};

export const revokeToken = () => {
    if (typeof window === 'undefined') {
        return;
    }
    window.localStorage.removeItem(TOKEN_LS_KEY);
};
