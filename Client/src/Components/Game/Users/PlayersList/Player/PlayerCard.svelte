<script lang="ts">
	import type Player from 'Models/Game/Player';

	import strToGradient from 'Utils/strToGradient';

	export let player: Player;

	const playerColor = `background: ${strToGradient(player.username)}`;
</script>

<div class="grid grid-cols-4 p-2 text-lg wrapper">
	<div class="col-span-1">
		<div class="active-circle" style="{playerColor}"></div>
	</div>
	<div class="col-span-3">
		{player.username}
	</div>
</div>

<style>
	.wrapper {
		background-color: rgba(0, 0, 0, 0.05);
		border-bottom: 1px solid rgba(0, 0, 0, 0.2);
	}

	.active-circle {
		width: 30px;
		height: 30px;

		border-radius: 50%;
	}
</style>
