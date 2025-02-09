export type ImageModel = {
	id: string;
	url: string;
	isProcessed: boolean;
	isFailed: boolean;
	createdAt: Date | null;
	updatedAt: Date | null;
};

export type Upload = {
	id: string;
	file?: File;
	uploadedAt: Date | null;
	image?: ImageModel;
	failed?: boolean;
	progress: number;
};
