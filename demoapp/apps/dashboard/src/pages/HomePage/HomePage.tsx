import React from 'react';
import {PageLayout, SEO} from "../../components";
import {Container} from "@mui/joy";
import {PetList} from "../../components/PetList";
import {EditPetDialog} from "../../components/EditPetDialog/EditPetDialog";

export function HomePage() {
	return (
		<PageLayout title="Pets" displayPageTitle>
			<Container>
				<PetList />
				<EditPetDialog />
			</Container>
		</PageLayout>
	);
}
