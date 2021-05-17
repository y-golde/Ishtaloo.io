import axios from 'axios';
import { getContext } from 'svelte';

const UsePlayersList = () => {
	const roomId = getContext('roomId');

	const fetchPlayers = async () => {
		const users = await axios.get(`/rooms/${roomId}/users`)
			.then(res => {
				return res.data;
			})
			.catch(err => {
				return [];
			})
		const filteredUsers = users.filter(user => user.userId && user.userName)
		return filteredUsers;
	};

	const getPlayers = async () => {
		const playersRes = await fetchPlayers();
		return playersRes;
	};

	return {
		getPlayers,
	};
};

export default UsePlayersList;
