import {fetchJSON} from "../../util/fetch";
import {Pet} from "../../models/pet";

export type UpdatePetRequest = {
	pet: Pet;
};

export type UpdatePetResponse = {
};

export async function updatePet(request: UpdatePetRequest, token: string): Promise<UpdatePetResponse> {
	return  await fetchJSON('/v3/pet/update', request, token);
}
