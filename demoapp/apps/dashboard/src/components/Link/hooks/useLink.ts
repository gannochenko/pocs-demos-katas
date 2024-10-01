import { LinkProps } from "../type";

export const useLink = (props: LinkProps) => {
  let { href, target, referrerPolicy } = props;
  let showRRLink = true;

  href = href ?? "/";

  if (!href.startsWith("/")) {
    target = "_blank";
    referrerPolicy = "no-referrer";
    showRRLink = true;
  }

  return {
    rootProps: {
      ...props,
      href,
      to: href,
      target,
      referrerPolicy,
    },
    showRRLink,
  };
};
