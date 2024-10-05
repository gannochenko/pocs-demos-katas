import { ThemedProps } from "../type";
import {Breakpoint, Breakpoints} from "@mui/system/createTheme/createBreakpoints";

// https://mui.com/joy-ui/customization/default-theme-viewer/

export const breakpointUp =
  (breakpointName: Breakpoint) =>
  ({ theme }: ThemedProps) =>
    theme?.breakpoints.up(breakpointName) ?? "@media";

export const breakpointDown =
  (breakpointName: Breakpoint) =>
  ({ theme }: ThemedProps) =>
    theme?.breakpoints.down(breakpointName) ?? "@media";

export const breakpointBetween =
  (
    breakpointStartName: Breakpoint,
    breakpointEndName: Breakpoint
  ) =>
  ({ theme }: ThemedProps) =>
    theme?.breakpoints.between(breakpointStartName, breakpointEndName) &&
    "@media";
