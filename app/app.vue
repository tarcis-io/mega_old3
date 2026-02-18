<script lang="ts" setup>
	const title = ref('Mega');
	const language = ref('en');
</script>

<template>
	<Html :lang="language">
		<Head>
			<Title>{{ title }}</Title>
		</Head>
	</Html>
	<h1>Hello World, {{ title }}!</h1>
</template>
