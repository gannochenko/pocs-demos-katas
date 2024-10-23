import React, {PropsWithChildren} from 'react';

import { useNetworkStatusProvider } from './hooks/useNetworkStatus';
import { NetworkStatusContextProvider } from './context';

export const NetworkStatusProvider = ({ children }: PropsWithChildren) => {
    const status = useNetworkStatusProvider();
    return (
        <NetworkStatusContextProvider value={status}>
            {children}
        </NetworkStatusContextProvider>
    );
};
