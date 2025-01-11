import fetchRetry from 'fetch-retry';

export type ErrorResponse = {
	error: string;
};

export function isError(error: any): error is ErrorResponse {
	return "error" in error;
}

export const apiUrl = process.env.REACT_APP_API_URL;

export const fetchWithRetry = fetchRetry(fetch, {
	retries: 5,
	retryDelay: function(attempt, error, response) {
		return Math.pow(2, attempt) * 1000;
	}
});

export const customFetch = async <I, O>(uri: string, body: I | null, token?: string): Promise<O | ErrorResponse> => {
	const headers: Record<string, string> = {
		'Content-Type': 'application/json',
	};
	if (token) {
		headers['Authorization'] = `Bearer ${token}`
	}

	try {
		const result = await fetchWithRetry(`${apiUrl}${uri}`, {
			method: 'POST',
			body: body ? JSON.stringify(body) : "{}",
			headers,
		});
		return await result.json();
	} catch (e) {
		return {
			error: (e as Error).message,
		};
	}
};
