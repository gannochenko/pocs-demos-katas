export { spacing } from "../mixins/spacing";

type Scalar = string | number;

export type MarginProps = Partial<{
  marginTop: Scalar;
  marginBottom: Scalar;
  marginRight: Scalar;
  marginLeft: Scalar;
  margin: Scalar;
}>;

export type PaddingProps = Partial<{
  paddingTop: Scalar;
  paddingBottom: Scalar;
  paddingRight: Scalar;
  paddingLeft: Scalar;
  padding: Scalar;
}>;

export type SizeProps = Partial<{
  width: Scalar;
  height: Scalar;
  maxWidth: Scalar;
  maxHeight: Scalar;
  minWidth: Scalar;
  minHeight: Scalar;
}>;

const renderScalar = (value: Scalar) => {
  if (typeof value === "number") {
    return `${value}rem`; // todo: use theme.spacing() here
  }

  return value;
};

export const marginProps = ({
  marginBottom,
  marginLeft,
  marginRight,
  marginTop,
  margin,
}: MarginProps) => {
  let result = "";

  if (margin !== undefined) {
    result += `margin: ${renderScalar(margin)};`;
  }
  if (marginBottom !== undefined) {
    result += `margin-bottom: ${renderScalar(marginBottom)};`;
  }
  if (marginLeft !== undefined) {
    result += `margin-left: ${renderScalar(marginLeft)};`;
  }
  if (marginRight !== undefined) {
    result += `margin-right: ${renderScalar(marginRight)};`;
  }
  if (marginTop !== undefined) {
    result += `margin-top: ${renderScalar(marginTop)};`;
  }

  return result;
};

export const paddingProps = ({
  paddingBottom,
  paddingLeft,
  paddingRight,
  paddingTop,
  padding,
}: PaddingProps) => {
  let result = "";

  if (padding !== undefined) {
    result += `padding: ${renderScalar(padding)};`;
  }
  if (paddingBottom !== undefined) {
    result += `padding-bottom: ${renderScalar(paddingBottom)};`;
  }
  if (paddingLeft !== undefined) {
    result += `padding-left: ${renderScalar(paddingLeft)};`;
  }
  if (paddingRight !== undefined) {
    result += `padding-right: ${renderScalar(paddingRight)};`;
  }
  if (paddingTop !== undefined) {
    result += `padding-top: ${renderScalar(paddingTop)};`;
  }

  return result;
};

export const sizeProps = ({
  width,
  height,
  minWidth,
  minHeight,
  maxWidth,
  maxHeight,
}: SizeProps) => {
  let result = "";

  if (width !== undefined) {
    result += `width: ${renderScalar(width)};`;
  }
  if (height !== undefined) {
    result += `height: ${renderScalar(height)};`;
  }
  if (minWidth !== undefined) {
    result += `min-width: ${renderScalar(minWidth)};`;
  }
  if (minHeight !== undefined) {
    result += `min-height: ${renderScalar(minHeight)};`;
  }
  if (maxWidth !== undefined) {
    result += `max-width: ${renderScalar(maxWidth)};`;
  }
  if (maxHeight !== undefined) {
    result += `max-height: ${renderScalar(maxHeight)};`;
  }

  return result;
};
