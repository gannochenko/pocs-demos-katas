import { useQuery } from "react-query";
import {listPets, ListPetsRequest} from "../services/listPets";
export const LIST_PETS_KEY = "list_pets";

export function useListPets(request: ListPetsRequest) {
	return useQuery(
		[LIST_PETS_KEY, ...[request]],
		() => {
			return listPets(request);
		},
		{
			keepPreviousData: true,
			refetchOnWindowFocus: false,
		}
	);
}
