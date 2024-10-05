import { useQuery } from "react-query";
import {listPets, ListPetsRequest} from "../services/listPets";
import {useAuth0} from "@auth0/auth0-react";
export const LIST_PETS_KEY = "list_pets";

export function useListPets(request: ListPetsRequest) {
	const { getAccessTokenSilently } = useAuth0();
	return useQuery(
		[LIST_PETS_KEY, ...[request]],
		async () => {
			return listPets(request, await getAccessTokenSilently());
		},
		{
			keepPreviousData: true,
			refetchOnWindowFocus: false,
		}
	);
}
