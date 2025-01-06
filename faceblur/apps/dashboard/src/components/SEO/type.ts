import { HTMLAttributes } from 'react';

export type SEOPropsType = Partial<{
    description: string;
    lang: string;
    meta: Meta[];
    keywords: string[] | string;
    title: string;
    image: string;
}> &
    HTMLAttributes<HTMLElement>;

export interface Meta {
    name: string;
    content: string;
}
