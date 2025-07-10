export const BASE_URL = "https://api.weatherapi.com/v1";

export async function customFetch(
  url: string,
  options: RequestInit = {},
  timeoutMs: number = 10000
): Promise<any> {
  const controller = new AbortController();
  const timeoutId = setTimeout(() => controller.abort(), timeoutMs);

  try {
    const response = await fetch(url, {
      ...options,
      signal: controller.signal,
      headers: {
        "Content-Type": "application/json",
        ...options.headers,
      },
    });

    if (!response.ok) {
      throw new Error(
        `HTTP error! status: ${response.status} - ${response.statusText}`
      );
    }

    const data = await response.json();
    return data;
  } catch (error) {
    if (error instanceof Error && error.name === "AbortError") {
      throw new Error(
        `Request timeout: The API request took longer than ${timeoutMs}ms`
      );
    }
    if (error instanceof TypeError && error.message.includes("fetch")) {
      throw new Error(
        `Network error: Unable to reach the API. Please check your internet connection.`
      );
    }
    throw error;
  } finally {
    clearTimeout(timeoutId);
  }
}
