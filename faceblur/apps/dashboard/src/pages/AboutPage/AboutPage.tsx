import React, { FC } from 'react';

import {AboutPageRoot} from './style';
import { AboutPageProps } from './type';
import {Link, PageLayout} from "../../components";
import {Container} from "@mui/joy";

export const AboutPage: FC<AboutPageProps> = ({
    children,
    ...restProps
}) => {
    return (
        <PageLayout title="About" displayPageTitle>
            <Container>
                <AboutPageRoot {...restProps}>
                    Hello
                </AboutPageRoot>
            </Container>
        </PageLayout>
    )
};
