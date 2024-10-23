import CheckCircle from '@mui/icons-material/CheckCircle';
import CheckCircleOutline from '@mui/icons-material/CheckCircleOutline';
import {useListTags} from "../../hooks/tags/useListTags";
import {Tag} from "../../models/tag";
import {TagSelectorRoot, TagSelectorChip} from "./style";

type TagSelectorProps = Partial<{
	value: string[];
	onChange: (newValue: string[]) => void;
}>;

export function TagSelector({value, onChange}: TagSelectorProps) {
	const tagsResult = useListTags({});
	const tags: Tag[] = tagsResult.data?.tags ?? [];

	return (
		<TagSelectorRoot>
			{tags?.map(tag => (
				<TagSelectorChip
					key={tag.id}
					variant="soft"
					size="lg"
					color={value?.includes(tag.id) ? "success" : "neutral"}
					startDecorator={value?.includes(tag.id) ? <CheckCircle /> : <CheckCircleOutline />}
					onClick={() => {
						let nextValue: string[] = [];
						if (value?.includes(tag.id)) {
							nextValue = nextValue.filter(id => id !== tag.id);
						} else {
							nextValue = [...(value ?? []), tag.id];
						}

						onChange?.(nextValue)
					}}
				>
					{tag.name}
				</TagSelectorChip>
			))}
		</TagSelectorRoot>
	);
}
