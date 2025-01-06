import { HTMLAttributes } from 'react';

export type LayoutPropsType = Partial<{
    // custom props here

    props: {
        location: {
            pathname: string;
        };
    };
}> &
    HTMLAttributes<HTMLElement>;
