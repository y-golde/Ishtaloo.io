<script lang="ts">
	import Headline from './Common/Headline.svelte';
	import Game from './Components/Game/Game.svelte';
import ThemeWrapper from './CSS/ThemeWrapper.svelte';

	const headlineText = 'ishtaloo.io'
</script>

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
  html {
	  background-color: var(--background);
  }
</style>

<ThemeWrapper>
	<div class="main-wrapper">
		<Headline text={headlineText}></Headline>
		<Game />
	</div>
</ThemeWrapper>