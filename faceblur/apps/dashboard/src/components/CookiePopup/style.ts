import { ZINDEX_HIGHER } from "../../util/constants";
import { breakpointDown } from "../../util/mixins";
import { spacing } from "../../util/props";
import { ThemedProps } from "../../util/type";
import {Button, styled, Typography} from "@mui/joy";

import image from "./assets/cookies.jpg";

export const CookiePopupRoot = styled("div")<
  ThemedProps<{
    fadingAway: boolean;
  }>
>`
  display: flex;
  flex-direction: row;
  align-items: stretch;
  position: fixed;
  bottom: ${spacing(2)};
  background-color: var(--joy-palette-background-body);
  border-bottom-left-radius: 2px;
  border-top-left-radius: 2px;
  overflow: hidden;
  box-shadow: 0 6px 16px 0 rgba(0, 0, 0, 0.1);
  right: ${({ fadingAway }) => (!fadingAway ? 0 : "-1rem")};
  opacity: ${({ fadingAway }) => (!fadingAway ? 1 : 0)};
  transition: right 500ms ease, opacity 500ms ease;
  z-index: ${ZINDEX_HIGHER};
`;

//background-image: url(${image.src});
export const CookiePopupPicture = styled("div")`
  background-repeat: no-repeat;
  background-position: center center;
  background-attachment: scroll;
  width: ${spacing(30)};
  ${breakpointDown("md")} {
    display: none;
  }
  position: relative;
  &:hover > * {
    opacity: 1;
  }
`;

export const CookiePopupText = styled(Typography)`
  padding: ${spacing(2)} ${spacing(4)};
  position: relative;
`;

export const CookiePopupAgreeButton = styled(Button)`
  position: absolute;
  right: ${spacing(4)};
  bottom: ${spacing(2)};
  ${breakpointDown("md")} {
    display: none;
  }
`;

export const CookiePopupAgreeButtonXS = styled(Button)`
  display: block;
  margin-top: ${spacing(2)};
  // ${breakpointDown("md")} {
  //   display: block;
  // }
`;
