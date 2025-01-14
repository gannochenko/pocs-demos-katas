export type ImageUploadProps = Partial<{
	upload: {
		id: string;
		file: File;
		createdAt: Date;
	};
	onSuccess: (id: string) => void;
}>;
