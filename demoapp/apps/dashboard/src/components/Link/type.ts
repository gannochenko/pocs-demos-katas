import {MouseEvent, PropsWithChildren} from "react";

export type LinkProps = Partial<PropsWithChildren<{
	href: string;
	target: string;
	referrerPolicy: string;
	onClick: (e: MouseEvent<HTMLAnchorElement>) => void;
}>>;
