import { useEffect, useState } from 'react';
import { NetworkStatusContextValueType } from '../type';
import { getWindow } from '../../../util/getWindow';

const win = getWindow();

export const useNetworkStatusProvider = () => {
    // @ts-ignore
    const [status, setState] = useState<NetworkStatusContextValueType>({
        online: true,
        setOnline: (isOnline: boolean) =>
            setState({
                ...status,
                online: isOnline,
            } as NetworkStatusContextValueType),
    });

    useEffect(() => {
        if (!win) {
            return;
        }

        const onOnline = () => {
            status.setOnline(true);
        };
        const onOffline = () => {
            status.setOnline(false);
        };

        win.addEventListener('online', onOnline);
        win.addEventListener('offline', onOffline);

        return () => {
            win.removeEventListener('online', onOnline);
            win.removeEventListener('offline', onOffline);
        };
    }, [status]);

    return status as NetworkStatusContextValueType;
};
