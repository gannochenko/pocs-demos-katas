import { FromDecoder, JsonDecoder } from "ts.data.json";

export const categoryDecoder = JsonDecoder.object(
	{
		id: JsonDecoder.string,
		name: JsonDecoder.string,
	},
	"Category"
);

export type Category = FromDecoder<typeof categoryDecoder>;
