import { styled, Typography } from "@mui/joy";
import { TypographyRootProps } from "./type";

export const TypographyRoot = styled(Typography)<TypographyRootProps>`
  ${({ muted, theme }) => {
    if (muted) {
      return `color: ${theme.palette.text.secondary};`;
    }

    return "";
  }}

  ${({ megaHeader }) => {
    if (megaHeader) {
      return `
        font-family: "Montserrat", monospace;
        font-size: 1.6rem;
      `;
    }

    return "";
  }}
`;
