import {useListPets} from "../../hooks/useListPets";

export function PetList() {
	const petsResult = useListPets({});

	console.log(petsResult.data);

	return <div>!</div>;
}
