import { HTMLAttributes } from 'react';

export type CopyrightPropsType = Partial<{
    // custom props here

    author: string;
    source: string;
    sourceText: string;
}> &
    HTMLAttributes<HTMLElement>;
