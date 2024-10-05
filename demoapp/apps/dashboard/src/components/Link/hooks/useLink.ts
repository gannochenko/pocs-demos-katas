import { LinkProps } from "../type";

export const useLink = ({ href, target, referrerPolicy, children, ...resProps }: LinkProps) => {
  let showRRLink = true;

  href = href ?? "/";

  if (!href.startsWith("/")) {
    target = "_blank";
    referrerPolicy = "no-referrer";
    showRRLink = true;
  }

  return {
    regularLinkProps: {
      ...resProps,
      href,
      target,
      rel: referrerPolicy,
    },
    routerLinkProps: {
      ...resProps,
      to: href,
    },
    showRRLink,
    children,
  };
};
