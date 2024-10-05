import { LinkRR, LinkRegular } from "./style";
import { LinkProps } from "./type";
import { useLink } from "./hooks/useLink";
import {Typography} from "@mui/joy";

export const Link = (props: LinkProps) => {
  const { regularLinkProps, routerLinkProps, showRRLink, children } = useLink(props);

  if (showRRLink) {
    return (
      <LinkRR {...routerLinkProps}>
        {children}
      </LinkRR>
    );
  }

  return (
    <LinkRegular href={regularLinkProps.href} target={regularLinkProps.target} rel={regularLinkProps.rel}>
      {children}
    </LinkRegular>
  );
};
