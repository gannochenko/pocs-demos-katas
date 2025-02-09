import { Theme } from "@mui/joy";

export type ThemedProps<P = {}, T = Theme> = {
	theme?: T;
} & P;

type Disallow<T> = {
	[key in keyof T]?: never;
};

export type MutuallyExclusive<T, U> = (T & Disallow<U>) | (U & Disallow<T>);
