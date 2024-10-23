import { useContext, createContext } from 'react';
import { NetworkStatusContextValueType } from './type';

export const NetworkStatusContext =
    createContext<NetworkStatusContextValueType>({
        online: true,
        setOnline: () => {},
    });

export const NetworkStatusContextProvider = NetworkStatusContext.Provider;

export const useNetworkStatus = () => useContext(NetworkStatusContext);
