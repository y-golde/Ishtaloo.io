<script lang="ts">
	import { navigate } from 'svelte-routing';

	import UsePortal from './usePortal';
	import authenticated from 'Store/authenticated';

	const { fetchRooms, joinRoom } = UsePortal();

	let isAuthenticated = false;
	const unsubscribe = authenticated.subscribe((value) => {
		isAuthenticated = value;
	});

	const autoJoinRoom = async () => {
		// this is a dumb stub route - get rid of it next version
		const rooms = await fetchRooms();
		if (rooms) {
			const { roomId } = rooms[0];
			const succssesful = await joinRoom(roomId);
			if (succssesful) {
				navigate(`/room/${roomId}`);
			}
		}
	};

	$: {
		if (isAuthenticated) {
			autoJoinRoom();
			unsubscribe();
		}
	}
</script>

<h1>portal</h1>
