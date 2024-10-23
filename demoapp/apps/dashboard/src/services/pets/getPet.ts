import { JsonDecoder } from "ts.data.json";
import {fetchJSON} from "../../util/fetch";
import {petDecoder, Pet} from "../../models/pet";

export type GetPetRequest = {
	id: string;
};

export type GetPetResponse = {
	pet: Pet;
};

const getPetResponseDecoder = JsonDecoder.object(
	{
		pet: petDecoder,
	},
	"GetPetResponse"
);

export const getPet = async (request: GetPetRequest, token: string): Promise<GetPetResponse> => {
	const response = await fetchJSON('/v3/pet/get', request, token);

	return getPetResponseDecoder.decodeToPromise(response);
};
