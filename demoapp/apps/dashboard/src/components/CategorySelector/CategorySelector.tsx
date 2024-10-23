import Option from "@mui/joy/Option";
import Select from "@mui/joy/Select";
import {useListCategories} from "../../hooks/categories/useListCategories";
import {Category} from "../../models/category";

type CategorySelectorProps = Partial<{
	value: string;
	onChange: (newValue: string) => void;
}>;

export function CategorySelector({ value, onChange }: CategorySelectorProps) {
	const categoriesResult = useListCategories({});
	const categories: Category[] = categoriesResult.data?.categories ?? [];

	return (
		<Select value={value} onChange={(_, newValue) => {
			if (newValue) {
				onChange?.(newValue);
			}
		}}>
			{
				categories.map(({id, name}) => {
					return (<Option value={id} key={id}>{name}</Option>);
				})
			}
		</Select>
	);
}
