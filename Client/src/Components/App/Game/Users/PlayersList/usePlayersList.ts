import { getContext } from 'svelte';

const UsePlayersList = () => {
	const roomId = getContext('roomId');

	const fetchPlayers = async () => {
		console.log(roomId);
		return [
			{
				id: '0e3dksms',
				username: 'danny',
			},
			{
				id: '0e3dkdken',
				username: 'yontan',
			},
		];
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
