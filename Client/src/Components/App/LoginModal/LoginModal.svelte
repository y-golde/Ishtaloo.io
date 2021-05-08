<script lang="ts">
	import Textfield from '../../Common/Form/Textfield.svelte';
	import Modal from '../../Common/Modal/Modal.svelte';
	import ModalHeader from '../../Common/Modal/ModalHeader.svelte';

	export let show: boolean;
	let userName = 'moti';

	const modalHeaderText = 'Login';
</script>

<Modal show="{show}">
	<ModalHeader slot="header" text="{modalHeaderText}" />
	<div slot="content">
		<Textfield bind:value="{userName}" />
	</div>
</Modal>
