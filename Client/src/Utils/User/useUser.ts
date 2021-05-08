import user from '../../Store/user';

const USER_CONTEXT_NAME = 'user';

export const shouldPropLogin = (userName: string) => {
	const existingUser = localStorage.getItem(USER_CONTEXT_NAME);
	if (existingUser && !userName) {
		setUser(existingUser);
	}

	return !Boolean(userName || existingUser);
};

export const setUser = (userName: string) => {
	user.set(userName);
	localStorage.setItem(USER_CONTEXT_NAME, userName);
};
