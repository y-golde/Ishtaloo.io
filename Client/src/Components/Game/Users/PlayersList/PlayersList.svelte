<script lang="ts">
	import { onMount } from 'svelte';
	import type Player from 'Models/Game/Player';

	import UsePlayersList from './usePlayersList';
	import PlayerCard from './Player/PlayerCard.svelte';

	let players: Player[] = [];
	const { getPlayers } = UsePlayersList();

	onMount(async () => {
		players = await getPlayers();
	});
</script>

<div class="grid grid-cols-1">
	{#each players as player}
		<div class="col-span-1">
			<PlayerCard player="{player}" />
		</div>
	{/each}
</div>
