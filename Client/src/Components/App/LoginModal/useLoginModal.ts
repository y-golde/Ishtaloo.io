import axios from 'axios';
import { setUser } from '../../../Utils/User/useUser';

const UseLoginModal = () => {
	const handleLoginClick = (userName: string) => {
		console.log('hey');
		postAuth(userName);
		//setUser(userName);
	};

	const shouldDisable = (userName: string) => {
		return userName === '';
	};

	const postAuth = (userName: string) => {
		const body = {
			userName
		}

		axios.post('/auth', body)
			.then(res => {
				console.log(res);

			})
			.catch(err => {
				console.log(err);
			})
		

	}

	return {
		handleLoginClick,
		shouldDisable,
	};
};

export default UseLoginModal;
