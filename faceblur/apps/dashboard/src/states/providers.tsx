import React, {PropsWithChildren} from 'react';

import { NoopState } from '.';

export const StateProviders = ({ children }: PropsWithChildren) => {
    return (
        <NoopState.Provider>{children}</NoopState.Provider>
    );
};
