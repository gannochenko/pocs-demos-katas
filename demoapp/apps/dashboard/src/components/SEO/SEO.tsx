import React, { FC } from 'react';
import Helmet from 'react-helmet';
import { SEOPropsType } from './type';
import { siteMetadata } from "../../meta/site";

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
        description || siteMetadata.description;

    if (title) {
        metaTitle = `${title} | ${siteMetadata.title}`;
        metaTitleOG = title;
    } else {
        metaTitle = siteMetadata.title;
        metaTitleOG = siteMetadata.description;
    }

    let allKeywords: string[] = [];
    if (typeof keywords === 'string') {
        allKeywords = allKeywords.concat(
            keywords.split(',').map((word) => word.trim()),
        );
    }
    allKeywords = allKeywords
        .concat(siteMetadata.keywords)
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
