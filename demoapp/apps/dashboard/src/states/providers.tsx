import React, {PropsWithChildren} from 'react';

import { AuthState } from '.';

export const StateProviders = ({ children }: PropsWithChildren) => {
    return (
        <AuthState.Provider>{children}</AuthState.Provider>
    );
};
