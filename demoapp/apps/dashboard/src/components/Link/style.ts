import { styled } from "@mui/joy";
import { Link } from "react-router-dom";

const css = `
  color: inherit;
  &:hover,
  &:visited,
  &:active,
  &:focus {
    color: inherit;
  }
  text-decoration: underline;
  &:hover {
    text-decoration: none;
  }
`;

export const LinkRegular = styled("a")`
  ${css}
`;

export const LinkRR = styled(Link)`
  ${css}
`;
