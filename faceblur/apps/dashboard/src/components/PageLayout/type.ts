import { HTMLAttributes } from 'react';

export type PageLayoutPropsType = Partial<{
    title: string;
    keywords: string[];
    description: string;
    displayPageTitle: boolean;
}> &
    HTMLAttributes<HTMLElement>;
