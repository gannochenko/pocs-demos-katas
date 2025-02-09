import { ThemedProps } from "../type";
import {DefaultTypographySystem} from "@mui/joy/styles/types/typography";

export const typography =
  (value: keyof DefaultTypographySystem) =>
  ({ theme }: ThemedProps) => theme?.typography[value];
