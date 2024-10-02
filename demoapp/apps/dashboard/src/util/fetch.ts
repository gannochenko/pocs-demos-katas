import {getToken} from "./token";

const apiUrl = process.env.REACT_APP_API_URL;

export const fetchJSON = (uri: string, body?: Record<string, unknown>) => {
	const token = getToken();

	// todo: retry upon invalid token
	return window.fetch(`${apiUrl}${uri}`, {
		method: 'POST',
		body: body ? JSON.stringify(body) : "",
		headers: {
			'Content-Type': 'application/json',
			// 'Authorization': `Bearer ${token}`,
		},
	}).then((result) => result.json());
};
