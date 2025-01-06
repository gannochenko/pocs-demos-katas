import { useQuery } from "react-query";
import {useAuth0} from "@auth0/auth0-react";
import {listCategories, ListCategoriesRequest} from "../../services/category/listCategories";
export const LIST_CATEGORIES_KEY = "list_categories";

export function useListCategories(request: ListCategoriesRequest) {
	const { getAccessTokenSilently } = useAuth0();
	return useQuery(
		[LIST_CATEGORIES_KEY, ...[request]],
		async () => {
			return listCategories(request, await getAccessTokenSilently());
		},
		{
			keepPreviousData: true,
			refetchOnWindowFocus: false,
		}
	);
}
