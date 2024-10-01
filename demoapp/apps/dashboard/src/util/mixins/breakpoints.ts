import { ThemedProps } from "../type";

// https://mui.com/joy-ui/customization/default-theme-viewer/
export type BreakpointNameType = "lg" | "sm" | "md" | "xs" | "xl";

export const breakpointUp =
  (breakpointName: BreakpointNameType) =>
  ({ theme }: ThemedProps) =>
    theme?.breakpoints.up(breakpointName) ?? "@media";

export const breakpointDown =
  (breakpointName: BreakpointNameType) =>
  ({ theme }: ThemedProps) =>
    theme?.breakpoints.down(breakpointName) ?? "@media";

export const breakpointBetween =
  (
    breakpointStartName: BreakpointNameType,
    breakpointEndName: BreakpointNameType
  ) =>
  ({ theme }: ThemedProps) =>
    theme?.breakpoints.between(breakpointStartName, breakpointEndName) &&
    "@media";
