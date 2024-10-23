import React, { FC } from 'react';
import { ApplicationLayoutRoot, ApplicationLayoutBody } from './style';
import { LayoutPropsType } from './type';
import { Header, Footer } from '../index';
import { CookiePopup } from '../CookiePopup';

/**
 * This is a top-level UI layout, it wraps every other visible element of the project.
 * Above it there is the Providers wrapper only.
 */
export const ApplicationLayout: FC<LayoutPropsType> = ({ children }) => {
    return (
        <ApplicationLayoutRoot>
            <Header />
            <ApplicationLayoutBody>{children}</ApplicationLayoutBody>
            <Footer />
            <CookiePopup />
        </ApplicationLayoutRoot>
    );
};
