import { FromDecoder, JsonDecoder } from "ts.data.json";

export const petDecoder = JsonDecoder.object(
	{
		id: JsonDecoder.string,
		name: JsonDecoder.string,
		status: JsonDecoder.string,
	},
	"Pet"
);

export type Pet = FromDecoder<typeof petDecoder>;
