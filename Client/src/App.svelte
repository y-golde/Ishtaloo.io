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
  @import url('https://fonts.googleapis.com/css2?family=Source+Sans+Pro:ital,wght@0,400;0,600;0,700;1,400&display=swap');
  html {
	  background-color: var(--background);
  }

  html * {
	font-family: 'Source Sans Pro', sans-serif;
  }
</style>

<ThemeWrapper>
	<div class="main-wrapper">
		<Headline text={headlineText}></Headline>
		<Game />
	</div>
</ThemeWrapper>