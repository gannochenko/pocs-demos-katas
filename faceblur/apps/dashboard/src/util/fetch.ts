import fetchRetry from 'fetch-retry';
import axios from "axios";
import {ErrorResponse} from "@/proto/common/error/v1/error";

export function isErrorResponse(response: any): response is ErrorResponse {
	return "error" in response;
}

export const apiUrl = process.env.REACT_APP_API_URL;

export const getAPIURL = () => apiUrl;

export const fetchWithRetry = fetchRetry(fetch, {
	retries: 3,
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

export const uploadFile = async (
	url: string,
	file: File,
	onProgress: (progress: number) => void
): Promise<void> => {
	// using axios because it supports progress out of the box
	const response = await axios.put(url, file, {
		headers: {
			"Content-Type": "application/octet-stream",
		},
		onUploadProgress: (progressEvent) => {
			const progress = Math.round(
				(progressEvent.loaded / progressEvent.total!) * 100
			);
			onProgress(progress);
		},
	});

	return response.data;
};
