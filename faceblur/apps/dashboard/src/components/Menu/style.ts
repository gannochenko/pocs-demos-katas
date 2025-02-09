import { styled, Container } from "@mui/joy";

import {breakpointDown, breakpointUp, contentAlignment, spacing, typography} from "../../util/mixins";
import {Link} from "../Link";
import logoImage from './assets/logo.png';

export const MenuRoot = styled("div")`
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    background-color: var(--joy-palette-background-body);
    z-index: 1000;
    box-shadow: 0 6px 30px -8px rgba(0, 0, 0, 0.2);
`;

export const MenuInnerContainer = styled(Container)`
    ${contentAlignment('center', 'center')};
    justify-content: space-between;
    position: relative;
`;

export const MenuItems = styled("nav")`
    padding: ${spacing(1)};
    ${contentAlignment('right', 'center')}
    ${breakpointDown('md')} {
        display: none;
    }
`;

//  ${muiTypography('caption')};
export const MenuHome = styled(Link)`
    ${contentAlignment('center', 'center')}
    font-weight: bold;

    color: var(--joy-palette-background-text-default);
    &:hover {
        color: var(--joy-palette-background-text-default);
    }
    text-decoration: none;
    flex-shrink: 0;
    height: ${spacing(10)};
    transition: color 200ms ease;
`;

export const MenuItem = styled(Link)`
    color: var(--joy-palette-background-text-default);
    text-transform: uppercase;
    text-decoration: none;
    position: relative;
    padding-bottom: ${spacing(0.5)};

    &:after {
        content: '';
        display: block;
        height: 2px;
        position: absolute;
        bottom: 0;
        left: 0;
        width: 0;
        background-color: var(--joy-palette-background-text-default);
        transition: width ease 200ms;
    }

    &:hover {
        &:after {
            width: 100%;
        }
    }
`;

export const MenuRight = styled("div")`
    ${contentAlignment('center', 'center')};
    flex-shrink: 0;
    position: relative;
`;

export const MenuHamburger = styled("div")`
    ${contentAlignment('center', 'center', 'column')};
    width: ${spacing(10)};
    height: ${spacing(10)};
    padding: ${spacing(2)};
    cursor: pointer;
    ${breakpointUp('md')} {
        display: none;
    }
`;

export const MenuBar = styled("div")`
    background-color: var(--joy-palette-background-header);
    height: ${spacing(2.5)};
    width: 100%;
    display: block;
`;

export const MenuMobileItems = styled("nav")<{
    open: boolean;
}>`
    background-color: var(--joy-palette-background-header);
    position: absolute;
    top: 100%;
    right: ${({ open }) => (open ? '0' : '-100%')};
    width: 100%;
    height: 100vh;
    overflow: hidden;
    transition: right ease 200ms;
`;

export const MenuMobileItem = styled(Link)`
    padding: ${spacing(4)} ${spacing(8)};
    position: relative;
    display: block;
    text-decoration: none;
    color: var(--joy-palette-background-header);
    border-bottom: 1px solid var(--joy-palette-background-header);

    &:before {
        content: '';
        position: absolute;
        width: 0;
        height: 100%;
        top: 0;
        left: 0;
        bottom: 0;
        background-color: var(--joy-palette-background-header);
        transition: width 200ms ease;
    }

    &:hover {
        &:before {
            width: ${spacing(2.5)};
        }
    }
`;

export const Logo = styled(Link)`
    width: 50px;
    height: 50px;
    background-image: url(${logoImage});
    background-position: center; /* Centers the image */
    background-repeat: no-repeat; /* Prevents the image from repeating */
    background-size: 180%;
    margin-right: ${spacing(2)};
`;