import authenticated from '../../Store/authenticated';

const AUTH_CONTEXT_NAME = 'authenticated';

export const shouldPropLogin = (authenticated: boolean) => {
	const existingUser = Boolean(sessionStorage.getItem(AUTH_CONTEXT_NAME));
	if (existingUser && !authenticated) {
		setAuthenticated(existingUser);
	}

	return !(authenticated || existingUser);
};

export const setAuthenticated = (auth: boolean) => {
	authenticated.set(auth);
	sessionStorage.setItem(AUTH_CONTEXT_NAME, String(auth));
};
