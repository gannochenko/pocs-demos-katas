import React, {useEffect} from 'react';
import {PageLayout} from "../../components";
import {Container} from "@mui/joy";
import {ImageList} from "../../components/ImageList";

export function HomePage() {
	useEffect(() => {
		console.log("PAGE UNMOUNT");
	}, []);
	
	return (
		<PageLayout title="Images" displayPageTitle>
			<Container>
				<ImageList />
			</Container>
		</PageLayout>
	);
}
