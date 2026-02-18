<script lang="ts" setup>
	const title = ref('Mega');
</script>

<template>
	<Head>
		<Meta charset="UTF-8" />
		<Meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" />
		<link rel="icon" href="/favicon.ico" sizes="32x32" />
		<Link rel="icon" href="/favicon.svg" type="image/svg+xml" />
		<Link rel="apple-touch-icon" href="/apple-touch-icon.png" />
		<Link rel="manifest" href="/manifest.json" />
		<Title>{{ title }}</Title>
	</Head>
	<Body>
		<h1>Hello World, {{ title }}!</h1>
	</Body>
</template>
