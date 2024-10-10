import {useMutation, useQuery, useQueryClient} from "react-query";
import {listPets, ListPetsRequest} from "../../services/pets/listPets";
import {useAuth0} from "@auth0/auth0-react";
import {updatePet, UpdatePetRequest} from "../../services/pets/updatePet";
import {LIST_PETS_KEY} from "./useListPets";

export function useUpdatePet() {
	const queryClient = useQueryClient();
	const { getAccessTokenSilently } = useAuth0();

	return useMutation(async (request: UpdatePetRequest) => updatePet(request, await getAccessTokenSilently()), {
		onSuccess: () => {
			queryClient.invalidateQueries(LIST_PETS_KEY);
		},
		onError: (error) => {
			// todo: notify user
			console.log("error", error);
		},
	});
}
