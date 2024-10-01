import { ThemedProps } from "../type";

export const spacing =
  (value: number) =>
  ({ theme }: ThemedProps) =>
    theme?.spacing(value) ?? "";
