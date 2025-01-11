import React from 'react';
import {PageLayout} from "../../components";
import {Container} from "@mui/joy";
import {ImageList} from "../../components/ImageList";

export function HomePage() {
	return (
		<PageLayout title="Images" displayPageTitle>
			<Container>
				<ImageList />
			</Container>
		</PageLayout>
	);
}
