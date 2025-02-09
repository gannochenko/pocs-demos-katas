import React, { FC } from 'react';
import Helmet from 'react-helmet';
import { SEOPropsType } from './type';
import {siteMeta} from "../../siteMeta";

export const SEO: FC<SEOPropsType> = ({
    description = '',
    lang = 'en',
    meta = [],
    keywords = [],
    title = '',
    image = '',
}) => {
    let metaTitle = '';
    let metaTitleOG = '';
    const metaDescription =
        description ?? siteMeta.description;

    if (title) {
        metaTitle = `${title} | ${siteMeta.title}`;
        metaTitleOG = title;
    } else {
        metaTitle = siteMeta.title;
        metaTitleOG = siteMeta.description;
    }

    let allKeywords: string[] = [];
    if (typeof keywords === 'string') {
        allKeywords = allKeywords.concat(
            keywords.split(',').map((word) => word.trim()),
        );
    }
    allKeywords = allKeywords
        .concat(siteMeta.keywords)
        .filter((x) => !!x);

    return (
        <Helmet
            htmlAttributes={{
                lang,
            }}
            title={metaTitle}
            // titleTemplate={
            //     title
            //         ? `%s | ${siteMetadata.title}`
            //         : metaTitle
            // }
            meta={[
                {
                    name: 'twitter:card',
                    content: 'summary',
                },
                {
                    name: 'twitter:creator',
                    content: '@gannochenko',
                },
                {
                    name: 'description',
                    content: metaDescription,
                },
                {
                    property: 'og:title',
                    content: metaTitleOG,
                },
                {
                    property: 'og:description',
                    content: metaDescription,
                },
                {
                    property: 'og:type',
                    content: 'website',
                },
                image
                    ? {
                        property: 'og:image',
                        content: image,
                    }
                    : {},
            ]
                .concat(
                    allKeywords.length > 0
                        ? [
                            {
                                name: `keywords`,
                                content: allKeywords.join(`, `),
                            },
                        ]
                        : [],
                )
                .concat(meta)
                .filter((x) => !!x)}
        />
    );
};
