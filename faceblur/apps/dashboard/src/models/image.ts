export type ImageModel = {
	id: string;
	url: string;
	isProcessed: boolean;
	isFailed: boolean;
	createdAt: Date | undefined;
	updatedAt: Date | undefined;
};

export type Upload = {
	id: string;
	file: File;
	uploadedAt: Date;
	image?: ImageModel;
	failed?: boolean;
	progress: number;
};
