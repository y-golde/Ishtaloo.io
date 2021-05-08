<script lang="ts">
	import UseLoginModal from './useLoginModal';
	import Button from '../../Common/Button/Button.svelte';
	import Textfield from '../../Common/Form/Textfield.svelte';
	import Modal from '../../Common/Modal/Modal.svelte';
	import ModalHeader from '../../Common/Modal/ModalHeader.svelte';

	export let show: boolean;

	let userName = '';
	let { handleLoginClick } = UseLoginModal();

	const modalHeaderText = 'Login';
	const placeHolderText = 'type here...';
</script>

<Modal show="{show}">
	<ModalHeader slot="header" text="{modalHeaderText}" />
	<div slot="content">
		<div class="flex flex-wrap p-3 gap-2 items-center text-md">
			<div>Name:</div>
			<Textfield bind:value="{userName}" placeholder="{placeHolderText}" />
			<div>
				<Button handleClick="{() => handleLoginClick(userName)}" />
			</div>
		</div>
	</div>
</Modal>
