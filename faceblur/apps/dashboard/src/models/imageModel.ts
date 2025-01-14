export type Image = {
	id: string;
	url: string;
	isProcessed: boolean;
	isFailed: boolean;
	createdAt: Date | undefined;
	updatedAt: Date | undefined;
};
