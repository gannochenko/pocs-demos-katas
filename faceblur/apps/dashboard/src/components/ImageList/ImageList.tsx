import Grid from "@mui/joy/Grid";
import { Root, ImageItem, NoImages } from "./style";
import React, { useEffect } from "react";
import { PortalToID } from "../PortalToID/PortalToID";
import { UploadButton } from "../UploadButton";
import { Image } from "../Image";
import { PetListProps } from "./type";
import { useImageList } from "./useImageList";

export function ImageList(props: PetListProps) {
  const {
    uploads,
    images,
    uploadButtonProps,
    getImageUploadProps,
    getImageUploadPropsByImage,
    empty,
  } = useImageList(props);

  return (
    <>
      <PortalToID id="page-header-portal">
        <UploadButton {...uploadButtonProps} />
      </PortalToID>
      <Root>
        {empty && <NoImages>No images so far</NoImages>}
        {!empty && (
          <Grid container spacing={2} sx={{ flexGrow: 1 }}>
            {uploads.map((upload) => (
              <Grid xs={4} key={upload.id}>
                <Image {...getImageUploadProps(upload)} />
              </Grid>
            ))}
            {images.map((image) => (
              <Grid xs={4} key={image.id}>
                <Image {...getImageUploadPropsByImage(image)} />
              </Grid>
            ))}
          </Grid>
        )}
      </Root>
    </>
  );
}
