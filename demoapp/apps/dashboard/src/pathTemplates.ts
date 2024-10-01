export const pathTemplates = {
    HOME: '/',
    PERSONAL_DATA_POLICY: '/personal-data-policy',
    COOKIE_POLICY: '/cookie-policy',
    ABOUT: '/about',

    fillTemplate: (template: string, values: Record<string, string>) => {
        let result = template;

        if (values) {
            Object.keys(values).forEach((key) => {
                const value = values[key];
                result = result
                    .replace(`#${key}#`, value)
                    .replace(`#${key.toUpperCase()}#`, value);
            });
        }

        return result;
    },
};
