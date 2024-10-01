import { BottomMenuRoot } from "./style";
import { BottomMenuProps } from "./type";
import { useBottomMenu } from "./hooks/useBottomMenu";
import { Link } from "../Link";
import {pathTemplates} from "../../pathTemplates";

export const BottomMenu = (props: BottomMenuProps) => {
  const { children } = props;
  const res = useBottomMenu(props);

  return (
    <BottomMenuRoot>
      <Link href={pathTemplates.COOKIE_POLICY}>Cookie policy</Link>
      &nbsp;&bull;&nbsp;
      <Link href={pathTemplates.PERSONAL_DATA_POLICY}>Personal data policy</Link>
    </BottomMenuRoot>
  );
};
