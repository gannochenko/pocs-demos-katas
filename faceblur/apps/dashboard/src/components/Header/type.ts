import { HTMLAttributes } from 'react';

export type HeaderPropsType = Partial<{
    inner: boolean;
}> &
    HTMLAttributes<HTMLElement>;
