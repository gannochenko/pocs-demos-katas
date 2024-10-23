import React, { FC, useState, useCallback } from 'react';

import {
    MenuRoot,
    MenuInnerContainer,
    MenuItems,
    MenuItem,
    MenuHome,
    MenuHamburger,
    MenuBar,
    MenuRight,
    MenuMobileItems,
    MenuMobileItem,
} from './style';

import { menu } from '../../menu';
import { siteMeta } from '../../siteMeta';

import { MenuPropsType } from './type';
import {Typography} from "../Typography";
import {AuthWidget} from "../AuthWidget";

export const Menu: FC<MenuPropsType> = () => {
    const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
    const onHamburgerClick = useCallback(() => {
        setMobileMenuOpen(!mobileMenuOpen);
    }, [mobileMenuOpen, setMobileMenuOpen]);
    const onMobileItemClick = useCallback(() => {
        setMobileMenuOpen(false);
    }, [setMobileMenuOpen]);

    return (
        <MenuRoot>
            <MenuInnerContainer>
                <MenuHome href="/" onClick={onMobileItemClick}>
                    <Typography>
                        {siteMeta.logoText}
                    </Typography>
                </MenuHome>
                <MenuRight>
                    <MenuItems>
                        {menu.map((item) => (
                            <MenuItem href={item.link} key={item.link}>
                                {item.text}
                            </MenuItem>
                        ))}
                    </MenuItems>

                    <AuthWidget />

                    {!!menu.length && (
                        <MenuHamburger onClick={onHamburgerClick}>
                            <MenuBar />
                            <MenuBar />
                            <MenuBar />
                        </MenuHamburger>
                    )}
                </MenuRight>
            </MenuInnerContainer>
            <MenuMobileItems open={mobileMenuOpen}>
                {menu.map((item) => (
                    <MenuMobileItem
                        href={item.link}
                        key={item.link}
                        onClick={onMobileItemClick}
                    >
                        {item.text}
                    </MenuMobileItem>
                ))}
            </MenuMobileItems>
        </MenuRoot>
    );
};
