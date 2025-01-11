import React, { FC, useCallback, useState } from "react";
import { Link } from "../Link";

import {
  CookiePopupAgreeButton,
  CookiePopupAgreeButtonXS,
  CookiePopupPicture,
  CookiePopupRoot,
  CookiePopupText,
} from "./style";
import { CookiePopupPropsType } from "./type";

export const getWindow = () => (typeof window !== "undefined" ? window : null);
const win = getWindow();

export const CookiePopup: FC<CookiePopupPropsType> = () => {
  const [displayed, setDisplayed] = useState(
    win ? !win.localStorage.getItem("cookie-accept") : false
  );

  const [fadingAway, setFadingAway] = useState(false);

  const onAcceptClick = useCallback(() => {
    if (typeof window === "undefined") {
      return;
    }

    window.localStorage.setItem("cookie-accept", "1");
    setFadingAway(true);
    setTimeout(() => setDisplayed(false), 600);
  }, []);

  if (!displayed) {
    return null;
  }

  return (
    <CookiePopupRoot fadingAway={fadingAway}>
      {/*<CookiePopupPicture />*/}
      <CookiePopupText>
        I use <b>cookies</b> to improve your experience with
        my website.
        <br />
        By further browsing you agree to accept the cookies.
        <br />
        More information <Link href="/cookie-policy">here</Link>.
        <div>
          <CookiePopupAgreeButtonXS onClick={onAcceptClick}>
            Accept!
          </CookiePopupAgreeButtonXS>
        </div>
      </CookiePopupText>
    </CookiePopupRoot>
  );
};

export default CookiePopup;
