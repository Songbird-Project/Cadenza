<script lang="ts">
	import "./style.css";
	import { Theme } from "../wailsjs/go/main/App.js";

	let theme = Theme().then((result) => (theme = result));

	import(/* @vite-ignore */ `${theme}`);
</script>

<main>
	<p>Hello, World!</p>
</main>

<style>
</style>
