const UsePlayersList = () => {
	const fetchPlayers = async () => {
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
