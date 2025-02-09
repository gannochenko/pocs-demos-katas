import Button from "@mui/joy/Button";
import {useRef, ChangeEvent} from "react";

type UploadButtonProps = Partial<{
	onChange: (files: File[]) => Promise<void>;
}>;

export function UploadButton({onChange}: UploadButtonProps) {
	const inputRef = useRef<HTMLInputElement | null>(null);

	const handleButtonClick = () => {
		if (inputRef.current) {
			inputRef.current.click();
		}
	};

	const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
		const files = event.target.files;
		if (files && files.length > 0) {
			onChange?.(Array.from(files));
		}
	};

	return (
		<div>
			<Button onClick={handleButtonClick}>Upload images</Button>
			<input
				type="file"
				id="photoInput"
				accept="image/*"
				multiple
				style={{ display: "none" }}
				ref={inputRef}
				onChange={handleFileChange}
			/>
		</div>
	);
}
