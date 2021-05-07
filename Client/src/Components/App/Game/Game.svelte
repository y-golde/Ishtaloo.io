<script lang="ts">
	import Users from './Users/Users.svelte';
	import GamePanel from './GamePanel/GamePanel.svelte';
</script>

<div class="container mx-auto">
	<div class="grid grid-cols-3 gap-4">
		<div class="md:col-span-1 col-span-3">
			<Users />
		</div>
		<div class="md:col-span-2 col-span-3">
			<GamePanel />
		</div>
	</div>
</div>
