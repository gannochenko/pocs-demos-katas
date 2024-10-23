import { PageLayoutPropsType } from '../type';

export const usePageLayout = ({
    title: titleProp = '',
    keywords: keywordsProp = [],
    description: descriptionProp = '',
    displayPageTitle: displayPageTitleProp,
    children,
}: PageLayoutPropsType) => {
    const actualTitle = titleProp ?? "";
    const actualKeywords = keywordsProp ?? [];
    const actualDescription = descriptionProp ?? "";

    const showTitle = actualTitle && displayPageTitleProp;

    return {
        seoProps: {
            title: actualTitle,
            keywords: actualKeywords,
            description: actualDescription,
        },
        showTitle,
        title: actualTitle,
        children,
    };
};
