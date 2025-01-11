import React from 'react';
import Container from "@mui/joy/Container";
import Box from "@mui/joy/Box";
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
                    <Container>
                        <Box sx={{display: 'flex', alignItems: 'center', justifyContent: 'space-between'}}>
                            {title}
                            <div id="page-header-portal" />
                        </Box>
                    </Container>
                </Typography>
            )}
            <br/>
            {children}
        </PageLayoutRoot>
    );
};
