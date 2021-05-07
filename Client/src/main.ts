import App from './App.svelte';
import axiosSetUp from './axiosSetUp';

axiosSetUp();

var app = new App({
	target: document.body,
});

export default app;
