import React, {PropsWithChildren} from 'react';
import { QueryClient, QueryClientProvider } from 'react-query';
import { SnackbarProvider } from 'notistack';
import { I18nextProvider } from 'react-i18next';
import { CssVarsProvider } from "@mui/joy/styles";
import { BrowserRouter } from "react-router-dom";

import { theme } from '../../style/theme';
import { StateProviders } from '../../states/providers';
import { NetworkStatusProvider } from '../NetworkStatusProvider';
import { i18n } from '../../i18n/i18n';
import {Auth0Provider} from "../Auth0Provider";

const queryClient = new QueryClient();

/**
 * This is a top-level wrapper, it wraps everything else, including the ApplicationLayout.
 */
export const Providers = ({ children }: PropsWithChildren) => {
    return (
        <BrowserRouter>
            <Auth0Provider>
                <CssVarsProvider theme={theme}>
                    <QueryClientProvider client={queryClient}>
                        <SnackbarProvider maxSnack={3}>
                            <NetworkStatusProvider>
                                <I18nextProvider i18n={i18n}>
                                    <StateProviders>
                                        {children}
                                    </StateProviders>
                                </I18nextProvider>
                            </NetworkStatusProvider>
                        </SnackbarProvider>
                    </QueryClientProvider>
                </CssVarsProvider>
            </Auth0Provider>
        </BrowserRouter>
    );
};
