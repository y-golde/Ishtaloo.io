<script>
	import CardHeader from './CardHeader/CardHeader.svelte';
	import PlayersList from './PlayersList/PlayersList.svelte';
</script>

<div class="panel-card w-full">
	<CardHeader />
	<PlayersList />
</div>

<style>
	.panel-card {
		/* stub */
		min-height: 400px;

		background-color: rgba(0, 0, 0, 0.01);

		border: 1px solid rgba(0, 0, 0, 0.1);
		border-radius: 15px;
	}
</style>
