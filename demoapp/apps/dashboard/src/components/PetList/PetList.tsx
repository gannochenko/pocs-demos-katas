import Table from '@mui/joy/Table';
import {useListPets} from "../../hooks/useListPets";
import {PetListRoot} from "./style";
import {Pet} from "../../models/pet";

export function PetList() {
	const petsResult = useListPets({});
	const pets: Pet[] = petsResult.data?.pets ?? [];
	
	return (
		<PetListRoot>
			<Table aria-label="basic table">
				<thead>
				<tr>
					<th style={{ width: '40%' }}>Name</th>
					<th>Status</th>
				</tr>
				</thead>
				<tbody>
				{
					pets.map(pet => {
						return (
							<tr key={pet.id}>
								<td>{pet.name}</td>
								<td>{pet.status}</td>
							</tr>
						);
					})
				}

				</tbody>
			</Table>
		</PetListRoot>
	);
}
