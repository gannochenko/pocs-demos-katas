import { JsonDecoder } from "ts.data.json";
import {fetchJSON} from "../../util/fetch";
import {tagDecoder, Tag} from "../../models/tag";

export type ListTagsRequest = {};

export type ListTagsResponse = {
	tags: Tag[];
};

const listTagsResponseDecoder = JsonDecoder.object(
	{
		tags: JsonDecoder.array(tagDecoder, "arrayOfTags"),
	},
	"ListTagsResponse"
);

export const listTags = async (request: ListTagsRequest, token: string): Promise<ListTagsResponse> => {
	const response = await fetchJSON('/v3/tag/list', null, token);

	return listTagsResponseDecoder.decodeToPromise(response);
};
