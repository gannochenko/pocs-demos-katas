import styled from '@emotion/styled';
import {
    muiTypography,
    muiColor,
    muiSpacing,
    backgroundCover,
    muiBreakpointDown,
} from '@gannochenko/ui.emotion';

// eslint-disable-next-line @typescript-eslint/no-var-requires
const image01 = require('../../../../static/assets/404/01.jpg') as string;

export const NotFoundRoot = styled.div`
    display: flex;
    margin-top: ${muiSpacing(16)};
`;

export const Image = styled.div`
    ${backgroundCover(image01)};
    width: ${muiSpacing(120)};
    height: ${muiSpacing(120)};
`;

export const Message = styled.div`
    padding-left: ${muiSpacing(8)};
    color: ${muiColor('text.primary')};
`;

export const Code = styled.div`
    font-size: ${muiSpacing(40)};
    line-height: 0.8;
`;

export const Explanation = styled.div`
    ${muiTypography('h6')};
    margin-top: ${muiSpacing(4)};
`;

export const Left = styled.div`
    position: relative;
    ${muiBreakpointDown('sm')} {
        display: none;
    }
`;
