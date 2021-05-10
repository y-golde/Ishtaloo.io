<script lang="ts">
	import UseLoginModal from './useLoginModal';
	import Button from '../../Common/Button/Button.svelte';
	import Textfield from '../../Common/Form/Textfield.svelte';
	import Modal from '../../Common/Modal/Modal.svelte';
	import ModalHeader from '../../Common/Modal/ModalHeader.svelte';

	export let show: boolean;

	let userName = '';
	let { handleLoginClick, shouldDisable } = UseLoginModal();

	const modalHeaderText = 'Login';
	const placeHolderText = 'nickname...';

	let disabled = true;
	$: {
		disabled = shouldDisable(userName);
	}
</script>

<Modal show="{show}">
	<ModalHeader slot="header" text="{modalHeaderText}" />
	<div slot="content" class="text-2xl">
		<p class="p-3">Hello, please provide a nickname:</p>
		<div class="flex flex-wrap p-3 gap-2 items-center justify-center">
			<Textfield bind:value="{userName}" placeholder="{placeHolderText}" />
			<Button
				disabled="{disabled}"
				handleClick="{() => handleLoginClick(userName)}"
				text="{'sign in'}"
			/>
		</div>
	</div>
</Modal>
