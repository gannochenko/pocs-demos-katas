import { breakpointDown, spacing } from "../../util/mixins";
import { styled } from "@mui/joy";

export const FooterRoot = styled("div")``;

export const Wrapper = styled("div")`
  padding: ${spacing(5)} ${spacing(2)};
  flex-direction: column;
  text-align: center;
`;

export const Left = styled("div")`
  margin-right: ${spacing(2)};
  ${breakpointDown("md")} {
    margin-right: 0;
    margin-bottom: ${spacing(3)};
  }
`;

export const Center = styled("div")``;

export const Right = styled("div")`
  margin-left: ${spacing(2)};
  ${breakpointDown("md")} {
    margin-left: 0;
    margin-top: ${spacing(3)};
  }
`;
