import { JsonDecoder } from "ts.data.json";
import {fetchJSON} from "../../util/fetch";
import {petDecoder, Pet} from "../../models/pet";

export type ListPetsRequest = {};

export type ListPetsResponse = {
	pets: Pet[];
};

const listPetsResponseDecoder = JsonDecoder.object(
	{
		pets: JsonDecoder.array(petDecoder, "arrayOfPets"),
	},
	"ListPetsResponse"
);

export const listPets = async (request: ListPetsRequest, token: string): Promise<ListPetsResponse> => {
	const response = await fetchJSON('/v3/pet/list', null, token);

	return listPetsResponseDecoder.decodeToPromise(response);
};
