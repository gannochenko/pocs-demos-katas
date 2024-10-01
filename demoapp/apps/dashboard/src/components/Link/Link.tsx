import { LinkRR, LinkRegular } from "./style";
import { LinkProps } from "./type";
import { useLink } from "./hooks/useLink";

export const Link = (props: LinkProps) => {
  const { rootProps, showRRLink } = useLink(props);

  if (showRRLink) {
    return <LinkRR {...rootProps} />;
  }

  return <LinkRegular {...rootProps} />;
};
