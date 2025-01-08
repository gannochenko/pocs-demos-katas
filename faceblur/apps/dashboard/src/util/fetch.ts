import fetchRetry from 'fetch-retry';

export const apiUrl = process.env.REACT_APP_API_URL;

export const fetchJSON = async (uri: string, body: Record<string, unknown> | null, token: string) => {
	const headers: Record<string, string> = {
		'Content-Type': 'application/json',
	};
	if (token) {
		headers['Authorization'] = `Bearer ${token}`
	}

	const result = await window.fetch(`${apiUrl}${uri}`, {
		method: 'POST',
		body: body ? JSON.stringify(body) : "{}",
		headers,
	});
	return await result.json();
};

export type ErrorResponse = {
	error: string;
};

export const fetchWithRetry = fetchRetry(fetch, {
	retries: 5,
	retryDelay: function(attempt, error, response) {
		return Math.pow(2, attempt) * 1000;
	}
});
