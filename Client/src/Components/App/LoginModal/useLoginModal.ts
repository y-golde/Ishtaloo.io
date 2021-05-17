import axios from 'axios';
import { setAuthenticated } from '../../../Utils/Authenticated/useAuthenticated';

const UseLoginModal = () => {
	const handleLoginClick = async (userName: string) => {
		const authenticated = await authenticate(userName) 
		setAuthenticated(authenticated);
	};

	const shouldDisable = (userName: string) => {
		return userName === '';
	};

	const authenticate = async (userName: string) => {
		const body = {
			userName
		}
		return await axios.post('/auth', body)
			.then(res => {
				return res.status === 200;
			})
			.catch(err => {
				return false;
			});
	}

	return {
		handleLoginClick,
		shouldDisable,
	};
};

export default UseLoginModal;
