import Option from "@mui/joy/Option";
import Select from "@mui/joy/Select";

export enum Status {
	Available = "available",
	Pending = "pending",
	Sold = "sold",
}

const labels = new Map<Status, string>([
	[Status.Available, "Available"],
	[Status.Pending, "Pending"],
	[Status.Sold, "Sold"],
]);

type StatusSelectorProps = Partial<{
	value: Status;
	onChange: (newValue: Status) => void;
}>;

export function StatusSelector({ value, onChange }: StatusSelectorProps) {
	return (
		<Select value={value} onChange={(_, newValue) => {
			onChange?.(newValue ?? Status.Pending);
		}}>
			{
				Array.from(labels.keys()).map((key) => {
					return (<Option value={key} key={key}>{labels.get(key)}</Option>);
				})
			}
		</Select>
	);
}
