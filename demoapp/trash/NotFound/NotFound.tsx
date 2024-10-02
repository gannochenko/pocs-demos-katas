import React, { FC } from 'react';

import { NotFoundRoot, Image, Message, Code, Explanation, Left } from './style';
import { NotFoundFramePropsType } from './type';
import { Link } from '../../apps/dashboard/src/components/Link';
import { Copyright } from '../../apps/dashboard/src/components/Copyright';

export const NotFound: FC<NotFoundFramePropsType> = ({
    children,
    ...restProps
}) => {
    return (
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
                    Not found. <Link to="/">Visit home page</Link>.
                </Explanation>
            </Message>
        </NotFoundRoot>
    );
};
