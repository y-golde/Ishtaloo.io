<script lang="ts">
	import type Player from 'Models/Game/Player';

	import strToGradient from 'Utils/strToGradient';

	export let player: Player;

	const playerColor = `background: ${strToGradient(player.userName)}`;
</script>

<div class="p-2 text-lg wrapper">
	<div class="flex flex-row align-center">
		<div class="flex-grow-0">
			<div class="active-circle align-middle" style="{playerColor}"></div>
		</div>
		<div class="flex-grow-0">
			<div class="user-name align-middle">{player.userName}</div>
		</div>
	</div>
</div>

<style>
	.wrapper {
		background-color: rgba(0, 0, 0, 0.05);
		border-bottom: 1px solid rgba(0, 0, 0, 0.2);
	}

	.active-circle {
		width: 15px;
		height: 15px;

		display: inline-block;
		margin-right: 0.5rem;

		border-radius: 50%;
	}

	.user-name {
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}
</style>
