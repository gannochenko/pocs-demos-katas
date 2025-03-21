import { useEffect, useState } from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { PetListProps } from "./type";
import { LIST_IMAGES_KEY, useListImages } from "@/hooks";
import { ImageModel, Upload } from "@/models/image";
import { isErrorResponse, uploadFile } from "@/util/fetch";
import { useNotification } from "@/hooks/notification";
import { useWebsocketContext } from "../WebsocketProvider";
import {
  ServerMessage,
  ServerMessageType,
} from "@/proto/websocket/v1/websocket";
import { GetUploadURL, SubmitImage } from "@/proto/image/v1/image";
import { useQueryClient } from "react-query";

const useUploader = () => {
  const [uploads, setUploads] = useState<Upload[]>([]);
  const { getAccessTokenSilently } = useAuth0();
  const { showError } = useNotification();

  const updateUpload = (id: string, updateCb: (upload: Upload) => Upload) => {
    setUploads((prevState) => {
      return prevState.map((upload) => {
        if (upload.id === id) {
          return updateCb(upload);
        }

        return upload;
      });
    });
  };

  const showUploadError = (reason: string) => {
    showError("Error uploading file", reason);
  };

  return {
    uploads,
    submit: async (newUploads: Upload[]) => {
      newUploads = newUploads.sort(
        (a, b) =>
          (b.uploadedAt?.getDate() ?? 0) - (a.uploadedAt?.getDate() ?? 0)
      );

      setUploads((prevUploads) => {
        return [...newUploads, ...prevUploads];
      });

      let token = "";
      try {
        token = await getAccessTokenSilently();
      } catch (e) {
        showError("Unauthorized");
        return;
      }

      for (const upload of newUploads) {
        const getUploadULRResponse = await GetUploadURL({}, token);
        if (isErrorResponse(getUploadULRResponse)) {
          showUploadError(getUploadULRResponse.error);
          updateUpload(upload.id, (upload) => ({
            ...upload,
            failed: true,
          }));
        } else {
          await uploadFile(
            getUploadULRResponse.url,
            upload.file!,
            (newProgress) => {
              updateUpload(upload.id, (upload) => ({
                ...upload,
                progress: newProgress,
              }));
            }
          );
          const submitImageResponse = await SubmitImage(
            {
              image: {
                objectName: getUploadULRResponse.objectName,
                uploadedAt: upload.uploadedAt!,
              },
            },
            token
          );
          if (isErrorResponse(submitImageResponse)) {
            showUploadError(submitImageResponse.error);
            updateUpload(upload.id, (upload) => ({
              ...upload,
              failed: true,
            }));
          } else {
            updateUpload(upload.id, (upload) => ({
              ...upload,
              image: submitImageResponse.image,
            }));
          }
        }
      }
    },
  };
};

export function useImageList(props: PetListProps) {
  const queryClient = useQueryClient();

  const imagesResult = useListImages({ pageNumber: 1 });
  const images: ImageModel[] = imagesResult.data?.images ?? [];

  const { addEventListener, removeEventListener } = useWebsocketContext();

  useEffect(() => {
    const handler = (payload: ServerMessage) => {
      queryClient.invalidateQueries([LIST_IMAGES_KEY]);
    };

    addEventListener(
      ServerMessageType.SERVER_MESSAGE_TYPE_IMAGE_PROCESSED,
      handler
    );

    return () =>
      removeEventListener(
        ServerMessageType.SERVER_MESSAGE_TYPE_IMAGE_PROCESSED,
        handler
      );
  }, [addEventListener, removeEventListener]);

  const { uploads, submit } = useUploader();

  const realUploads = uploads.filter(
    (upload) =>
      images.find((image) => image.id === upload.image?.id) === undefined
  );

  return {
    uploads: realUploads,
    images,
    empty: !images.length && !realUploads.length && !imagesResult.isLoading,
    uploadButtonProps: {
      onChange: async (files: File[]) => {
        if (files.length) {
          submit(
            files.map((file) => ({
              id: Math.floor(Math.random() * 100000).toString(),
              file,
              uploadedAt: new Date(),
              progress: 0,
            }))
          );
        }
      },
    },
    getImageUploadProps: (upload: Upload) => {
      return {
        upload,
      };
    },
    getImageUploadPropsByImage: (image: ImageModel) => {
      return {
        upload: {
          id: image.id,
          file: undefined,
          image,
          uploadedAt: image.updatedAt,
          failed: image.isFailed,
          progress: 100,
        },
      };
    },
  };
}
