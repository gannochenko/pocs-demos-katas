import { useQuery } from "react-query";
import {useAuth0} from "@auth0/auth0-react";
import {getPet, GetPetRequest} from "../../services/pets/getPet";
export const GET_PET_KEY = "get_pet";

export function useGetPet(request: GetPetRequest) {
	const { getAccessTokenSilently } = useAuth0();
	return useQuery(
		[GET_PET_KEY, ...[request]],
		async () => {
			return getPet(request, await getAccessTokenSilently());
		},
		{
			refetchOnWindowFocus: false,
			enabled: !!request.id,
		}
	);
}
