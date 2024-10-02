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
import { siteMetadata } from '../../meta/site';

import { MenuPropsType } from './type';
import {Typography} from "../Typography";

// import { AuthWidget } from '../../../../bullshit/AuthWidget';

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
                <MenuHome to="/" onClick={onMobileItemClick}>
                    <Typography>
                        {siteMetadata.logoText}
                    </Typography>
                </MenuHome>
                <MenuRight>
                    <MenuItems>
                        {menu.map((item) => (
                            <MenuItem to={item.link} key={item.link}>
                                {item.text}
                            </MenuItem>
                        ))}
                    </MenuItems>

                    {/*<AuthWidget />*/}

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
                        to={item.link}
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
