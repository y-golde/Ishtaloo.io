<script lang="ts">
	import axios from 'axios';
	import { onMount } from 'svelte';
	
	import Headline from './Common/Headline.svelte';

	const headlineText = 'ishtaloo.io'

	let name = 'world';
	let message = 'loading...';
	
	onMount(async () => {
		const res = await axios.get('')
			.then(res => {
				return res.data;
			})
			.catch(err => {
				console.log(err);
				return '';
			})
		message = res;
	});
</script>

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>

<Headline text={headlineText}></Headline>

<h2>Message from server: {message}</h2>