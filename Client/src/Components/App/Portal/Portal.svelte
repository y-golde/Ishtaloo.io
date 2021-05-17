<script lang="ts">
	import { navigate } from 'svelte-routing';

	import UsePortal from './usePortal';

	const { fetchRooms, joinRoom } = UsePortal();

	(async () => {
		// this is a dumb stub route - get rid of it next version
		const rooms = await fetchRooms();
		if (rooms) {
			const { roomId } = rooms[0];
			const succssesful = await joinRoom(roomId);
			if (succssesful) {
				navigate(`/room/${roomId}`);
			}
		}
	})();
</script>

<h1>portal</h1>
