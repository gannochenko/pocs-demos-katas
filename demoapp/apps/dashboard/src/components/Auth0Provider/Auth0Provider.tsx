import { Auth0Provider as Auth0ProviderLib, AppState } from "@auth0/auth0-react";
import React, { PropsWithChildren } from "react";
import { useNavigate } from "react-router-dom";

type Auth0ProviderWithNavigateProps = {
	children: React.ReactNode;
}

export const Auth0Provider = ({
  children,
}: PropsWithChildren<Auth0ProviderWithNavigateProps>): JSX.Element | null => {
	const navigate = useNavigate();

	const domain = process.env.REACT_APP_AUTH0_DOMAIN;
	const clientId = process.env.REACT_APP_AUTH0_CLIENT_ID;
	const redirectUri = process.env.REACT_APP_AUTH0_CALLBACK_URL;
	const audience = process.env.REACT_APP_AUTH0_AUDIENCE;

	const onRedirectCallback = (appState?: AppState) => {
		navigate(appState?.returnTo || window.location.pathname);
	};

	if (!(domain && clientId && redirectUri)) {
		return null;
	}

	return (
		<Auth0ProviderLib
			domain={domain}
			clientId={clientId}
			authorizationParams={{
				redirect_uri: redirectUri,
				audience,
			}}
			onRedirectCallback={onRedirectCallback}
		>
			{children}
		</Auth0ProviderLib>
	);
};
