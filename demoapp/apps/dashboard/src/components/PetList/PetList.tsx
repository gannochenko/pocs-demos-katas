import Table from '@mui/joy/Table';
import {useListPets} from "../../hooks/pets/useListPets";
import {PetListRoot} from "./style";
import {Pet} from "../../models/pet";
import {Link} from "../Link";

type PetListProps = Partial<{
	onRowClick: (petID: string) => void;
}>;

export function PetList({onRowClick}: PetListProps) {
	const petsResult = useListPets({});
	const pets: Pet[] = petsResult.data?.pets ?? [];
	
	return (
		<PetListRoot>
			<Table aria-label="basic table">
				<thead>
				<tr>
					<th style={{ width: '40%' }}>Name</th>
					<th>Status</th>
					<th />
				</tr>
				</thead>
				<tbody>
				{
					pets.map(pet => {
						return (
							<tr key={pet.id}>
								<td>{pet.name}</td>
								<td>{pet.status}</td>
								<td align="right">
									<Link onClick={() => onRowClick?.(pet.id)}>Edit</Link>
								</td>
							</tr>
						);
					})
				}

				</tbody>
			</Table>
		</PetListRoot>
	);
}
