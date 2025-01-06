import { styled } from "@mui/joy";
import { Link } from "react-router-dom";
import {typography} from "../../util/mixins";

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
  ${css};
  ${typography("body-md")};
`;

export const LinkRR = styled(Link)`
  ${css}
  ${typography("body-md")};
`;
