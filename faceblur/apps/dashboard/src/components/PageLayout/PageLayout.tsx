import React from 'react';
import Container from "@mui/joy/Container";
import { PageLayoutRoot } from './style';
import { PageLayoutPropsType } from './type';
import { SEO } from '../SEO';
import { Typography } from '../Typography';
import { usePageLayout } from './hooks/usePageLayout';

export function PageLayout(props: PageLayoutPropsType) {
    const { seoProps, showTitle, title, children } = usePageLayout(props);

    return (
        <PageLayoutRoot>
            <SEO {...seoProps} />
            {showTitle && (
                <Typography level="h1">
                    <br/>
                    <Container>{title}</Container>
                </Typography>
            )}
            <br/>
            {children}
        </PageLayoutRoot>
    );
};
