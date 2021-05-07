<script lang="ts">
	import Headline from './Common/Headline.svelte';
	import Game from './Components/Game/Game.svelte';
	import ThemeWrapper from './CSS/ThemeWrapper.svelte';

	const headlineText = 'ishtaloo.io';
</script>

<ThemeWrapper>
	<div class="main-wrapper">
		<Headline text="{headlineText}" />
		<Game />
	</div>
</ThemeWrapper>

<style global lang="postcss">
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
	html {
		background-color: var(--background);
	}

	html * {
		font-family: 'Source Sans Pro', sans-serif;
		font-weight: 400;
	}
</style>
