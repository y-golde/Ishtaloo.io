<script lang="ts">
	import axios from 'axios';
	import { onMount } from 'svelte';

	let name = 'world';
	let message = 'loading...';
	

	// env later
	onMount(async () => {
		const res = await axios.get('/')
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

<h1>Hello {name}!</h1>

<h2>Message from server: {message}</h2>