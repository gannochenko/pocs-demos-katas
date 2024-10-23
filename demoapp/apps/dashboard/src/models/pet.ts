import { FromDecoder, JsonDecoder } from "ts.data.json";
import {categoryDecoder} from "./category";
import {tagDecoder} from "./tag";

export const petDecoder = JsonDecoder.object(
	{
		id: JsonDecoder.string,
		name: JsonDecoder.string,
		status: JsonDecoder.string,
		category: JsonDecoder.optional(categoryDecoder),
		tags: JsonDecoder.array(tagDecoder, "arrayOfTags")
	},
	"Pet"
);

export type Pet = FromDecoder<typeof petDecoder>;
