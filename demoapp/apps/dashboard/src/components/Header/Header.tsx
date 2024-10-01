import React, { FC } from 'react';
import { HeaderPropsType } from './type';
import { Menu } from '../Menu';
import {
    HeaderOffset,
    HeaderRoot,
} from './style';

export const Header: FC<HeaderPropsType> = () => (
    <HeaderRoot>
        <HeaderOffset />
        <Menu />
    </HeaderRoot>
);

export default Header;
