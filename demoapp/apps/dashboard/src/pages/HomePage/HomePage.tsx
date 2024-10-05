import React from 'react';
import {PageLayout, SEO} from "../../components";
import {Container} from "@mui/joy";
import {PetList} from "../../components/PetList";

export function HomePage() {
	return (
		<PageLayout title="Pets" displayPageTitle>
			<Container>
				<PetList />
			</Container>
		</PageLayout>
	);
}
