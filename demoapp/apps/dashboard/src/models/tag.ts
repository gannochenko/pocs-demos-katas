import { FromDecoder, JsonDecoder } from "ts.data.json";

export const tagDecoder = JsonDecoder.object(
	{
		id: JsonDecoder.string,
		name: JsonDecoder.string,
	},
	"Tag"
);

export type Tag = FromDecoder<typeof tagDecoder>;
