import React, { FC } from 'react';
import Container from "@mui/joy/Container";
import { PageLayoutRoot } from './style';
import { PageLayoutPropsType } from './type';
import { SEO } from '../SEO';
import { Typography } from '../Typography';
import { usePageLayout } from './hooks/usePageLayout';

export const PageLayout: FC<PageLayoutPropsType> = (props) => {
    const { seoProps, showTitle, title, children } = usePageLayout(props);

    return (
        <PageLayoutRoot>
            <SEO {...seoProps} />
            {showTitle && (
                <Typography variant="plain" component="h1">
                    <br/>
                    <Container>{title}</Container>
                </Typography>
            )}
            {children}
        </PageLayoutRoot>
    );
};
