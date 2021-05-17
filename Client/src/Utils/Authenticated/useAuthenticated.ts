import authenticated from '../../Store/authenticated';

const AUTH_CONTEXT_NAME = 'authenticated';

export const shouldPropLogin = (authenticated: boolean) => {
	const existingUser = Boolean(localStorage.getItem(AUTH_CONTEXT_NAME));
	if (existingUser && !authenticated) {
		setAuthenticated(existingUser);
	}

	return !(authenticated || existingUser);
};

export const setAuthenticated = (auth: boolean) => {
	authenticated.set(auth);
	localStorage.setItem(AUTH_CONTEXT_NAME, String(auth));
};
