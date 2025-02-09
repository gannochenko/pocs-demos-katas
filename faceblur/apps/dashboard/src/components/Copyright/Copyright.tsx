import React, { FC } from 'react';

import { Link } from '../Link';
import { CopyrightRoot } from './style';
import { CopyrightPropsType } from './type';

export const Copyright: FC<CopyrightPropsType> = ({
    author,
    source,
    sourceText,
}) => {
    if (!author && !source) {
        return null;
    }

    return (
        <CopyrightRoot>
            {!!author && <span>Photo by {author}</span>}
            {!!source && (
                <span>
                    {author ? ' on ' : ''}
                    <Link href={source} target="_blank" referrerPolicy="noopener noreferrer">
                        {sourceText || source}
                    </Link>
                </span>
            )}
        </CopyrightRoot>
    );
};
