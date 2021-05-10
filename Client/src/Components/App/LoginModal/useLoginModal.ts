import { setUser } from '../../../Utils/User/useUser';

const UseLoginModal = () => {
	const handleLoginClick = (userName: string) => {
		setUser(userName);
	};

	const shouldDisable = (userName: string) => {
		return userName === '';
	};

	return {
		handleLoginClick,
		shouldDisable,
	};
};

export default UseLoginModal;
