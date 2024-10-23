import { JsonDecoder } from "ts.data.json";
import {fetchJSON} from "../../util/fetch";
import {Category, categoryDecoder} from "../../models/category";

export type ListCategoriesRequest = {};

export type ListCategoriesResponse = {
	categories: Category[];
};

const listCategoriesResponseDecoder = JsonDecoder.object(
	{
		categories: JsonDecoder.array(categoryDecoder, "arrayOfCategories"),
	},
	"ListCategoriesResponse"
);

export const listCategories = async (request: ListCategoriesRequest, token: string): Promise<ListCategoriesResponse> => {
	const response = await fetchJSON('/v3/categories/list', null, token);

	return listCategoriesResponseDecoder.decodeToPromise(response);
};
