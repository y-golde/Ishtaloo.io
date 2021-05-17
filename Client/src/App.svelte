<script lang="ts">
	import authenticated from './Store/authenticated';
	import { shouldPropLogin } from './Utils/Authenticated/useAuthenticated';
	import Headline from './Components/Common/Headline.svelte';
	import Game from './Components/App/Game/Game.svelte';
	import ThemeWrapper from './CSS/ThemeWrapper.svelte';
	import LoginModal from './Components/App/LoginModal/LoginModal.svelte';

	let isAuthenticated = false;
	const unsubscribe = authenticated.subscribe((value) => {
		isAuthenticated = value;
	});

	let showLoginModal = false;
	$: showLoginModal = shouldPropLogin(isAuthenticated);

	const headlineText = 'ishtaloo.io';
</script>

<ThemeWrapper>
	<div class="main-wrapper">
		<Headline text="{headlineText}" />
		<Game />
	</div>
	<LoginModal show="{showLoginModal}" />
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
