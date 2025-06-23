import { AxiosResponse } from 'axios';
import { firstValueFrom, map, Observable, retry, timeout, timer } from 'rxjs';

export const retryRequest = <T>(
  response$: Observable<AxiosResponse<T, any>>,
  timeoutValue = 5000,
  retryCountValue = 10,
): Promise<T> => {
  return firstValueFrom(
    response$.pipe(
      timeout(timeoutValue),
      retry({
        count: retryCountValue,
        delay: (error, retryCount) => timer(retryCount * 1000),
      }),
      map((response) => response.data),
    ),
  );
};
