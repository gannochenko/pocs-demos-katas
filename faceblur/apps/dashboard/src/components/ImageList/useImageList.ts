import {PetListProps} from "./type";
import {useListImages} from "../../hooks";
import {Image} from "../../models/image";
import {useState, useEffect, useRef} from "react";
import {GetUploadURL} from "../../proto/image/v1/image";
import {useAuth0} from "@auth0/auth0-react";

enum UploadStage {
	INITIAL,
	URL,
	SEND,
}

type Upload = {
	file: File;
	cancelled: boolean;
	url?: string;
};

type UploaderOnChangeCallback = (uploads: Upload[]) => void;

class Uploader {
	uploads: Upload[] = [];
	onChange: UploaderOnChangeCallback = () => {};
	timer?: NodeJS.Timer;
	stopped = false;

	async scheduleUpload(files: File[], token: string) {
		for (const file of files) {
			this.uploads.push({
				file,
				cancelled: false,
			});
		}
	}

	start(onChange: UploaderOnChangeCallback) {
		console.log('START');
		this.stopped = false;
		this.onChange = onChange;

		const onTimer = async () => {
			if (this.stopped) {
				return;
			}

			console.log('!!!');
			// ...


			// structuredClone(

			for(const upload of this.uploads) {
				const url = await GetUploadURL({}, token);
			}

			this.timer = setTimeout(onTimer, 300);
		};

		onTimer();
	}

	stop() {
		console.log('STOP');
		this.stopped = true;
		if (this.timer) {
			// @ts-expect-error
			clearTimeout(this.timer);
		}
	}
}

export function useImageList(props: PetListProps) {
	const imagesResult = useListImages({pageNumber: 1});
	const images: Image[] = imagesResult.data?.images ?? [];
	const [uploads, setUploads] = useState<Upload[]>([]);

	const { getAccessTokenSilently } = useAuth0();

	const uploader = useRef(new Uploader());
	useEffect(() => {
		uploader.current.start(setUploads);

		return () => {
			uploader.current.stop();
		};
	}, []);

	return {
		images,
		uploadButtonProps: {
			onChange: async (files: File[]) => {
				if (files.length) {
					await uploader.current.scheduleUpload(files, await getAccessTokenSilently());
				}
			}
		},
	}
}
