import { PropsWithChildren } from "react";
import { TypographyProps as TP } from "@mui/joy";

export type TypographyProps = Partial<
  PropsWithChildren<{
    muted: boolean;
    megaHeader: boolean;
  }>
> &
  TP;

export type TypographyRootProps = Partial<{
  muted: boolean;
  megaHeader: boolean;
}>;
