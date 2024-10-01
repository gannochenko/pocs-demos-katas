import { useRef, ForwardedRef, RefObject } from 'react';

export const useRootRef = <E extends HTMLElement>(ref: ForwardedRef<E>) => {
    const localRef = useRef<E>(null);

    // we only allow modern reference syntax
    return (ref as RefObject<E>) || localRef;
};
