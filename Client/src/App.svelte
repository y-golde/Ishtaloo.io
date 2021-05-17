<script lang="ts">
	import { Router, Route } from 'svelte-routing';

	import authenticated from './Store/authenticated';
	import Game from './Components/App/Game/Game.svelte';
	import ThemeWrapper from './CSS/ThemeWrapper.svelte';
	import Portal from './Components/App/Portal/Portal.svelte';
	import Headline from './Components/Common/Headline.svelte';
	import LoginModal from './Components/App/LoginModal/LoginModal.svelte';
	import { shouldPropLogin } from './Utils/Authenticated/useAuthenticated';

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
		<Router url="">
			<Route path="/">
				<Portal />
			</Route>
			<Route path="/room/:roomid" let:params>
				<Game roomId="{params.roomId}" />
			</Route>
		</Router>
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
