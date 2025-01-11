import { extendTheme } from "@mui/joy/styles";

declare module "@mui/joy/styles" {
  interface PaletteBackgroundOverrides {
    default: true;
    contrast: true;
    active: true;
    pronounced: true;

    header: true;
    subHeader: true;
    footer: true;
    subtle: true;
  }
  interface PaletteTextOverrides {
    default: true;
    easier: true;
  }

  interface ColorSchemeOverrides {
    default: true;
    contrast: true;
    active: true;
    pronounced: true;

    header: true;
    subHeader: true;
    footer: true;
    subtle: true;
  }
}

// https://mui.com/joy-ui/customization/default-theme-viewer/
export const theme = extendTheme({
  radius: { sm: "2px" },
  colorSchemes: {
    light: {
      palette: {
        background: {
          body: "#fefefe",
    //       contrast: "#efeeee",
    //       active: "#dedede",
    //       pronounced: "#707070",
    //
          header: "#efeeee",
    //       subHeader: "#f8f8f8",
           footer: "#252525",
    //       subtle: "#f8f8f8",
        },
        text: {
          primary: "#2a2a2a",
          // easier: "#484848",
        },
      },
    },
    dark: {
      palette: {
        background: {
          // header: "#000",
          // subHeader: "#000",
          // default: "#000",
        },
        text: {
          // default: "#fff",
        },
      },
    },
  },
  typography: {
    h1: {
      fontSize: "var(--joy-fontSize-xl3)",
      fontFamily: "Montserrat",
      fontWeight: 500,
    },
    h2: {
      fontSize: "var(--joy-fontSize-xl2)",
      fontFamily: "Montserrat",
      fontWeight: 500,
    },
    h3: {
      fontSize: "var(--joy-fontSize-xl)",
      fontFamily: "Montserrat",
      fontWeight: 500,
    },
    "body-lg": {
      fontFamily: "Victor Mono",
    },
    "body-md": {
      fontFamily: "Victor Mono",
    },
    "body-sm": {
      fontFamily: "Victor Mono",
    },
    "body-xs": {
      fontFamily: "Victor Mono",
    },
  },
  components: {
    JoyTypography: {
      defaultProps: {
        levelMapping: {
          h1: "h2",
          h2: "h2",
          h3: "h3",
          h4: "h3",
          // @ts-expect-error
          body1: "p",
          // "title-lg": "p",
          // "title-md": "p",
          // "title-sm": "p",
          // "body-md": "p",
          // "body-sm": "p",
          // "body-xs": "span",
        },
      },
    },
  },
});
