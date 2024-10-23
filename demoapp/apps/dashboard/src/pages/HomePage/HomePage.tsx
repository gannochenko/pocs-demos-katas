import React, {useState} from 'react';
import {PageLayout, SEO} from "../../components";
import {Container} from "@mui/joy";
import {PetList} from "../../components/PetList";
import {EditPetDialog} from "../../components/EditPetDialog/EditPetDialog";

export function HomePage() {
	const [dialogOpen, setDialogOpen] = useState(false);
	const [petID, setPetID] = useState("");

	return (
		<PageLayout title="Pets" displayPageTitle>
			<Container>
				<PetList onRowClick={(petID) => {
					setPetID(petID);
					setDialogOpen(true);
				}} />
				<EditPetDialog open={dialogOpen} petID={petID} onClose={() => {
					setDialogOpen(false)
				}} />
			</Container>
		</PageLayout>
	);
}
