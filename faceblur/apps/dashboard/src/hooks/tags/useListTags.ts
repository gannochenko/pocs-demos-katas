import { useQuery } from "react-query";
import {useAuth0} from "@auth0/auth0-react";
import {listTags, ListTagsRequest} from "../../services/tags/listTags";
export const LIST_TAGS_KEY = "list_tags";

export function useListTags(request: ListTagsRequest) {
	const { getAccessTokenSilently } = useAuth0();
	return useQuery(
		[LIST_TAGS_KEY, ...[request]],
		async () => {
			return listTags(request, await getAccessTokenSilently());
		},
		{
			keepPreviousData: true,
			refetchOnWindowFocus: false,
		}
	);
}
