import { AnchorHTMLAttributes, PropsWithChildren } from "react";

export type LinkProps = Partial<PropsWithChildren<{}>> &
  AnchorHTMLAttributes<HTMLAnchorElement>;
