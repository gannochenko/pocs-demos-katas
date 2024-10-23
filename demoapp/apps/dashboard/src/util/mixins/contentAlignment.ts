type Alignment =
  | "start"
  | "end"
  | "left"
  | "right"
  | "top"
  | "bottom"
  | "middle"
  | "flex-start"
  | "flex-end"
  | "center";

const replaceAlignment = (how?: Alignment) => {
  if (how === "start" || how === "left" || how === "top") {
    return "flex-start";
  }
  if (how === "end" || how === "right" || how === "bottom") {
    return "flex-end";
  }
  if (how === "middle") {
    return "center";
  }
  return how;
};

export const contentAlignment = (
  alignmentX?: Alignment,
  alignmentY?: Alignment,
  direction = "row"
) => {
  const realAlignmentX = replaceAlignment(alignmentX);
  const realAlignmentY = replaceAlignment(alignmentY);
  if (direction === "column" || direction === "col") {
    return `
            display: flex;
            flex-direction: column;
            ${
              realAlignmentY !== null
                ? `justify-content: ${realAlignmentY};`
                : ""
            }
            ${realAlignmentX !== null ? `align-items: ${realAlignmentX};` : ""}
        `;
  }

  return `
        display: flex;
        flex-direction: row;
        ${realAlignmentX !== null ? `justify-content: ${realAlignmentX};` : ""}
        ${realAlignmentY !== null ? `align-items: ${realAlignmentY};` : ""}
    `;
};
