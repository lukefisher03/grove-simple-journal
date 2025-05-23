<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { Editor } from '@tiptap/core';
	import StarterKit from '@tiptap/starter-kit';
	import Typography from '@tiptap/extension-typography';
	import Highlight from '@tiptap/extension-highlight';

	let { className = '' } = $props();

	let element: Element;
	let editor: Editor;

	let content = `
	<h1>
	Hello world
	</h1>
	<h2>
		Hello world
	</h2>
	<h3>
		Hello world
	</h3>
	<h4>
		Hello world
	</h4>
	<h5>
		Hello world
	</h5>
	<h6>
		Hello world
	</h6>
	<br/>
	<h4>This is a test paragraph!</h4>
	Lorem ipsum dolor sit amet, <strong>consectetur adipiscing elit, sed do</strong> eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
	<br/>
	<br/>
	Lorem ipsum dolor sit amet, consectetur <em>adipiscing</em> elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim <strong>ad minim veniam</strong>, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
	<br/>
	<br/>
	Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

	`;

	onMount(() => {
		editor = new Editor({
			element: element,
			extensions: [StarterKit, Highlight, Typography],
			content: content,
			onTransaction: () => {
				// force re-render so `editor.isActive` works as expected
				editor = editor;
			}
		});
	});

	onDestroy(() => {
		if (editor) {
			editor.destroy();
		}
	});
</script>

<div bind:this={element} class={className} id="tiptap-editor"></div>

<style>
</style>
