import { getToken } from './token';

export const fetchJSON = async (
    url: string,
    data?: Record<string, unknown>,
    headers?: Record<string, unknown>,
) => {
    const token = getToken();
    const response = await fetch(url, {
        method: 'POST',
        ...(data
            ? {
                  body: JSON.stringify(data),
              }
            : {}),
        headers: {
            'Content-Type': 'application/json',
            ...(token ? { token } : {}),
            ...(headers ?? {}),
        },
    });

    return response.json();
};
