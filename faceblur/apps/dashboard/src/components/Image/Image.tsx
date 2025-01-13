import {Root} from "./style";

type ImageProps = Partial<{
	url: string;
}>;

export function Image({ url }: ImageProps) {
	return <Root url={url ?? ""} />
}
