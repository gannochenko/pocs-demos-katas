import { BottomMenu } from "../BottomMenu";
import { Copyright } from "../Copyright";
import { useFooter } from "./hooks/useFooter";
import { FooterRoot, Left, Right, Wrapper } from "./style";
import { FooterProps } from "./type";

export const Footer = (props: FooterProps) => {
  useFooter(props);

  return (
    <FooterRoot>
      <Wrapper>
        <Left>
          <Copyright />
        </Left>
        <Right>
          <BottomMenu />
        </Right>
      </Wrapper>
    </FooterRoot>
  );
};
