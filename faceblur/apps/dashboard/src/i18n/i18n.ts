import i18nFactory from 'i18next';
import { initReactI18next } from 'react-i18next';
import { en } from './en';
// import LanguageDetector from 'i18next-browser-languagedetector';

// https://dev.to/adrai/how-to-properly-internationalize-a-react-application-using-i18next-3hdb
i18nFactory
    // // detect user language
    // // learn more: https://github.com/i18next/i18next-browser-languageDetector
    // .use(LanguageDetector)
    // pass the i18n instance to react-i18next.
    .use(initReactI18next)
    // init i18next
    // for all options read: https://www.i18next.com/overview/configuration-options
    .init({
        debug: false,
        fallbackLng: 'en',
        interpolation: {
            escapeValue: false, // not needed for react as it escapes by default
        },
        resources: {
            en,
        },
    });

export const i18n = i18nFactory;
