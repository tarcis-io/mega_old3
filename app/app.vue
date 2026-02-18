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
	<UApp>
		<h1>Hello World, {{ title }}!</h1>
	</UApp>
</template>
