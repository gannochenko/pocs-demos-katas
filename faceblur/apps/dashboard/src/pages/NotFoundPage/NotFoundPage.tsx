import React, { FC } from 'react';

import { NotFoundRoot, Message, Code, Explanation, Left, Image } from './style';
import { NotFoundFramePropsType } from './type';
import {Copyright, Link, PageLayout} from "../../components";
import {Container} from "@mui/joy";

export const NotFoundPage: FC<NotFoundFramePropsType> = ({
    children,
    ...restProps
}) => {
    return (
        <PageLayout title="Not found">
            <Container>
                <NotFoundRoot {...restProps}>
                    <Left>
                        <Image />
                        <Copyright
                            author="Zeynep"
                            source="https://unsplash.com/@zeynep_e"
                            sourceText="Unsplash"
                        />
                    </Left>
                    <Message>
                        <Code>404</Code>
                        <Explanation>
                            Not found. <Link href="/">Visit home page</Link>.
                        </Explanation>
                    </Message>
                </NotFoundRoot>
            </Container>
        </PageLayout>
    )
};
