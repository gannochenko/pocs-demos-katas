import React from 'react';
import {PageLayout, SEO} from "../../components";
import {Container} from "@mui/joy";
import {PetList} from "../../components/PetList/PetList";

export function HomePage() {
	return (
		<PageLayout>
			<Container>
				<PetList />
			</Container>
		</PageLayout>
	);
}
