<script lang="ts">
	import user from './Store/user';
	import { shouldPropLogin } from './Utils/User/useUser';
	import Headline from './Components/Common/Headline.svelte';
	import Game from './Components/App/Game/Game.svelte';
	import ThemeWrapper from './CSS/ThemeWrapper.svelte';
	import LoginModal from './Components/App/LoginModal/LoginModal.svelte';

	let userName: string = '';
	const unsubscribe = user.subscribe((value) => {
		userName = value;
	});

	let showLoginModal = false;
	$: showLoginModal = shouldPropLogin(userName);

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
